package controllers

import (
	"net/http"
	"platform/commons/codes"

	"github.com/gin-gonic/gin"
)

// Game game controller
type Game struct {
	Base
}

// Types game types
func (g *Game) Types(c *gin.Context) {
	data := map[string]interface{}{
		"list": gameTypes,
	}
	respformat := g.Response(c, codes.ErrorCodeSuccess, data)
	c.JSON(http.StatusOK, respformat)
	return
}

// GameInfo game info
type GameInfo struct {
	GameID       string   `json:"gameID"`
	Name         string   `json:"name"`
	GameTypeID   int      `json:"gameTypeID"`
	GameTypeName string   `json:"gameTypeName"`
	Description  string   `json:"description"`
	Cover        string   `json:"cover"`
	Screenshots  []string `json:"screenshots"`
	PlayTimes    int      `json:"playTimes"`
	PlayerNum    int      `json:"playerNum"`
	IsFree       bool     `json:"isFree"`
	PayStatus    bool     `json:"payStatus"`
}

var games []GameInfo

type gameType struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var gameTypes []gameType

func init() {
	gameTypes = []gameType{
		{1, "动作游戏"},
		{2, "冒险游戏"},
		{3, "街机游戏"},
		{4, "桌面游戏"},
		{5, "卡牌游戏"},
		{6, "娱乐场游戏"},
		{7, "休闲游戏"},
		{8, "教育游戏"},
		{9, "音乐游戏"},
		{10, "解谜游戏"},
		{11, "竞速游戏"},
		{12, "角色扮演"},
		{13, "模拟游戏"},
		{14, "体育游戏"},
		{15, "策略游戏"},
		{16, "文字游戏"},
	}

	games = []GameInfo{
		{
			GameID:       "123456789",
			Name:         "world of warcraft",
			GameTypeID:   12,
			GameTypeName: "角色扮演",
			Description:  "Wow is a grate game in history",
			Cover:        "http://img4.imgtn.bdimg.com/it/u=2977877931,2204864369&fm=21&gp=0.jpg",
			Screenshots: []string{
				"http://img2.imgtn.bdimg.com/it/u=1803873670,2284693330&fm=11&gp=0.jpg",
				"http://img0.imgtn.bdimg.com/it/u=3951912182,2498656724&fm=11&gp=0.jpg",
			},
			PlayerNum: 1,
			PlayTimes: 10000,
			IsFree:    false,
			PayStatus: true,
		},
		{
			GameID:       "123456788",
			Name:         "league of legends",
			GameTypeID:   1,
			GameTypeName: "动作游戏",
			Description:  "Lol is a grate game in history",
			Cover:        "https://encrypted-tbn3.gstatic.com/images?q=tbn:ANd9GcT9FXQ0uqQXlxFsd16A3YKLxCxJkDSqwLZWTBoh6psWjVR-KHkL",
			Screenshots: []string{
				"https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRJYt-xRPgp1JlWG0NzHHVDCetE-CWJfIAiW97NJ3WUamOT3QR2",
				"https://encrypted-tbn1.gstatic.com/images?q=tbn:ANd9GcTZq5K7qq9zyjvyabULU9gkREaXl85CHTA-qAmnlaWnbvHmaj6g",
			},
			PlayerNum: 10,
			PlayTimes: 20000,
			IsFree:    true,
			PayStatus: false,
		},
	}
}

// List game list
func (g *Game) List(c *gin.Context) {
	data := map[string]interface{}{
		"list":      games,
		"pageNum":   1,
		"pageSize":  20,
		"totalPage": 5,
	}

	respformat := g.Response(c, codes.ErrorCodeSuccess, data)
	c.JSON(http.StatusOK, respformat)
	return
}

// Search search games
func (g *Game) Search(c *gin.Context) {
	data := map[string]interface{}{
		"list":      games,
		"query":     g.getQuery(c),
		"pageNum":   1,
		"pageSize":  20,
		"totalPage": 5,
	}

	respformat := g.Response(c, codes.ErrorCodeSuccess, data)
	c.JSON(http.StatusOK, respformat)
	return

}

// GameVM game vm info
type GameVM struct {
	IP   string `json:"ip"`
	Port int    `json:"port"`
}

// Start user start a game
func (g *Game) Start(c *gin.Context) {
	gameVM := GameVM{
		IP:   "220.181.57.217",
		Port: 10001,
	}

	respformat := g.Response(c, codes.ErrorCodeSuccess, gameVM)
	c.JSON(http.StatusOK, respformat)
	return
}

// End user quit a game
func (g *Game) End(c *gin.Context) {
	status := struct {
		Success bool `json:"success"`
	}{
		Success: true,
	}
	respformat := g.Response(c, codes.ErrorCodeSuccess, status)
	c.JSON(http.StatusOK, respformat)
	return

}

// UpdatePreference when a user update a preference of a game
func (g *Game) UpdatePreference(c *gin.Context) {

}

// Preference fetch a preference by user_id and game_id
func (g *Game) Preference(c *gin.Context) {

}
