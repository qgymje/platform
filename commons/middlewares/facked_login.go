package middlewares

import "github.com/gin-gonic/gin"

// user 1
const token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE0Nzg4NDQ2MjcsImlkIjoiNTdlMjI2N2VjODZhYjQ1YWYzZDE0ODA2In0.QEfA_RFWmSGpDU1bT_Gzdh0Jl06ft8oXBNntB5l3CEU"

func FakedLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Request.Header.Add("Authorization", "bearer "+token)
		c.Next()
	}
}
