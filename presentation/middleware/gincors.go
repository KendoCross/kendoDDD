package middleware

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func GinCors() gin.HandlerFunc {
	config := cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization", "Access-Control-Allow-Origin"},
		AllowCredentials: true,
		AllowOrigins:     []string{"*"},
		AllowOriginFunc: func(str string) bool {
			return true
		},
		MaxAge: 12 * time.Hour,
	}
	//config.AllowAllOrigins = true
	return cors.New(config)
}
