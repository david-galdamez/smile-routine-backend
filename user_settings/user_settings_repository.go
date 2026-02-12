package user_settings

import (
	"context"

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

func (usr *UserSettingsRepository) UpdateUserSetting(ctx context.Context, userId int, minutes int) error {
	return usr.db.UpdateUserSetting(ctx, database.UpdateUserSettingParams{
		UserID:      int32(userId),
		WaitMinutes: int32(minutes),
	})
}
