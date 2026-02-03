package users

import (
	"context"
	"log"
	"time"

	"github.com/david-galdamez/smile-routine-backend/database"
)

type UserRepository struct {
	Db *database.Queries
}

func NewUserRepository(db *database.Queries) *UserRepository {
	return &UserRepository{Db: db}
}

func (ur *UserRepository) IsEmailRegistered(ctx context.Context, email string) bool {
	exists, err := ur.Db.IsEmailRegistered(ctx, email)
	if err != nil {
		log.Printf("error checking email registration, %v", err)
		return false
	}

	return exists
}

func (ur *UserRepository) CreateUser(ctx context.Context, user *RegisterUserDto, passwordHash string, birthDate time.Time) (int, error) {
	id, err := ur.Db.CreateUser(ctx, database.CreateUserParams{
		Name:         user.Name,
		Email:        user.Email,
		PasswordHash: passwordHash,
		BirthDate:    birthDate,
		Gender:       string(user.Gender),
	})
	if err != nil {
		log.Printf("error creating user, %v", err)
		return 0, err
	}

	return int(id), nil
}

func (ur *UserRepository) GetUserByEmail(ctx context.Context, email string) (*database.User, error) {
	user, err := ur.Db.GetUserByEmail(ctx, email)
	if err != nil {
		log.Printf("error getting user, %v", err)
		return nil, err
	}

	return &user, nil
}
