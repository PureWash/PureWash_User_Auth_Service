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
	SetupRoutes(router *gin.Engine)     // Asosiy yo'llarni o'rnatish
	SetupMiddleware(router *gin.Engine) // Middlewarelarni sozlash
}

type controllerImpl struct {
	mainHandler handler.MainHandler
}

func NewController(mainHandler handler.MainHandler) Controller {
	return &controllerImpl{
		mainHandler: mainHandler,
	}
}

// @title API Gateway
// @version 1.0
// @description This is a user service.
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @schemes http
func (c *controllerImpl) SetupRoutes(router *gin.Engine) {

	// Swagger endpointini sozlash
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// User routerlarini sozlash
	router.POST("/register", c.mainHandler.User().RegisterUserHandler)
	router.POST("/login", c.mainHandler.User().LoginUserHandler)

	// user guruhlash
	user := router.Group("/users")
	user.Use(middleware.IsAuthenticated())
	{
		user.GET("/profile", c.mainHandler.User().GetUserHandler)
		user.PUT("/update", c.mainHandler.User().UpdateUserHandler)
		user.DELETE("/delete", c.mainHandler.User().DeleteUserHandler)
		user.PUT("/password", c.mainHandler.User().UpdatePassword)
	}

}

func (c *controllerImpl) SetupMiddleware(router *gin.Engine) {
	router.Use(gin.Logger())
	router.Use(middleware.CorsMiddileware())
}
