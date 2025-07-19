package routes

import (
	"example/hello/handlers"
	"example/hello/utils"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(r *gin.Engine) {
	user := r.Group("/users", LogStartEnd())
	{
		user.GET("/", handlers.GetAllUsers)
	}
}

func LogStartEnd() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		// Ambil stack trace sebelum handler dipanggil

		method := c.Request.Method
		path := c.FullPath()

		utils.InfoWithContext(c, fmt.Sprintf("Start processing %s %s", method, path))

		c.Next()

		duration := time.Since(start).Milliseconds()

		utils.InfoWithContext(c, fmt.Sprintf(
			"End processing %s %s in %dms",
			method, path , duration,
		))
	}
}
