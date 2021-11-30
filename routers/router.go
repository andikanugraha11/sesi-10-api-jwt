package routers

import (
	"github.com/andikanugraha11/rest-api-jwt/controllers"
	"github.com/andikanugraha11/rest-api-jwt/middlewares"
	"github.com/gin-gonic/gin"
)

func InitApplication() *gin.Engine  {
	r := gin.Default()
	// user/register
	// user/login
	userGroup := r.Group("/user")
	userGroup.POST("/register", controllers.UserRegistration)
	userGroup.POST("/login", controllers.UserLogin)


	productGroup := r.Group("/product")
	productGroup.Use(middlewares.Authentication())
	productGroup.POST("/", controllers.CreateProduct)
	productGroup.POST("/:productId", middlewares.Authorization(), controllers.UpdateProduct)


	return r
}