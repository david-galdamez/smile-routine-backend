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

func (s *UserSettingService) GetUserSetting(ctx context.Context, userId int) utils.ServiceResponse[int] {
	waitMinute, err := s.userSettingRepo.GetUserSetting(ctx, userId)
	if err != nil {
		return utils.Error[int]("Error obteniendo configuracion")
	}
	return utils.Ok[int](waitMinute)
}

func (s *UserSettingService) UpdateUserSetting(ctx context.Context, userID int, minutes int) utils.ServiceResponse[int] {

	waitMinute, err := s.userSettingRepo.UpdateUserSetting(ctx, userID, minutes)
	if err != nil {
		return utils.Error[int]("Error actualizando configuracion")
	}
	return utils.Ok[int](waitMinute)
}
