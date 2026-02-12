package users

import (
	"context"
	"log"
	"strings"
	"time"

	mealtimes "github.com/david-galdamez/smile-routine-backend/meal_times"
	"github.com/david-galdamez/smile-routine-backend/user_settings"
	"github.com/david-galdamez/smile-routine-backend/utils"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepository         *UserRepository
	mealTimeRepository     *mealtimes.MealTimeRepository
	userSettingsRepository *user_settings.UserSettingsRepository
}

func NewUserService(userRepository *UserRepository, mealTimeRepository *mealtimes.MealTimeRepository, userSettingsRepository *user_settings.UserSettingsRepository) *UserService {
	return &UserService{
		userRepository:         userRepository,
		mealTimeRepository:     mealTimeRepository,
		userSettingsRepository: userSettingsRepository,
	}
}

func (us *UserService) IsEmailRegistered(ctx context.Context, email string) bool {
	return us.userRepository.IsEmailRegistered(ctx, email)
}

func (us *UserService) DoesUserExist(ctx context.Context, userId int) bool {
	return us.userRepository.DoesUserExist(ctx, userId)
}

func (us *UserService) RegisterUser(ctx context.Context, user *RegisterUserDto) utils.ServiceResponse[LoginResponseDto] {
	if !user.Gender.IsValid() {
		return utils.Error[LoginResponseDto]("Genero invalido")
	}

	birthDate, err := time.Parse("2006-01-02", user.BirthDate)
	if err != nil {
		return utils.Error[LoginResponseDto]("Fecha de nacimiento invalida")
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)
	if err != nil {
		return utils.Error[LoginResponseDto]("Error al hashear contraseña")
	}

	id, err := us.userRepository.CreateUser(ctx, user, string(passwordHash), birthDate)
	if err != nil {
		return utils.Error[LoginResponseDto]("Error al crear usuario")
	}

	timeString := "00:00"
	value, err := time.Parse("15:34", timeString)
	if err != nil {
		return utils.Error[LoginResponseDto]("Error al parsear hora")
	}

	_, err = us.mealTimeRepository.RegisterMealTime(ctx, id, value, value, value)
	if err != nil {
		return utils.Error[LoginResponseDto]("Error al registrar horarios")
	}

	err = us.userSettingsRepository.RegisterUserSetting(ctx, id)
	if err != nil {
		return utils.Error[LoginResponseDto]("Error al registrar configuracion de usuario")
	}

	token, err := utils.GenerateJWT(id)
	if err != nil {
		return utils.Error[LoginResponseDto]("Error al generar token")
	}

	return utils.Ok[LoginResponseDto](LoginResponseDto{
		Token: token,
	})
}

func (us *UserService) LoginUser(ctx context.Context, login *LoginUserDto) utils.ServiceResponse[LoginResponseDto] {
	user, err := us.userRepository.GetUserByEmail(ctx, login.Email)
	if err != nil {
		if strings.Contains(err.Error(), "no rows") {
			return utils.Error[LoginResponseDto]("Email no registrado")
		}
		return utils.Error[LoginResponseDto]("Error al obtener usuario")
	}

	if user == nil {
		return utils.Error[LoginResponseDto]("Email no registrado")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(login.Password)); err != nil {
		return utils.Error[LoginResponseDto]("Contraseña incorrecta")
	}

	token, err := utils.GenerateJWT(int(user.ID))
	if err != nil {
		log.Printf("error: %v", err)
		return utils.Error[LoginResponseDto]("Error al generar token")
	}

	return utils.Ok[LoginResponseDto](LoginResponseDto{
		Token: token,
	})
}

func (us *UserService) GetUserById(ctx context.Context, id int) utils.ServiceResponse[UserDto] {
	user, err := us.userRepository.GetUserById(ctx, id)
	if err != nil {
		if strings.Contains(err.Error(), "no rows") {
			return utils.Error[UserDto]("Usuario no encontrado")
		}
		return utils.Error[UserDto]("Error al obtener usuario")
	}

	if user == nil {
		return utils.Error[UserDto]("Usuario no encontrado")
	}

	return utils.Ok[UserDto](UserDto{
		Id:        int(user.ID),
		Name:      user.Name,
		Email:     user.Email,
		BirthDate: user.BirthDate.Format("2006-01-02"),
		Gender:    ApiGender(user.Gender),
	})
}

func (us *UserService) UpdateUser(ctx context.Context, id int, updateUser *UpdateUserDto) utils.ServiceResponse[UserDto] {

	birthDate, err := time.Parse("2006-01-02", updateUser.BirthDate)
	if err != nil {
		return utils.Error[UserDto]("Fecha de nacimiento invalida")
	}

	updateResult, err := us.userRepository.UpdateUser(ctx, id, updateUser, birthDate)
	if err != nil {
		return utils.Error[UserDto]("Error al actualizar usuario")
	}

	return utils.Ok[UserDto](UserDto{
		Id:        int(updateResult.ID),
		Name:      updateResult.Name,
		Email:     updateResult.Email,
		BirthDate: updateResult.BirthDate.Format("2006-01-02"),
		Gender:    ApiGender(updateResult.Gender),
	})
}

func (us *UserService) UpdatePassword(ctx context.Context, userId int, updatePassword *UpdatePasswordDto) utils.ServiceResponse[any] {
	user, err := us.userRepository.GetUserById(ctx, userId)
	if err != nil {
		return utils.Error[any]("Error al obtener usuario")
	}

	if user == nil {
		return utils.Error[any]("Usuario no encontrado")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(updatePassword.OldPassword)); err != nil {
		return utils.Error[any]("Contraseña actual incorrecta")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(updatePassword.NewPassword), bcrypt.MinCost)
	if err != nil {
		return utils.Error[any]("Error al generar contraseña")
	}

	err = us.userRepository.UpdatePassword(ctx, userId, string(hashedPassword))
	if err != nil {
		return utils.Error[any]("Error al actualizar contraseña")
	}

	return utils.Ok[any](nil)
}
