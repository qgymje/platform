package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Company 开发商信息管理
type Company struct {
	ID        bson.ObjectId `bson:"_id"`
	UserID    bson.ObjectId `bson:"userID"`
	Name      string        `bson:"name"`
	Loactions []string      `bson:"locations"`
	Valid     bool          `bson:"valid"`
	CreatedAt time.Time     `bson:"created_at"`
	UpdatedAt time.Time     `bson:"updated_at"`
}

// Create 插入一个用户数据
func (c *Company) Create() error {
	session := GetMongo()
	defer session.Close()

	c.ID = bson.NewObjectId()
	c.CreatedAt = time.Now()
	c.UpdatedAt = time.Now()

	return session.DB(DBName).C(ColNameCompany).Insert(&c)
}
