package user_settings

import (
	"context"

	"github.com/david-galdamez/smile-routine-backend/utils"
)

type UserSettingService struct {
	userSettingRepo *UserSettingsRepository
}

func NewUserSettingService(userSettingRepo *UserSettingsRepository) *UserSettingService {
	return &UserSettingService{
		userSettingRepo: userSettingRepo,
	}
}

func (s *UserSettingService) UpdateUserSetting(ctx context.Context, userID int, minutes int) utils.ServiceResponse[any] {

	err := s.userSettingRepo.UpdateUserSetting(ctx, userID, minutes)
	if err != nil {
		return utils.Error[any]("Error actualizando configuracion")
	}
	return utils.Ok[any](nil)
}
