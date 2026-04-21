package user_settings

import (
	"context"
	"log"

	"github.com/david-galdamez/smile-routine-backend/database"
)

type UserSettingsRepository struct {
	db *database.Queries
}

func NewUserSettingsRepository(db *database.Queries) *UserSettingsRepository {
	return &UserSettingsRepository{db: db}
}

func (usr *UserSettingsRepository) RegisterUserSetting(ctx context.Context, userId int) error {
	return usr.db.RegisterUserSetting(ctx, database.RegisterUserSettingParams{
		UserID:      int32(userId),
		WaitMinutes: 30,
	})
}

func (usr *UserSettingsRepository) GetUserSetting(ctx context.Context, userId int) (int, error) {
	waitMinute, err := usr.db.GetUserSetting(ctx, int32(userId))
	if err != nil {
		log.Printf("err: %v", err.Error())
		return 0, err
	}
	return int(waitMinute), nil
}

func (usr *UserSettingsRepository) UpdateUserSetting(ctx context.Context, userId int, minutes int) (int, error) {
	waitMinute, err := usr.db.UpdateUserSetting(ctx, database.UpdateUserSettingParams{
		UserID:      int32(userId),
		WaitMinutes: int32(minutes),
	})
	if err != nil {
		log.Printf("err: %v", err.Error())
		return 0, err
	}
	return int(waitMinute), nil
}
