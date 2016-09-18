package users

import (
	"crypto/md5"
	"errors"
	"fmt"

	"platform/utils"
)

var (
	// should remove for security issuse
	ErrPasswordTooShort = errors.New("密码不能小于6位")
)

const (
	SaltLength           = 6
	PasswordMin          = 6
	RandomPasswordLength = 32
)

const letterBytes = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func generateRandomString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[utils.GetRand().Intn(len(letterBytes))]
	}
	return string(b)
}

type Password struct {
	pwd          string
	salt         string
	encryptedPwd string
}

func NewPassword(pwd string) *Password {
	return &Password{
		pwd: pwd,
	}
}

func (p *Password) Len() int {
	return len(p.pwd)
}

func (p *Password) IsConfirmSame(pwd string) bool {
	return p.pwd == pwd
}

func (p *Password) SetSalt(salt string) {
	p.salt = salt
}

func (p *Password) GenSalt() string {
	salt := generateRandomString(SaltLength)
	if p.salt == "" {
		p.salt = salt
	}
	return salt
}

func (p *Password) GenPwd() string {
	passSalt := []byte(p.pwd + p.salt)
	p.encryptedPwd = fmt.Sprintf("%x", md5.Sum(passSalt))
	return p.encryptedPwd
}

// 如果是第三方注册的, 生成一个随机的密码
func (p *Password) GenRandwomPwd() string {
	return generateRandomString(RandomPasswordLength)
}

func (p *Password) IsEncryptedSame(encryptedPwd string) bool {
	return p.Encryped() == encryptedPwd
}

func (p *Password) Encryped() string {
	if p.encryptedPwd == "" {
		p.encryptedPwd = p.GenPwd()
	}
	return p.encryptedPwd
}

func (p *Password) Valid() error {
	if p.Len() < PasswordMin {
		return ErrPasswordTooShort
	}

	return nil
}

func (p *Password) Salt() string {
	if p.salt == "" {
		p.salt = p.GenSalt()
	}
	return p.salt
}
