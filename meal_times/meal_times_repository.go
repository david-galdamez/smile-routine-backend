package mealtimes

import (
	"context"
	"time"

	"github.com/david-galdamez/smile-routine-backend/database"
)

type MealTimeRepository struct {
	db *database.Queries
}

func NewMealTimeRepository(db *database.Queries) *MealTimeRepository {
	return &MealTimeRepository{
		db: db,
	}
}

func (mlr *MealTimeRepository) GetMealTimes(ctx context.Context, userId int) (*database.MealTime, error) {
	mealTimes, err := mlr.db.GetMealTimes(ctx, int32(userId))
	if err != nil {
		return nil, err
	}
	return &mealTimes, nil
}

func (mlr *MealTimeRepository) RegisterMealTime(ctx context.Context, userID int, breakfast, lunch, dinner time.Time) (*database.MealTime, error) {
	mealTime, err := mlr.db.RegisterMealTime(ctx, database.RegisterMealTimeParams{
		UserID:    int32(userID),
		Breakfast: breakfast,
		Lunch:     lunch,
		Dinner:    dinner,
	})
	if err != nil {
		return nil, err
	}
	return &mealTime, nil
}

func (mlr *MealTimeRepository) UpdateMealTime(ctx context.Context, userId int, breakfast, lunch, dinner time.Time) (*database.MealTime, error) {
	mealTime, err := mlr.db.UpdateMealTime(ctx, database.UpdateMealTimeParams{
		UserID:    int32(userId),
		Breakfast: breakfast,
		Lunch:     lunch,
		Dinner:    dinner,
	})
	if err != nil {
		return nil, err
	}
	return &mealTime, nil
}
