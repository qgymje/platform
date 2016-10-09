package models

import (
	"math"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// GameFinder an object find games
type GameFinder struct {
	status      int
	gameTypeID  int
	search      string
	skip, limit int
	order       string

	query *mgo.Query
	games *[]Game
}

// NewGameFinder create a new GameFinder
func NewGameFinder() *GameFinder {
	f := new(GameFinder)
	f.games = &[]Game{}
	return f
}

// Limit limit
func (g *GameFinder) Limit(offset, limit int) *GameFinder {
	g.skip = int(math.Max(0, float64(offset-1))) * limit
	g.limit = limit
	return g
}

// Status status
func (g *GameFinder) Status(s GameStatus) *GameFinder {
	g.status = int(s)
	return g
}

// Order order
func (g *GameFinder) Order(o string) *GameFinder {
	g.order = o
	return g
}

// GameTypeID game type id
func (g *GameFinder) GameTypeID(id int) *GameFinder {
	g.gameTypeID = id
	return g
}

// Search search
func (g *GameFinder) Search(search string) *GameFinder {
	g.search = search
	return g
}

func (g *GameFinder) condition() bson.M {
	where := bson.M{}

	if g.search != "" {
		where["$text"] = bson.M{"$search": g.search}
	}

	if g.gameTypeID > 0 {
		where[GameColumns.GameTypeID] = g.gameTypeID
	}
	where[GameColumns.Status] = g.status

	return where
}

// Do do the search job
func (g *GameFinder) Do() (err error) {
	session := GetMongo()
	defer session.Close()

	err = session.DB(DBName).C(ColNameGame).Find(g.condition()).Skip(g.skip).Limit(g.limit).All(g.games)
	if err != nil {
		if err == mgo.ErrNotFound {
			return ErrNotFound
		}
		return err
	}
	return nil
}

// Result return the games that found
func (g *GameFinder) Result() *[]Game {
	return g.games
}

// Count return the total num of the query
func (g *GameFinder) Count() int64 {
	session := GetMongo()
	defer session.Close()

	n, err := session.DB(DBName).C(ColNameGame).Find(g.condition()).Count()
	if err != nil {
		return 0
	}
	return int64(n)
}
