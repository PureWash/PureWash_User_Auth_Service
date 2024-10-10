package router

import (
	"user-service/api/handler"
	"user-service/api/middleware"

	_ "user-service/api/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Controller interface {
	SetupRoutes()     // Asosiy yo'llarni o'rnatish
}

type controllerImpl struct {
	mainHandler handler.MainHandler
	router      *gin.Engine
}

func NewController(mainHandler handler.MainHandler, router *gin.Engine) Controller {
	return &controllerImpl{
		mainHandler: mainHandler,
		router:      router,
	}
}

// @title API Gateway
// @version 1.0
// @description This is a user service.
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @schemes http
func (c *controllerImpl) SetupRoutes() {
	c.router.Use(middleware.CorsMiddileware())
	// Swagger endpointini sozlash
	c.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// User routerlarini sozlash
	c.router.POST("/register", c.mainHandler.User().RegisterUserHandler)
	c.router.POST("/login", c.mainHandler.User().LoginUserHandler)
	c.router.POST("/refresh-token", c.mainHandler.User().UpdateAccessTokenHandler)

	// user guruhlash
	user := c.router.Group("/users")
	user.Use(middleware.IsAuthenticated())
	{
		user.GET("/profile", c.mainHandler.User().GetUserHandler)
		user.PUT("/update", c.mainHandler.User().UpdateUserHandler)
		user.DELETE("/delete", c.mainHandler.User().DeleteUserHandler)
		user.PUT("/password", c.mainHandler.User().UpdatePassword)
	}

}
