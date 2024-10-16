package service

import (
	"context"
	"fmt"
	"log/slog"
	"time"
	"user-service/internal/config"
	"user-service/internal/models"
	"user-service/internal/security"
	"user-service/storage"

	"github.com/google/uuid"
)

type UserService interface {
	RegisterUser(context.Context, models.UserRegisterRequst) (*models.UserRegisterResponce, error)
	CheckIfUserExist(ctx context.Context, checkUser models.CheckUser) (bool, error)
	LoginUser(context.Context, models.LoginRequest) (*models.LoginResponse, error)
	GetUserProfile(context.Context, string) (*models.UserProfile, error)
	UpdateUserProfile(context.Context, models.UpdateUserParams) error
	UpdateUserProfileAdmin(ctx context.Context, updateUser models.UpdateUserAdmin) error
	DeleteUser(context.Context, string) error
	UpdatePassword(context.Context, models.UpdatePasswordParams) error
	GetAllUsers(ctx context.Context, fEmployee models.GetAllUsersReq) (*storage.GetAllUsersRow, error)
}

type userServiceImpl struct {
	userRepository *storage.Queries
	logger         *slog.Logger
	cfg            config.Config
}

func NewUserService(queries *storage.Queries, logger *slog.Logger) UserService {
	return &userServiceImpl{
		userRepository: queries,
		logger:         logger,
		cfg:            config.Load(),
	}
}

func (us *userServiceImpl) RegisterUser(ctx context.Context, user models.UserRegisterRequst) (*models.UserRegisterResponce, error) {
	hashedPassword, err := security.HashPassword(user.Password)
	if err != nil {
		us.logger.ErrorContext(ctx, fmt.Sprintf("Error in hashed password: %s", err.Error()))
		return nil, err
	}
	uid, err := us.userRepository.CreateUser(ctx, storage.CreateUserParams{
		ID:           uuid.New(),
		Username:     user.Username,
		FullName:     user.FullName,
		PhoneNumber:  user.PhoneNumber,
		PasswordHash: hashedPassword,
		Role:         user.Role,
	})

	return &models.UserRegisterResponce{
		Id:      uid.String(),
		Message: "User registered successfully",
	}, nil
}

func (us *userServiceImpl) CheckIfUserExist(ctx context.Context, checkUser models.CheckUser) (bool, error) {
	count, err := us.userRepository.CheckIfUserExists(ctx, storage.CheckIfUserExistsParams{
		Username:    checkUser.Username,
		PhoneNumber: checkUser.PhoneNumber,
	})

	if err != nil {
		us.logger.Error("User that ckecked error")
		return false, err
	}
	return count > 0, nil
}

func (us *userServiceImpl) LoginUser(ctx context.Context, login models.LoginRequest) (*models.LoginResponse, error) {
	loginUser, err := us.userRepository.LoginUser(ctx, login.Username)
	if err != nil {
		us.logger.Error(fmt.Sprintf("Error in login user: %s", err.Error()))
		return nil, err
	}

	check := security.CheckPasswordHash(login.Password, loginUser.PasswordHash)
	if !check {
		us.logger.Error(fmt.Sprintf("Passwrod is incorrect"))
		return nil, fmt.Errorf("username or password is incorrect")
	}

	accessToken, err := security.GenerateJWTToken(security.TokenClaims{
		ID:       loginUser.ID.String(),
		Username: loginUser.Username,
		Role:     loginUser.Role,
	}, us.cfg.SECRET_KEY, time.Duration(time.Minute*20))
	if err != nil {
		us.logger.Error(fmt.Sprintf("Error in generate access token: %s", err.Error()))
	}

	refreshToken, err := security.GenerateJWTToken(security.TokenClaims{
		ID:       loginUser.ID.String(),
		Username: loginUser.Username,
		Role:     loginUser.Role,
	}, us.cfg.SECRET_KEY, time.Duration(7*24*time.Hour))
	if err != nil {
		us.logger.Error(fmt.Sprintf("Error in generate refresh token: %s", err.Error()))
	}

	return &models.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (us *userServiceImpl) GetUserProfile(ctx context.Context, id string) (*models.UserProfile, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		us.logger.Error(fmt.Sprintf("Error in parse uuid: %s", err.Error()))
		return nil, err
	}
	userProfile, err := us.userRepository.GetUser(ctx, uid)
	if err != nil {
		us.logger.Error(fmt.Sprintf("Error in get user profile: %s", err.Error()))
		return nil, err
	}

	return &models.UserProfile{
		ID:          userProfile.ID.String(),
		Username:    userProfile.Username,
		FullName:    userProfile.FullName,
		PhoneNumber: userProfile.PhoneNumber,
		Role:        userProfile.Role,
	}, nil
}

