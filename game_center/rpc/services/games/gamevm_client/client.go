package gamevmClient

import (
	"bytes"
	"crypto/sha256"
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"net"

	"golang.org/x/protobuf/proto"

	pb "platform/tests/gamevm_mock_server/protos"
	"platform/utils"
)

func init() {
	log.SetFlags(log.Ltime | log.Lshortfile)
	utils.InitRander()
}

const (
	serverAddr = "192.168.0.233:6060"
	userID     = "57e2267ec86ab45af3d14806"
	gameID     = "57e226dac86ab45af3d14807"
	key        = "cloudzen"
	clientID   = 123456
)

// Request the server
type Request struct {
	ClientID uint32
	Len      uint32
	NameLen  uint32
}

type requestProtobuf struct {
	message *pb.RequestVmIp
}

func newRequestProtobuf(userID, gameID string) *requestProtobuf {
	req := &requestProtobuf{
		message: &pb.RequestVmIp{
			GameId:  proto.String(gameID),
			Account: proto.String(userID),
		},
	}

	req.prepare()
	return req
}

func (r *requestProtobuf) prepare() error {
	r.clientIP()
	r.key()
	r.sign(key)

	return nil
}

func (r *requestProtobuf) Name() []byte {
	n := proto.MessageName(r.message)
	utils.Dump(n)
	return []byte("request_vm_ip")
}

func (r *requestProtobuf) clientIP() {
	r.message.PalyerIp = proto.Uint32(ip2int(net.IPv4(192, 168, 0, 155)))
}

func (r *requestProtobuf) key() {
	randstr := utils.GetRandomName()
	r.message.Key = proto.String(randstr)
}

func (r *requestProtobuf) sign(secretKey string) {
	str := secretKey + r.message.GetKey()
	hasher := sha256.New()
	hasher.Write([]byte(str))
	sha := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	r.message.Sign = proto.String(sha)
}

func (r *requestProtobuf) Marshal() ([]byte, error) {
	msg, err := proto.Marshal(r.message)
	if err != nil {
		return nil, err
	}
	log.Println("pbbyte:", msg)
	var buff bytes.Buffer
	buff.Write(r.Name())
	buff.Write(msg)
	return buff.Bytes(), nil
}

// Response the response body
type Response struct {
	ServerID uint32 //who sent this message.
	Len      uint32 //length of total packet msg.
	NameLen  uint32 //length of msg's type
	Message  *pb.ResponseVmIp
}

func ip2int(ip net.IP) uint32 {
	if len(ip) == 16 {
		return binary.BigEndian.Uint32(ip[12:16])
	}
	return binary.BigEndian.Uint32(ip)
}

func int2ip(nn uint32) net.IP {
	ip := make(net.IP, 4)
	binary.BigEndian.PutUint32(ip, nn)
	return ip
}

func main() {
	conn, err := net.Dial("tcp", serverAddr)
	if err != nil {
		log.Fatal(err)
	}

	pbreq := newRequestProtobuf(userID, gameID)
	pbbytes, err := pbreq.Marshal()
	if err != nil {
		log.Fatal(err)
	}
	utils.Dump("pbbytes:", pbbytes)

	req := &Request{
		ClientID: clientID,
		NameLen:  uint32(len(pbreq.Name())),
	}
	totalLength := 12 + len(pbbytes) + int(req.NameLen)
	var buff bytes.Buffer
	binary.Write(&buff, binary.BigEndian, req.ClientID)
	binary.Write(&buff, binary.BigEndian, uint32(totalLength))
	binary.Write(&buff, binary.BigEndian, req.NameLen)
	pbbuffer := make([]byte, len(pbbytes))
	copy(pbbuffer[:], pbbytes[:])
	binary.Write(&buff, binary.BigEndian, pbbuffer)

	log.Println(buff.Bytes())

	conn.Write(buff.Bytes())

	/////////////////////////////////// ///////////////////////////////////
	for {
		respBytes := make([]byte, 1024) // using small tmo buffer for demonstrating
		n, err := conn.Read(respBytes)
		if err != nil {
			if err != io.EOF {
				fmt.Println("read error:", err)
			}
			break
		}
		fmt.Println("got", n, "bytes.")

		serverID := binary.BigEndian.Uint32(respBytes[0:4])
		log.Println("server id:", serverID)

		respTotalLength := binary.BigEndian.Uint32(respBytes[4:8])
		log.Println(respTotalLength)

		nameLength := binary.BigEndian.Uint32(respBytes[8:12])
		log.Println(nameLength)

		pbLengthStart := 12 + nameLength
		log.Println(pbLengthStart)

		pbBuffBytes := respBytes[pbLengthStart:respTotalLength]
		log.Println(len(pbBuffBytes))
		log.Println(pbBuffBytes)

		var pbResp pb.ResponseVmIp
		err = proto.Unmarshal(pbBuffBytes, &pbResp)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("vm ip:", int2ip(pbResp.GetVmIp()))
		log.Println("vm port:", pbResp.GetVmPort())
		log.Println("account:", pbResp.GetAccount())
		log.Println("game id:", pbResp.GetGameId())
	}
}
