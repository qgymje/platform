// Package models 表示数据操作层, 不涉及任何具体的业务逻辑, 只做数据操作, 保持功能唯一
package models

import mgo "gopkg.in/mgo.v2"

var mongoSession *mgo.Session
var ErrNotFound = mgo.ErrNotFound

const DBName = "accountCenter"
const ColNameUser = "users"
const ColNameUserLogin = "userLogins"

// InitMongodb 根据配置初始化mongodb
func InitMongodb(sess *mgo.Session) {
	mongoSession = sess
}

// GetMongo 封装一个session, 在使用后需要Close()
func GetMongo() *mgo.Session {
	return mongoSession.Copy()
}