func (us *userServiceImpl) UpdateUserProfile(ctx context.Context, updateUser models.UpdateUserParams) error {
	uid, err := uuid.Parse(updateUser.ID)
	if err != nil {
		us.logger.Error(fmt.Sprintf("Error in parse uuid: %s", err.Error()))
		return err
	}

	hashedPassword, err := us.userRepository.GetUserPassword(ctx, uid)
	if err != nil {
		us.logger.Error(fmt.Sprintf("Error in check password %s", err.Error()))
		return err
	}

	check := security.CheckPasswordHash(updateUser.Password, hashedPassword)
	if !check {
		us.logger.Error(fmt.Sprintf("Passwrod is incorrect"))
		return fmt.Errorf("username or password is incorrect")
	}

	err = us.userRepository.UpdateUser(ctx, storage.UpdateUserParams{
		ID:          uid,
		Username:    updateUser.Username,
		FullName:    updateUser.FullName,
		PhoneNumber: updateUser.PhoneNumber,
	})

	if err != nil {
		us.logger.Error(fmt.Sprintf("Error in update user: %s", err.Error()))
		return err
	}

	return nil
}

func (us *userServiceImpl) UpdateUserProfileAdmin(ctx context.Context, updateUser models.UpdateUserAdmin) error {
	uid, err := uuid.Parse(updateUser.ID)
	if err != nil {
		us.logger.Error(fmt.Sprintf("Error in parse uuid: %s", err.Error()))
		return err
	}

	hashedPassword, err := security.HashPassword(updateUser.Password)
	if err != nil {
		us.logger.Error(fmt.Sprintf("Error in genereted hash password: %s", err.Error()))
		return err
	}

	err = us.userRepository.UpdateAdminUser(ctx, storage.UpdateUserAdminParams{
		ID:          uid,
		Username:    updateUser.Username,
		FullName:    updateUser.FullName,
		PhoneNumber: updateUser.PhoneNumber,
		Role:        updateUser.Role,
		Password:    hashedPassword,
	})

	if err != nil {
		us.logger.Error(fmt.Sprintf("Error in update user: %s", err.Error()))
		return err
	}

	return nil
}

func (us *userServiceImpl) DeleteUser(ctx context.Context, id string) error {
	uid, err := uuid.Parse(id)
	if err != nil {
		us.logger.Error(fmt.Sprintf("Error in parse to uuid: %s", err.Error()))
		return err
	}

	err = us.userRepository.DeleteUser(ctx, uid)
	if err != nil {
		us.logger.Error(fmt.Sprintf("Error in deleted user: %s", err.Error()))
		return err
	}

	return nil
}

func (us *userServiceImpl) UpdatePassword(ctx context.Context, updatePass models.UpdatePasswordParams) error {
	uid, err := uuid.Parse(updatePass.ID)
	if err != nil {
		us.logger.Error(fmt.Sprintf("Error in parse uuid: %s", err.Error()))
		return err
	}

	hashedPassword, err := us.userRepository.GetUserPassword(ctx, uid)
	if err != nil {
		us.logger.Error(fmt.Sprintf("Error in check password %s", err.Error()))
		return err
	}

	check := security.CheckPasswordHash(updatePass.OldPassword, hashedPassword)
	if !check {
		us.logger.Error(fmt.Sprintf("Passwrod is incorrect"))
		return fmt.Errorf("password is incorrect")
	}
	hashedPassword, err = security.HashPassword(updatePass.NewPassword)
	if err != nil {
		us.logger.ErrorContext(ctx, fmt.Sprintf("Error in hashed password: %s", err.Error()))
		return err
	}
	err = us.userRepository.UpdatePassword(ctx, storage.UpdatePasswordParams{
		ID:           uid,
		PasswordHash: hashedPassword,
	})

	if err != nil {
		us.logger.Error("Error in updated password")
		return err
	}

	return nil
}

func (us *userServiceImpl) GetAllUsers(ctx context.Context, fEmployee models.GetAllUsersReq) (*storage.GetAllUsersRow, error) {
	resp, err := us.userRepository.GetAllUsers(ctx, storage.GetAllUsersParams{
		Username:    fEmployee.Username,
		FullName:    fEmployee.FullName,
		PhoneNumber: fEmployee.PhoneNumber,
		Role:        fEmployee.Role,
		Limit:       fEmployee.Limit,
		Offset:      (fEmployee.Page - 1) * fEmployee.Limit,
	})
	if err != nil {
		us.logger.Error(fmt.Sprintf("Eror in get all users: %s", err.Error()))
		return nil, err
	}

	return resp, nil
}
