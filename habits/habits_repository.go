package habits

import (
	"context"
	"time"

	"github.com/david-galdamez/smile-routine-backend/database"
)

type HabitsRepository struct {
	db *database.Queries
}

func NewHabitsRepository(db *database.Queries) *HabitsRepository {
	return &HabitsRepository{db: db}
}

func (r *HabitsRepository) GetHabits(ctx context.Context, userID int, year int, month int) ([]database.GetHabitsRow, error) {

	start := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)
	end := start.AddDate(0, 1, 0)

	return r.db.GetHabits(ctx, database.GetHabitsParams{
		UserID:  int32(userID),
		Column2: start,
		Column3: end,
	})
}

func (r *HabitsRepository) GetHabitsOfDay(ctx context.Context, userID int, habitDate time.Time) ([]database.GetHabitsOfDayRow, error) {
	return r.db.GetHabitsOfDay(ctx, database.GetHabitsOfDayParams{
		UserID: int32(userID),
		Column2: time.Date(
			habitDate.Year(),
			habitDate.Month(),
			habitDate.Day(),
			0, 0, 0, 0,
			time.UTC,
		),
	})
}

func (r *HabitsRepository) RegisterHabit(ctx context.Context, userID int, habitDate time.Time, completed bool, mealID int) error {
	return r.db.RegisterHabit(ctx, database.RegisterHabitParams{
		UserID:    int32(userID),
		HabitDate: habitDate,
		Completed: completed,
		MealID:    int32(mealID),
	})
}
