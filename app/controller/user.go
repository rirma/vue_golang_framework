package controller

import(
	"github.com/gin-gonic/gin"
	"net/http"
	"app/service"
)

func UserInformation(c *gin.Context) {
	user, is_ok := service.GetLoginUserInformation(c)
	if !is_ok {
		c.JSON(http.StatusOK, gin.H{
			"is_valid": false,
			"name": "",
			"email": "",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"is_valid": true,
			"name": user.Name,
			"email": user.Email,
		})
	}
}
