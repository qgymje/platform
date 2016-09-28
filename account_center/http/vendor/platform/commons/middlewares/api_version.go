package middlewares

import "github.com/gin-gonic/gin"

func APIVersion() gin.HandlerFunc {
	return func(c *gin.Context) {
		v := c.Request.Header.Get("Version")
		c.Set("version", v)
		c.Next()
	}
}
