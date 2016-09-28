package models

import (
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Game 表示一个游戏基本属性
//go:generate gen_columns -tag=bson -path=game.go
type Game struct {
	GameID      bson.ObjectId `bson:"_id"`
	GameTypeID  int           `bson:"game_type_id"`
	Name        string        `bson:"name"`
	Cover       string        `bson:"cover"`
	Screenshots []string      `bson:"screenshots"`
	Description string        `bson:"description"`
	PlayTimes   int           `bson:"play_times"`
	PlayerNum   int           `bson:"player_num"`
	IsFree      bool          `bson:"is_free"`
	Status      int           `bson:"status"`
	CreatedAt   time.Time     `bson:"created_at"`
	UpdatedAt   time.Time     `bson:"updated_at"`
	PublishedAt time.Time     `bson:"published_at"`
}

// GetID get a game id
func (g *Game) GetID() string {
	return g.GameID.Hex()
}

// Create 插入一个用户数据
func (g *Game) Create() error {
	session := GetMongo()
	defer session.Close()

	g.GameID = bson.NewObjectId()
	g.CreatedAt = time.Now()
	g.UpdatedAt = time.Now()

	return session.DB(DBName).C(ColNameGame).Insert(&g)
}

func (g *Game) update(m bson.M) error {
	session := GetMongo()
	defer session.Close()

	m[GameColumns.UpdatedAt] = time.Now()
	change := bson.M{"$set": m}
	return session.DB(DBName).C(ColNameGame).Update(bson.M{GameColumns.GameID: g.GameID}, change)
}

// Update update a game config
func (g *Game) Update(change bson.M) error {
	return g.update(change)
}

func findGame(m bson.M) (*Game, error) {
	session := GetMongo()
	defer session.Close()

	var game Game
	err := session.DB(DBName).C(ColNameGame).Find(m).One(&game)
	if err != nil {
		if err == mgo.ErrNotFound {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return &game, nil
}

func findGames(m bson.M) ([]*Game, error) {
	session := GetMongo()
	defer session.Close()

	var games []*Game
	err := session.DB(DBName).C(ColNameGame).Find(m).All(&games)
	if err != nil {
		if err == mgo.ErrNotFound {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return games, nil
}

func findGamesByStatus(status GameStatus) ([]*Game, error) {
	change := bson.M{GameColumns.Status: status}
	return findGames(change)
}

// FindGamesUnvalid just game info, no game itself
func FindGamesUnvalid() ([]*Game, error) {
	return findGamesByStatus(Unvalid)
}

// FindGamesUploaded uploaded but not valied games
func FindGamesUploaded() ([]*Game, error) {
	return findGamesByStatus(Uploaded)
}

// FindGamesValidFailed valied games
func FindGamesValidFailed() ([]*Game, error) {
	return findGamesByStatus(ValidFailed)
}

// FindGamesPublished published games
func FindGamesPublished() ([]*Game, error) {
	return findGamesByStatus(Published)
}

// FindGamesDown down games
func FindGamesDown() ([]*Game, error) {
	return findGamesByStatus(Down)
}

// FindGamesAvailable available games
func FindGamesAvailable() ([]*Game, error) {
	return findGamesByStatus(Available)
}
