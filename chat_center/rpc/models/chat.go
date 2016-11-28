package models

import (
	"crypto/md5"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"
)

// Chat chat model
type Chat struct {
	ID        int64 `orm:"column(id)"`
	Name      string
	Avatar    string
	UserID    string    `orm:"column(user_id)"` // creater
	Sign      string    `orm:"size(32)"`        // ordered user_id include owner id in md5
	Members   []*Member `orm:"reverse(many)"`
	CreatedAt time.Time
}

// TableName table name
func (Chat) TableName() string {
	return TableNameChat
}

// GetID get chat id
func (c *Chat) GetID() string {
	return strconv.FormatInt(c.ID, 10)
}

// Create a chat
func (c *Chat) Create(userIDs []string) (err error) {
	c.Sign = c.GenSign(userIDs)
	c.CreatedAt = time.Now()

	GetDB().Begin()
	if _, err = GetDB().Insert(c); err != nil {
		GetDB().Rollback()
		return err
	}

	members := c.genMembers(userIDs)
	if err = CreateMembers(members); err != nil {
		GetDB().Rollback()
		return err
	}

	GetDB().Commit()
	return err
}

func (c *Chat) genMembers(userIDs []string) (members []*Member) {
	now := time.Now()
	for i := range userIDs {
		member := &Member{
			Chat:      c,
			UserID:    userIDs[i],
			CreatedAt: now,
		}
		members = append(members, member)
	}
	return
}

func (c *Chat) membersToStrings() []string {
	s := []string{}
	for _, m := range c.Members {
		s = append(s, m.GetUserID())
	}
	return s
}

// GenSign generate sign
func (c *Chat) GenSign(ss []string) string {
	sort.Strings(ss)
	s := strings.Join(ss, "*")
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

// GetSign get sing
func (c *Chat) GetSign() string {
	ss := c.membersToStrings()
	return c.GenSign(ss)
}

// Find find
func (c *Chat) Find() (err error) {
	if err = GetDB().QueryTable(TableNameChat).Filter("user_id", c.UserID).One(c); err != nil {
		return
	}
	if c.withMembers(); err != nil {
		return
	}
	return
}

// FindByID one chat
func (c *Chat) FindByID() (err error) {
	if err = GetDB().QueryTable(TableNameChat).Filter("id", c.ID).One(c); err != nil {
		return
	}
	if c.withMembers(); err != nil {
		return
	}
	return
}

// FindBySign find by sign
func (c *Chat) FindBySign() (err error) {
	if err = GetDB().QueryTable(TableNameChat).Filter("sign", c.Sign).One(c); err != nil {
		return
	}
	if c.withMembers(); err != nil {
		return
	}
	return
}

func (c *Chat) withMembers() (err error) {
	_, err = GetDB().QueryTable(TableNameMember).Filter("Chat__id", c.ID).RelatedSel().All(&c.Members)
	return
}
