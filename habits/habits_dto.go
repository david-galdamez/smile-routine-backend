package habits

type RegisterHabitDto struct {
	HabitDate string `json:"habit_date" validate:"required,datetime=2006-01-02"`
	Completed bool   `json:"completed" validate:"required"`
	MealID    int    `json:"meal_id" validate:"required"`
}

type HabitDto struct {
	HabitDate  string  `json:"habit_date" validate:"datetime=2006-01-02"`
	Porcentage float64 `json:"porcentage"`
}

type HabitOfDayDto struct {
	Id       *int        `json:"id"`
	MealName string      `json:"meal_name"`
	Status   HabitStatus `json:"status"`
}

type HabitStatus string

const (
	StatusPending   HabitStatus = "pending"
	StatusCompleted HabitStatus = "completed"
	StatusFailed    HabitStatus = "failed"
)
