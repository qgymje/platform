package models

import (
	"math"

	"gopkg.in/mgo.v2/bson"
)

// CouponFinder coupon finder
type CouponFinder struct {
	skip, limit int
	campaignID  bson.ObjectId
	ids         []bson.ObjectId
	where       bson.M
	coupons     []*Coupon

	err error
}

// NewCouponFinder new cupon finder
func NewCouponFinder() *CouponFinder {
	f := new(CouponFinder)
	f.where = bson.M{}
	f.coupons = []*Coupon{}
	return f
}

// ByIDs by coupon id
func (c *CouponFinder) ByIDs(couponIDs []string) *CouponFinder {
	c.ids, c.err = StringsToObjectIDs(couponIDs)
	if c.err == nil {
		c.where[CouponColumns.CouponID] = bson.M{"$in": c.ids}
	}
	return c
}

// ByCampaignID by campain id
func (c *CouponFinder) ByCampaignID(campaignID string) *CouponFinder {
	c.campaignID, c.err = StringToObjectID(campaignID)
	if c.err == nil {
		c.where[CouponColumns.CampaignID] = c.campaignID
	}
	return c
}

// Limit limit
func (c *CouponFinder) Limit(offset, limit int) *CouponFinder {
	c.skip = int(math.Max(0, float64(offset-1))) * limit
	c.limit = limit
	return c
}

// Do do the search job
func (c *CouponFinder) Do() (err error) {
	session := GetMongo()
	defer session.Close()

	err = session.DB(DBName).C(ColNameCoupon).Find(c.where).Skip(c.skip).Limit(c.limit).All(&c.coupons)
	if err != nil {
		return err
	}

	if len(c.coupons) == 0 {
		return ErrNotFound
	}
	return nil
}

// Result result
func (c *CouponFinder) Result() []*Coupon {
	return c.coupons
}

// Count count
func (c *CouponFinder) Count() int64 {
	session := GetMongo()
	defer session.Close()

	n, err := session.DB(DBName).C(ColNameCoupon).Find(c.where).Count()
	if err != nil {
		return 0
	}
	return int64(n)
}
