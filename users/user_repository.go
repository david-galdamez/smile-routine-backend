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

func (ur *UserRepository) DoesUserExist(ctx context.Context, userId int) bool {
	exists, err := ur.Db.DoesUserExist(ctx, int32(userId))
	if err != nil {
		log.Printf("error checking if user exists, %v", err)
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

func (ur *UserRepository) GetUserById(ctx context.Context, id int) (*database.User, error) {
	user, err := ur.Db.GetUserById(ctx, int32(id))
	if err != nil {
		log.Printf("error getting user, %v", err)
		return nil, err
	}

	return &user, nil
}

func (ur *UserRepository) UpdateUser(ctx context.Context, id int, updateUser *UpdateUserDto, birthDate time.Time) (*database.User, error) {
	user, err := ur.Db.UpdateUser(ctx, database.UpdateUserParams{
		ID:        int32(id),
		Name:      updateUser.Name,
		Email:     updateUser.Email,
		BirthDate: birthDate,
		Gender:    string(updateUser.Gender),
	})
	if err != nil {
		log.Printf("error updating user, %v", err)
		return nil, err
	}

	return &user, nil
}

func (ur *UserRepository) UpdatePassword(ctx context.Context, userId int, passwordHash string) error {
	err := ur.Db.UpdatePassword(ctx, database.UpdatePasswordParams{
		ID:           int32(userId),
		PasswordHash: passwordHash,
	})
	if err != nil {
		log.Printf("error updating password, %v", err)
		return err
	}

	return nil
}
