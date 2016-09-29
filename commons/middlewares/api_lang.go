package middlewares

import "github.com/gin-gonic/gin"

// APILang accepts header version
func APILang() gin.HandlerFunc {
	return func(c *gin.Context) {
		v := c.Request.Header.Get("Accept-Language")
		c.Set("lang", v)
		c.Next()
	}
}
