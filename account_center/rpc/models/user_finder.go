package models

import (
	"math"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// UserFinder user model finder
type UserFinder struct {
	skip, limit int
	search      string
	ids         []bson.ObjectId
	err         error
	users       []*User
}

// NewUserFinder create a new UserFinder
func NewUserFinder() *UserFinder {
	f := new(UserFinder)
	f.users = []*User{}
	return f
}

// Limit limit
func (u *UserFinder) Limit(offset, limit int) *UserFinder {
	u.skip = int(math.Max(0, float64(offset-1))) * limit
	u.limit = limit
	return u
}

// Search search
func (u *UserFinder) Search(search string) *UserFinder {
	u.search = search
	return u
}

// ByIDs by user ids
func (u *UserFinder) ByIDs(ids []string) *UserFinder {
	u.ids, u.err = StringsToObjectIDs(ids)
	u.skip = 0
	u.limit = len(u.ids)
	return u
}

func (u *UserFinder) condition() bson.M {
	where := bson.M{}

	if u.search != "" {
		where["$text"] = bson.M{"$search": u.search}
	}

	if len(u.ids) > 0 {
		where[UserColumns.ID] = bson.M{"$in": u.ids}
	}
	return where
}

// Do the query work
func (u *UserFinder) Do() (err error) {
	session := GetMongo()
	defer session.Close()

	err = session.DB(DBName).C(ColNameUser).Find(u.condition()).Skip(u.skip).Limit(u.limit).All(&u.users)
	if err != nil {
		if err == mgo.ErrNotFound {
			return ErrNotFound
		}
		return err
	}
	return nil
}

// Result return the users that found
func (u *UserFinder) Result() []*User {
	return u.users
}

// Count return the total number under the condition
func (u *UserFinder) Count() int64 {
	session := GetMongo()
	defer session.Close()

	n, err := session.DB(DBName).C(ColNameUser).Find(u.condition()).Count()
	if err != nil {
		return 0
	}
	return int64(n)
}
