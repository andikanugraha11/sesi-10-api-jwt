package controllers

import (
	database "github.com/andikanugraha11/rest-api-jwt/databases"
	"github.com/andikanugraha11/rest-api-jwt/helpers"
	"github.com/andikanugraha11/rest-api-jwt/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserRegistration(c *gin.Context) {
	db := database.GetDB()

	var User models.User

	ct := helpers.GetContentType(c)

	if ct == "application/json" {
		_ = c.ShouldBindJSON(&User)
	}else {
		_ = c.ShouldBind(&User)
	}

	err := db.Debug().Create(&User).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":        User.ID,
		"email":     User.Email,
		"full_name": User.FullName,
	})
}

func UserLogin(c *gin.Context) {
	db := database.GetDB()

	var User models.User

	ct := helpers.GetContentType(c)

	if ct == "application/json" {
		_ = c.ShouldBindJSON(&User)
	}else {
		_ = c.ShouldBind(&User)
	}

	password := User.Password

	err := db.Debug().Where("email = ?", User.Email).Take(&User).Error
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "Invalid email/password",
		})
		return
	}


	comparePass := helpers.ComparePassword([]byte(User.Password), []byte(password))

	if !comparePass {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "Invalid email/password",
		})
		return
	}

	token := helpers.GenerateToken(User.ID, User.Email)

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})

}
