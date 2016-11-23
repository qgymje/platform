package middlewares

import "github.com/gin-gonic/gin"

const (
	authKey     = "Authorization"
	tokenPrefix = "bearer "
	token       = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE0ODI0Njc3MzcsImlkIjoiNTdlMjI2N2VjODZhYjQ1YWYzZDE0ODA2In0.LrJZD29FyAgyz13Ugn-mpEWGlim_edE8HyzH0-h7TCg"
)

// FakedLogin add a facked user token to the header
func FakedLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get(authKey)
		if authHeader == "" {
			c.Request.Header.Add(authKey, tokenPrefix+token)
		}
		c.Next()
	}
}
