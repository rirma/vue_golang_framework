package middleware

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"app/service"
)

func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
		_, is_ok := service.GetLoginUserInformation(c)

        if !is_ok {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"msg": "session invalid"})
			c.Abort()
        }
    }
}