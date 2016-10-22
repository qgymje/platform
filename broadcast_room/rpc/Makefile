pack: clean
	godep save

clean:
	@echo "removing pkgs.." && go clean

http:
	fswatch

rpc:
	fswatch

build:
	env GOOS=linux GOARCH=amd64 go build -o broadcast_room_rpc

sync:
	scp broadcast_room_rpc root@node1:~/platform/broadcast_room_rpc
	scp ./configs/* root@node1:~/platform/broadcast_room_rpc/configs/

