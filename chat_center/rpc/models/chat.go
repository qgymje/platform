package models

import "time"

// Chat chat model
type Chat struct {
	ID        int64 `orm:"column(id)"`
	Name      string
	Avatar    string
	UserID    string    `orm:"column(user_id)"` // creater
	Members   []*Member `orm:"reverse(many)"`
	CreatedAt time.Time
}

// TableName table name
func (Chat) TableName() string {
	return TableNameChat
}

// Create a chat
func (c *Chat) Create() (err error) {
	_, err = GetDB().Insert(c)
	return err
}

// Find find
func (c *Chat) Find() (err error) {
	return GetDB().QueryTable(TableNameChat).RelatedSel("Members").Filter("user_id", c.UserID).One(c)
}

// FindByID one chat
func (c *Chat) FindByID() (err error) {
	return GetDB().QueryTable(TableNameChat).RelatedSel("Members").Filter("id", c.ID).One(c)
}
