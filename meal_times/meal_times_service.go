package mealtimes

import (
	"context"
	"strings"
	"time"

	"github.com/david-galdamez/smile-routine-backend/user_settings"
	"github.com/david-galdamez/smile-routine-backend/utils"
)

type MealTimeService struct {
	mtr *MealTimeRepository
	usr *user_settings.UserSettingsRepository
}

func NewMealTimeService(repository *MealTimeRepository, usr *user_settings.UserSettingsRepository) *MealTimeService {
	return &MealTimeService{
		mtr: repository,
		usr: usr,
	}
}

func (mts *MealTimeService) GetMealTimes(ctx context.Context, userId int) utils.ServiceResponse[MealTimeDto] {
	mealTimes, err := mts.mtr.GetMealTimes(ctx, userId)
	if err != nil {
		if strings.Contains(err.Error(), "no rows") {
			return utils.Ok(MealTimeDto{})
		}
		return utils.Error[MealTimeDto]("Error al obtener las horas de comida")
	}

	userSetting, err := mts.usr.GetUserSetting(ctx, userId)
	if err != nil {
		return utils.Error[MealTimeDto]("Error al obtener el tiempo de espera.")
	}

	mealTimeDto := MealTimeDto{
		Breakfast: mealTimes.Breakfast.Format("15:04"),
		Lunch:     mealTimes.Lunch.Format("15:04"),
		Dinner:    mealTimes.Dinner.Format("15:04"),
		Offset:    userSetting,
	}

	return utils.Ok(mealTimeDto)
}

func (mts *MealTimeService) UpdateMealTime(ctx context.Context, userId int, updateRequest *MealTimeDto) utils.ServiceResponse[MealTimeDto] {

	breakfastTime, err := time.Parse("15:04", updateRequest.Breakfast)
	if err != nil {
		return utils.Error[MealTimeDto]("Formato de hora incorrecto")
	}

	lunchTime, err := time.Parse("15:04", updateRequest.Lunch)
	if err != nil {
		return utils.Error[MealTimeDto]("Formato de hora incorrecto")
	}

	dinnerTime, err := time.Parse("15:04", updateRequest.Dinner)
	if err != nil {
		return utils.Error[MealTimeDto]("Formato de hora incorrecto")
	}

	userSetting, err := mts.usr.UpdateUserSetting(ctx, userId, updateRequest.Offset)
	if err != nil {
		return utils.Error[MealTimeDto]("Error al actualizar los minutos.")
	}

	mealTime, err := mts.mtr.UpdateMealTime(ctx, userId, breakfastTime, lunchTime, dinnerTime)
	if err != nil {
		return utils.Error[MealTimeDto]("Error al actualizar la hora de comida")
	}

	mealTimeDto := MealTimeDto{
		Breakfast: mealTime.Breakfast.Format("15:04"),
		Lunch:     mealTime.Lunch.Format("15:04"),
		Dinner:    mealTime.Dinner.Format("15:04"),
		Offset:    userSetting,
	}

	return utils.Ok(mealTimeDto)
}
