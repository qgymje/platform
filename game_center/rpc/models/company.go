package models

import (
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Company  company info
//go:generate gen_columns -tag=bson -path=./company.go
type Company struct {
	ID        bson.ObjectId `bson:"_id"`
	UserID    bson.ObjectId `bson:"userID"`
	Name      string        `bson:"name"`
	Loactions []string      `bson:"locations"`
	Valid     bool          `bson:"valid"`
	CreatedAt time.Time     `bson:"created_at"`
	UpdatedAt time.Time     `bson:"updated_at"`
}

// Create  create a company info
func (c *Company) Create() error {
	session := GetMongo()
	defer session.Close()

	c.ID = bson.NewObjectId()
	c.CreatedAt = time.Now()
	c.UpdatedAt = time.Now()

	return session.DB(DBName).C(ColNameCompany).Insert(&c)
}

func findCompanies(m bson.M) ([]*Company, error) {
	session := GetMongo()
	defer session.Close()

	var companies []*Company
	err := session.DB(DBName).C(ColNameCompany).Find(m).All(&companies)
	if err != nil {
		if err == mgo.ErrNotFound {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return companies, nil
}

func findCompany(m bson.M) (*Company, error) {
	session := GetMongo()
	defer session.Close()

	var company Company
	err := session.DB(DBName).C(ColNameCompany).Find(m).All(&company)
	if err != nil {
		if err == mgo.ErrNotFound {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return &company, nil
}
