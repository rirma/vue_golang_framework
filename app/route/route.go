package route

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"app/controller"
	"app/middleware"
)

var router *gin.Engine

func init() {
	router = gin.Default()
	// CORS for https://foo.com and https://github.com origins, allowing:
	// - PUT and PATCH methods
	// - Origin header
	// - Credentials share
	// - Preflight requests cached for 12 hours
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:8080"},
		AllowMethods: []string{
			"GET",
			"POST",
		},
		AllowHeaders: []string{
			"X-Requested-With",
			"Origin",
			"X-Csrftoken",
			"Content-Type",
			"Accept",
			"X-XSRF-TOKEN",
		},
		ExposeHeaders: []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge: 12 * time.Hour,
	}))
}

func Routing() {
	api := router.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			v1.POST("/login", controller.Login)
			v1.POST("/logout", controller.Logout)
			auth := v1.Group("/auth")
			auth.Use(middleware.AuthMiddleware())
			{
				auth.GET("/user_information", controller.UserInformation)
			}
		}
	}

	router.Run(":8082")
}
