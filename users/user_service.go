package users

import (
	"context"
	"log"
	"strings"
	"time"

	"github.com/david-galdamez/smile-routine-backend/utils"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepository *UserRepository
}

func NewUserService(userRepository *UserRepository) *UserService {
	return &UserService{userRepository: userRepository}
}

func (us *UserService) IsEmailRegistered(ctx context.Context, email string) bool {
	return us.userRepository.IsEmailRegistered(ctx, email)
}

func (us *UserService) RegisterUser(ctx context.Context, user *RegisterUserDto) utils.ServiceResponse[int] {
	if !user.Gender.IsValid() {
		return utils.Error[int]("Genero invalido")
	}

	birthDate, err := time.Parse("2006-01-02", user.BirthDate)
	if err != nil {
		return utils.Error[int]("Fecha de nacimiento invalida")
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)
	if err != nil {
		return utils.Error[int]("Error al hashear contraseña")
	}

	id, err := us.userRepository.CreateUser(ctx, user, string(passwordHash), birthDate)
	if err != nil {
		return utils.Error[int]("Error al crear usuario")
	}

	return utils.Ok[int](id)
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
