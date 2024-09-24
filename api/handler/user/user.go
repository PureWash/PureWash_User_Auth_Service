package userHandler

import (
	"fmt"
	"log/slog"
	"user-service/internal/models"
	"user-service/internal/service"

	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	RegisterUserHandler(ctx *gin.Context)
	LoginUserHandler(ctx *gin.Context)
	GetUserHandler(ctx *gin.Context)
	DeleteUserHandler(ctx *gin.Context)
	UpdateUserHandler(ctx *gin.Context)
	UpdatePassword(ctx *gin.Context)
}

type userHandlerImpl struct {
	userService service.UserService
	logger      *slog.Logger
}

func NewUserHandler(userService service.UserService, logger *slog.Logger) UserHandler {
	return &userHandlerImpl{
		userService: userService,
		logger: logger,
	}
} 

// @Summary Register a new user
// @Description This endpoint registers a new user with provided details
// @Tags auth
// @Accept json
// @Produce json
// @Param user body models.UserRegisterRequst true "User Register Request"
// @Success 201 {object} models.UserRegisterResponce
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /register [post]
func (uh *userHandlerImpl) RegisterUserHandler(ctx *gin.Context) {
	uh.logger.Info(fmt.Sprint("User register methods"))
	var user models.UserRegisterRequst

	// JSON binding error
	if err := ctx.BindJSON(&user); err != nil {
		uh.logger.Error(fmt.Sprintf("Error in user binding json: %v", err))
		
		// Return detailed error response
		ctx.JSON(400, models.ErrorResponse{
			Status:  400,
			Message: "Invalid input",
			Error:   err.Error(),
		})
		return
	}

	// Service layer error (user creation failed)
	resp, err := uh.userService.RegisterUser(ctx, user)
	if err != nil {
		uh.logger.Error(fmt.Sprintf("Error creating user: %v", err))
		
		ctx.JSON(500, models.ErrorResponse{
			Status:  500,
			Message: "Failed to create user",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(201, resp)
}


// @Summary User login
// @Description This endpoint allows a user to log in by providing login credentials
// @Tags auth
// @Accept json
// @Produce json
// @Param login body models.LoginRequest true "Login Request"
// @Success 200 {object} models.LoginResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /login [post]
func (uh *userHandlerImpl) LoginUserHandler(ctx *gin.Context) {
	uh.logger.Info("User login methods")
	var login models.LoginRequest

	// JSON binding error
	if err := ctx.BindJSON(&login); err != nil {
		uh.logger.Error(fmt.Sprintf("Error in binding json: %v", err))
		
		ctx.JSON(400, models.ErrorResponse{
			Status:  400,
			Message: "Invalid login input",
			Error:   err.Error(),
		})
		return
	}

	// Service layer error (login failed)
	resp, err := uh.userService.LoginUser(ctx, login)
	if err != nil {
		uh.logger.Error(fmt.Sprintf("Error in server login user: %v", err))
		
		ctx.JSON(500, models.ErrorResponse{
			Status:  500,
			Message: "Failed to login",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(200, resp)
}


// @Summary Delete a user
// @Description This endpoint deletes a user by their ID
// @Tags users
// @Produce json
// @Param id query string true "User ID"
// @Success 200 {object} models.SuccessResponce
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /users/delete [delete]
func (uh *userHandlerImpl) DeleteUserHandler(ctx *gin.Context) {
	uh.logger.Info("User delete methods")
	var id = ctx.Query("id")

	// Service layer error (delete user failed)
	err := uh.userService.DeleteUser(ctx, id)
	if err != nil {
		uh.logger.Error(fmt.Sprintf("Error in delete user: %v", err))
		
		ctx.JSON(500, models.ErrorResponse{
			Status:  500,
			Message: "Failed to delete user",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(200, models.SuccessResponce{
		Status:  200,
		Message: "User deleted successfully",
	})
}


// @Summary Update user profile
// @Description This endpoint updates user profile details
// @Tags users
// @Accept json
// @Produce json
// @Param user body models.UpdateUserProfile true "User Profile"
// @Success 200 {object} models.SuccessResponce
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /users/update [put]
func (uh *userHandlerImpl) UpdateUserHandler(ctx *gin.Context) {
	uh.logger.Info("User update method")
	var user models.UpdateUserProfile

	// JSON binding error
	if err := ctx.BindJSON(&user); err != nil {
		uh.logger.Error(fmt.Sprintf("Error binding JSON: %v", err))
		ctx.JSON(400, models.ErrorResponse{
			Status:  400,
			Message: "Invalid user input",
			Error:   err.Error(),
		})
		return
	}

	// Service layer error
	err := uh.userService.UpdateUserProfile(ctx, user)
	if err != nil {
		uh.logger.Error(fmt.Sprintf("Error updating user profile: %v", err))
		ctx.JSON(500, models.ErrorResponse{
			Status:  500,
			Message: "Failed to update user profile",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(200, models.SuccessResponce{
		Status:  200,
		Message: "User profile updated successfully",
	})
}

// @Summary Get user profile
// @Description This endpoint retrieves the user profile by their ID
// @Tags users
// @Produce json
// @Param id query string true "User ID"
// @Success 200 {object} models.UserProfile
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /users/profile [get]
func (uh *userHandlerImpl) GetUserHandler(ctx *gin.Context) {
	uh.logger.Info("User profile get method")
	id := ctx.Query("id")

	// Service layer error
	resp, err := uh.userService.GetUserProfile(ctx, id)
	if err != nil {
		uh.logger.Error(fmt.Sprintf("Error retrieving user profile: %v", err))
		ctx.JSON(500, models.ErrorResponse{
			Status:  500,
			Message: "Failed to get user profile",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(200, resp)
}


// @Summary Update password
// @Description This endpoint updates the user's password
// @Tags users
// @Accept json
// @Produce json
// @Param updatePass body models.UpdatePasswordRequest true "Update Password Request"
// @Success 200 {object} models.SuccessResponce
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /users/password [put]
func (uh *userHandlerImpl) UpdatePassword(ctx *gin.Context) {
	uh.logger.Info("Update password method")
	var updatePass models.UpdatePasswordRequest

	// JSON binding error
	if err := ctx.BindJSON(&updatePass); err != nil {
		uh.logger.Error(fmt.Sprintf("Error binding JSON: %v", err))
		ctx.JSON(400, models.ErrorResponse{
			Status:  400,
			Message: "Invalid input",
			Error:   err.Error(),
		})
		return
	}

	// Service layer error
	err := uh.userService.UpdatePassword(ctx, updatePass)
	if err != nil {
		uh.logger.Error(fmt.Sprintf("Error updating password: %v", err))
		ctx.JSON(500, models.ErrorResponse{
			Status:  500,
			Message: "Failed to update password",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(200, models.SuccessResponce{
		Status:  200,
		Message: "Password updated successfully",
	})
}