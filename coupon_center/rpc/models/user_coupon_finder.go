package models

import (
	"math"

	"gopkg.in/mgo.v2/bson"
)

// UserCouponFinder user coupon finder
type UserCouponFinder struct {
	skip, limit int
	userID      bson.ObjectId
	err         error
	where       bson.M
	userCoupons []*UserCoupon
}

// NewUserCouponFinder new user coupon finder
func NewUserCouponFinder() *UserCouponFinder {
	f := new(UserCouponFinder)
	f.where = bson.M{}
	f.userCoupons = []*UserCoupon{}
	return f
}

// Limit limit
func (c *UserCouponFinder) Limit(offset, limit int) *UserCouponFinder {
	c.skip = int(math.Max(0, float64(offset-1))) * limit
	c.limit = limit
	return c

}

// ByUserID by user id
func (c *UserCouponFinder) ByUserID(userID string) *UserCouponFinder {
	c.userID, c.err = StringToObjectID(userID)
	if c.err == nil {
		c.where[UserCouponColumns.UserID] = c.userID
	}
	return c
}

// Do do the query work
func (c *UserCouponFinder) Do() (err error) {
	session := GetMongo()
	defer session.Close()

	err = session.DB(DBName).C(ColNameUserCoupon).Find(c.where).Skip(c.skip).Limit(c.limit).All(&c.userCoupons)
	if err != nil {
		return err
	}

	if len(c.userCoupons) == 0 {
		return ErrNotFound
	}

	return nil
}

// Result result
func (c *UserCouponFinder) Result() []*UserCoupon {
	return c.userCoupons
}

// One one record
func (c *UserCouponFinder) One() *UserCoupon {
	return c.userCoupons[0]
}

// Count returns the total number of query
func (c *UserCouponFinder) Count() int64 {
	session := GetMongo()
	defer session.Close()

	n, err := session.DB(DBName).C(ColNameUserCoupon).Find(c.where).Count()
	if err != nil {
		return 0
	}
	return int64(n)
}
