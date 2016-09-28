package middlewares

import (
	"time"

	"github.com/gin-gonic/gin"
)

func RecordRequestBegin() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("request_begin", time.Now())
		c.Next()
	}
}
