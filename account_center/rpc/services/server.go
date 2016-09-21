package servers

import (
	"errors"
	"platform/account_center/rpc/services/sms"
	"platform/account_center/rpc/services/users"
	pb "platform/commons/protos/user"

	"golang.org/x/net/context"
)

// UserServer implement pb.UserServer
type UserServer struct {
}

func (s *UserServer) getUserInfo(u *users.UserInfo) *pb.UserInfo {
	userInfo := pb.UserInfo{
		UserID:    u.ID,
		Phone:     u.Phone,
		Email:     u.Email,
		Nickname:  u.Nickname,
		Token:     u.Token,
		Avatar:    u.Avatar,
		CreatedAt: u.CreatedAt.Unix(),
	}
	return &userInfo
}

// Register provider register rpc call
func (s *UserServer) Register(ctx context.Context, regInfo *pb.RegisterInfo) (*pb.UserInfo, error) {
	config := &users.RegisterConfig{
		Account:         regInfo.Account,
		Password:        regInfo.Password,
		PasswordConfirm: regInfo.PasswordConfirm,
		Nickname:        regInfo.Nickname,
	}
	r := users.NewRegister(config)
	if err := r.Do(); err != nil {
		return nil, errors.New(r.ErrorCode().String())
	}

	u, _ := r.GetUserInfo()
	return s.getUserInfo(u), nil
}

// Auth provider auth rpc call
func (s *UserServer) Auth(ctx context.Context, token *pb.Token) (*pb.UserInfo, error) {
	var err error

	t := users.NewToken()
	if ok, err := t.Verify(token.Token); !ok && err != nil {
		return nil, errors.New(t.ErrorCode().String())
	}

	u, err := t.GetUserInfo()
	if err != nil {
		return nil, errors.New(u.ErrorCode().String())
	}
	return s.getUserInfo(u), nil
}

// Login provider login rpc call
func (s *UserServer) Login(ctx context.Context, in *pb.LoginInfo) (*pb.UserInfo, error) {
	config := &users.LoginConfig{
		Account:  in.Account,
		Password: in.Password,
	}
	login := users.NewLogin(config)
	if err := login.Do(); err != nil {
		return nil, errors.New(login.ErrorCode().String())
	}

	u, _ := login.GetUserInfo()
	return s.getUserInfo(u), nil
}

// Logout provider logout rpc call
func (s *UserServer) Logout(ctx context.Context, token *pb.Token) (*pb.Status, error) {
	out := users.NewLogout(token.Token)
	if err := out.Do(); err != nil {
		return nil, errors.New(out.ErrorCode().String())
	}
	return &pb.Status{Success: true}, nil
}

// Info provider info rpc call
func (s *UserServer) Info(ctx context.Context, userID *pb.UserID) (*pb.UserInfo, error) {
	ui := users.NewUserInfo()
	err := ui.GetByID(userID.UserID)
	if err != nil {
		return nil, errors.New(ui.ErrorCode().String())
	}
	return s.getUserInfo(ui), nil
}

// ValidCode request a sms code before register
func (s *UserServer) ValidCode(ctx context.Context, in *pb.Phone) (*pb.Code, error) {
	code := sms.NewCode(in.Phone, in.Country)
	if err := code.Do(); err != nil {
		return nil, errors.New(code.ErrorCode().String())
	}
	return &pb.Code{Code: code.GetCode()}, nil
}
