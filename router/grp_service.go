package router

import (
	"github.com/gin-gonic/contrib/gzip"
	"github.com/gin-gonic/gin"
	message "github.com/grpc_server_demo/controllers/message"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token")
		c.Writer.Header().Set("Content-Encoding", "gzip")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
		}
	}
}

func GRPCServiceRouter(r *gin.Engine) {

	r.Use(CORSMiddleware())
	r.Use(gzip.Gzip(gzip.DefaultCompression))
	v1 := r.Group("api/v1/")
	{
		cr := v1.Group("message")
		{
			cr.GET("/:message", message.GetMessages)
		}
	}
}
