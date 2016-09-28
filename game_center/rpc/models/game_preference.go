package models

import (
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// GamePreference  game preference of each user each game
//go:generate gen_columns -tag=bson -path=game_preference.go
type GamePreference struct {
	UserID     bson.ObjectId          `bson:"user_id"`
	GameID     bson.ObjectId          `bson:"game_id"`
	Preference map[string]interface{} `bson:"preference"`
	CreatedAt  time.Time              `bson:"created_at"`
	UpdatedAt  time.Time              `bson:"updated_at"`
}

// Create create a game preference each game each user
func (g *GamePreference) Create() error {
	session := GetMongo()
	defer session.Close()

	g.CreatedAt = time.Now()
	g.UpdatedAt = time.Now()

	return session.DB(DBName).C(ColNameGamePreference).Insert(&g)
}

func (g *GamePreference) update(m bson.M) error {
	session := GetMongo()
	defer session.Close()

	m[GameColumns.UpdatedAt] = time.Now()
	change := bson.M{"$set": m}
	return session.DB(DBName).C(ColNameGamePreference).Update(bson.M{GamePreferenceColumns.GameID: g.GameID, GamePreferenceColumns.UserID: g.UserID}, change)
}

// Update game preference
func (g *GamePreference) Update(change bson.M) error {
	return g.update(change)
}

func findGamePreference(m bson.M) (*GamePreference, error) {
	session := GetMongo()
	defer session.Close()

	var game GamePreference
	err := session.DB(DBName).C(ColNameGamePreference).Find(m).One(&game)
	if err != nil {
		if err == mgo.ErrNotFound {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return &game, nil
}
