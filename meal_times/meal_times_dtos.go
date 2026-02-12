package mealtimes

type MealTimeDto struct {
	Breakfast string `json:"breakfast" validate:"required"`
	Lunch     string `json:"lunch" validate:"required"`
	Dinner    string `json:"dinner" validate:"required"`
}
