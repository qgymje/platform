package users

import (
	"errors"
	pb "platform/commons/protos/user"

	"golang.org/x/net/context"
)

type UserServer struct {
}

func (s *UserServer) getUserInfo(u *UserInfo) *pb.UserInfo {
	userInfo := pb.UserInfo{
		UserID:   u.ID,
		Name:     u.Name,
		Nickname: u.Nickname,
		Token:    u.Token,
		HeadImg:  u.HeadImg,
		RegTime:  u.RegTime.Unix(),
	}
	return &userInfo
}

func (s *UserServer) Register(ctx context.Context, regInfo *pb.RegisterInfo) (*pb.UserInfo, error) {
	var err error

	loginConfig := LoginConfig{
		Name:     regInfo.Name,
		Password: regInfo.Password,
	}
	config := RegisterConfig{
		Nickname: regInfo.Nickname,
	}
	config.LoginConfig = loginConfig

	r := NewRegister(&config)
	if err = r.Do(); err != nil {
		return nil, errors.New(r.ErrorCode().String())
	}

	u, _ := r.GetUserInfo()
	return s.getUserInfo(u), nil
}

func (s *UserServer) Auth(ctx context.Context, token *pb.Token) (*pb.UserInfo, error) {
	var err error

	t := NewToken()
	if ok, err := t.Verify(token.Token); !ok && err != nil {
		return nil, errors.New(t.ErrorCode().String())
	}

	u, err := t.GetUserInfo()
	if err != nil {
		return nil, errors.New(u.ErrorCode().String())
	}
	return s.getUserInfo(u), nil
}

func (s *UserServer) Login(ctx context.Context, in *pb.LoginInfo) (*pb.UserInfo, error) {

	config := LoginConfig{
		Name:     in.Name,
		Password: in.Password,
	}

	login := NewLogin(&config)
	if err := login.Do(); err != nil {
		return nil, errors.New(login.ErrorCode().String())
	}

	u, _ := login.GetUserInfo()
	return s.getUserInfo(u), nil
}

func (s *UserServer) Logout(ctx context.Context, token *pb.Token) (*pb.Status, error) {
	out := NewLogout(token.Token)
	if err := out.Do(); err != nil {
		return nil, errors.New(out.ErrorCode().String())
	}
	return &pb.Status{Success: true}, nil
}

func (s *UserServer) Info(ctx context.Context, userID *pb.UserID) (*pb.UserInfo, error) {
	ui := NewUserInfo()
	err := ui.GetByID(userID.UserID)
	if err != nil {
		return nil, errors.New(ui.ErrorCode().String())
	}
	return s.getUserInfo(ui), nil
}
