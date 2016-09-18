package middlewares

import "github.com/gin-gonic/gin"

// user 1
const token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE0NzYyNTUwNDYsImlkIjoiNTdjZThmN2VjODZhYjRkMjA0ZjExZDI3In0.XOvgzEDAtL8j-DjYjq0VqZIFwgu7uyPUJ6Oljr9kn7M"

func FakedLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Request.Header.Add("Authorization", "bearer "+token)
		c.Next()
	}
}
