package controllers

import "github.com/gin-gonic/gin"

// Game game controller
type Game struct {
	Base
}

// List game list
func (g *Game) List(c *gin.Context) {

}

// Start user start a game
func (g *Game) Start(c *gin.Context) {

}

// UpdatePreference when a user update a preference of a game
func (g *Game) UpdatePreference(c *gin.Context) {

}

// Preference fetch a preference by user_id and game_id
func (g *Game) Preference(c *gin.Context) {

}
