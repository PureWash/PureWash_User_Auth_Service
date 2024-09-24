package handler

import (
	"log/slog"
	userHandler "user-service/api/handler/user"
	"user-service/internal/service"
)


type MainHandler interface {
	User() userHandler.UserHandler
}


type mainHandlerImpl struct {
	userService service.UserService
	logger *slog.Logger
}

func NewMainHandler(userService service.UserService, logger *slog.Logger) MainHandler {
	return &mainHandlerImpl{
		userService: userService,
		logger: logger,
	}
}

func (mh *mainHandlerImpl) User() userHandler.UserHandler {
	return userHandler.NewUserHandler(mh.userService, mh.logger)
}