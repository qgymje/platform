// Package models  data access layer
package models

import (
	"errors"
	"fmt"
	"log"
	"platform/utils"

	"github.com/astaxie/beego/orm"
	"gopkg.in/mgo.v2/bson"
)

var (
	// ErrObjectID error object id
	ErrObjectID = errors.New("not a valid objectID")
)

// ErrNotFound not found error
var ErrNotFound = orm.ErrNoRows

// DBName db name
const DBName = "chat_center"

// TableNameChat collection name
const TableNameChat = "chats"

// TableNameMember member
const TableNameMember = "members"

// TableNameMessage collection name
const TableNameMessage = "messages"

// StringToObjectID string to bson objectId
func StringToObjectID(id string) (bson.ObjectId, error) {
	if !bson.IsObjectIdHex(string(id)) {
		return bson.ObjectId(""), ErrObjectID
	}
	return bson.ObjectIdHex(id), nil
}

// StringsToObjectIDs strings to bson objectIds
func StringsToObjectIDs(ids []string) ([]bson.ObjectId, error) {
	IDHexs := []bson.ObjectId{}
	for _, id := range ids {
		if !bson.IsObjectIdHex(string(id)) {
			return nil, ErrObjectID
		}
		IDHexs = append(IDHexs, bson.ObjectIdHex(string(id)))
	}
	return IDHexs, nil
}

var db orm.Ormer

// InitModels init models
func InitModels() (err error) {
	orm.RegisterDriver("mysql", orm.DRMySQL)

	c := utils.GetConf().GetStringMapString("mysql")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset-utf8&parseTime=True&loc=Local", c["username"], c["password"], c["host"], c["port"], c["dbname"])

	err = orm.RegisterDataBase("default", "mysql", dsn)
	if err != nil {
		log.Fatalf("InitModels error: %v\n", err)
	}

	if utils.IsDev() {
		if err = createTables(); err != nil {
			log.Fatalf("create tables error: %v\n", err)
		}
	}

	db = orm.NewOrm()

	return
}

func createTables() (err error) {
	orm.Debug = true
	orm.RegisterModel(new(Chat), new(Member), new(Message))

	if err = orm.RunSyncdb("default", false, true); err != nil {
		return
	}
	return
}

// GetDB get *sqlx.DB
func GetDB() orm.Ormer {
	return db
}
