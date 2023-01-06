package controller

import(
	"github.com/gin-gonic/gin"
	"net/http"
	"app/service"
	"fmt"
)

type LoginRequest struct {
	Email    string `form:"id" json:"id" binding:"required"`
	Password string `form:"pass" json:"pass" binding:"required"`
}

func Login(c *gin.Context) {
	var request LoginRequest
	if request_err := c.ShouldBindJSON(&request); request_err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": request_err.Error()})
		return
	}

    email := request.Email
	pass := request.Password

	loginService := service.LoginService{}
	result, err := loginService.Login(c, email, pass)

	if !result {
		fmt.Println(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"is_ok": true,
	})
}

func Logout(c *gin.Context) {
	loginService := service.LoginService{}
	result := loginService.Logout(c)

	if !result {
		c.JSON(http.StatusBadRequest, gin.H{})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"is_ok": true,
		})
	}
}