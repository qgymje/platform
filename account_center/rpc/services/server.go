package services

import (
	"errors"
	"platform/account_center/rpc/services/email"
	"platform/account_center/rpc/services/sms"
	"platform/account_center/rpc/services/users"
	pb "platform/commons/protos/user"
	"platform/utils"

	"golang.org/x/net/context"
)

// UserServer implement pb.UserServer
type UserServer struct {
}

func srvUserToPbUser(u *users.UserInfo) *pb.UserInfo {
	userInfo := pb.UserInfo{
		UserID:     u.ID,
		Phone:      u.Phone,
		Email:      u.Email,
		Nickname:   u.Nickname,
		Token:      u.Token,
		Avatar:     u.Avatar,
		Level:      u.Level,
		FollowNum:  u.FollowNum,
		Popularity: u.Popularity,
		CreatedAt:  u.CreatedAt.Unix(),
	}
	return &userInfo
}

// Register provider register rpc call
func (s *UserServer) Register(ctx context.Context, regInfo *pb.RegisterInfo) (*pb.UserInfo, error) {
	var err error
	defer func() {
		if err != nil {
			utils.GetLog().Error("rpc.user.Register error: %+v", err)
		}
	}()

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
	return srvUserToPbUser(u), nil
}

// Auth provider auth rpc call
func (s *UserServer) Auth(ctx context.Context, token *pb.Token) (*pb.UserInfo, error) {
	var err error
	defer func() {
		if err != nil {
			utils.GetLog().Error("rpc.user.Auth error: %+v", err)
		}
	}()

	t := users.NewToken()
	if ok, err := t.Verify(token.Token); !ok && err != nil {
		return nil, errors.New(t.ErrorCode().String())
	}

	u, err := t.GetUserInfo()
	if err != nil {
		return nil, errors.New(u.ErrorCode().String())
	}
	return srvUserToPbUser(u), nil
}

// Login provider login rpc call
func (s *UserServer) Login(ctx context.Context, in *pb.LoginInfo) (*pb.UserInfo, error) {
	var err error
	defer func() {
		if err != nil {
			utils.GetLog().Error("rpc.user.Login error: %+v", err)
		}
	}()

	config := &users.LoginConfig{
		Account:  in.Account,
		Password: in.Password,
	}
	login := users.NewLogin(config)
	if err := login.Do(); err != nil {
		return nil, errors.New(login.ErrorCode().String())
	}

	u, _ := login.GetUserInfo()
	return srvUserToPbUser(u), nil
}

// Logout provider logout rpc call
func (s *UserServer) Logout(ctx context.Context, token *pb.Token) (*pb.Status, error) {
	var err error
	defer func() {
		if err != nil {
			utils.GetLog().Error("rpc.user.Logout error: %+v", err)
		}
	}()

	out := users.NewLogout(token.Token)
	if err := out.Do(); err != nil {
		return nil, errors.New(out.ErrorCode().String())
	}
	return &pb.Status{Success: true}, nil
}

// Info provider info rpc call
func (s *UserServer) Info(ctx context.Context, userID *pb.UserID) (*pb.UserInfo, error) {
	var err error
	defer func() {
		if err != nil {
			utils.GetLog().Error("rpc.user.Info error: %+v", err)
		}
	}()

	ui := users.NewUserInfo()
	err = ui.GetByID(userID.UserID)
	if err != nil {
		return nil, errors.New(ui.ErrorCode().String())
	}
	return srvUserToPbUser(ui), nil
}

// SMSCode request a sms code before register
func (s *UserServer) SMSCode(ctx context.Context, in *pb.Phone) (*pb.Code, error) {
	var err error
	defer func() {
		if err != nil {
			utils.GetLog().Error("rpc.user.SMSCode error: %+v", err)
		}
	}()

	code := sms.NewCode(in.Phone, in.Country)
	if err := code.Do(); err != nil {
		return nil, errors.New(code.ErrorCode().String())
	}
	return &pb.Code{Code: code.GetCode()}, nil
}

// EmailCode request a sms code before register
func (s *UserServer) EmailCode(ctx context.Context, in *pb.Email) (*pb.Code, error) {
	var err error
	defer func() {
		if err != nil {
			utils.GetLog().Error("rpc.user.EmailCode error: %+v", err)
		}
	}()

	code := email.NewCode(in.Email)
	if err := code.Do(); err != nil {
		return nil, errors.New(code.ErrorCode().String())
	}
	return &pb.Code{Code: code.GetCode()}, nil
}

// List query a buntch of users
func (s *UserServer) List(ctx context.Context, in *pb.UserQuery) (*pb.UsersInfo, error) {
	var err error
	defer func() {
		if err != nil {
			utils.GetLog().Error("rpc.user.List error: %+v", err)
		}
	}()

	config := &users.Config{
		PageNum:  int(in.Num),
		PageSize: int(in.Size),
		Search:   in.Search,
		IDs:      in.IDs,
	}

	users := users.NewUsers(config)
	if err = users.Do(); err != nil {
		return nil, errors.New(users.ErrorCode().String())
	}

	srvUsers := users.Users()
	count := users.Count()

	var pbUsers []*pb.UserInfo
	for _, srvUser := range srvUsers {
		pbUser := srvUserToPbUser(srvUser)
		pbUsers = append(pbUsers, pbUser)
	}

	return &pb.UsersInfo{Users: pbUsers, TotalNum: count}, nil
}
