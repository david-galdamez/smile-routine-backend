package habits

import (
	"context"
	"strconv"
	"time"

	"github.com/david-galdamez/smile-routine-backend/utils"
)

type HabitsService struct {
	habitsRepository *HabitsRepository
}

func NewHabitsService(habitsRepository *HabitsRepository) *HabitsService {
	return &HabitsService{
		habitsRepository: habitsRepository,
	}
}

func (s *HabitsService) GetHabits(ctx context.Context, userId int, year string, month string) utils.ServiceResponse[[]HabitDto] {

	yearInt, err := strconv.Atoi(year)
	if err != nil {
		return utils.Error[[]HabitDto]("Formato de año incorrecto")
	}

	monthInt, err := strconv.Atoi(month)
	if err != nil {
		return utils.Error[[]HabitDto]("Formato de mes incorrecto")
	}

	habits, err := s.habitsRepository.GetHabits(ctx, userId, yearInt, monthInt)
	if err != nil {
		return utils.Error[[]HabitDto]("Error al obtener los hábitos")
	}

	habitsDto := make([]HabitDto, 0, len(habits))
	for _, habit := range habits {

		dateString := habit.Day.Format("2006-01-02")
		porcentage := (float64(habit.CompletedCount) / 3.0) * 100

		if porcentage > 100 {
			porcentage = 100
		}

		newHabit := HabitDto{
			HabitDate:  dateString,
			Porcentage: porcentage,
		}

		habitsDto = append(habitsDto, newHabit)
	}

	return utils.Ok(habitsDto)
}

func (s *HabitsService) GetHabitsOfDay(ctx context.Context, userId int, habitDate string) utils.ServiceResponse[[]HabitOfDayDto] {
	if habitDate == "" {
		return utils.Error[[]HabitOfDayDto]("La fecha es requerida")
	}

	habitDateValue, err := time.Parse("2006-01-02", habitDate)
	if err != nil {
		return utils.Error[[]HabitOfDayDto]("Formato de fecha incorrecto")
	}

	habits, err := s.habitsRepository.GetHabitsOfDay(ctx, userId, habitDateValue)
	if err != nil {
		return utils.Error[[]HabitOfDayDto]("Error al obtener los hábitos")
	}

	habitsDto := make([]HabitOfDayDto, 0, len(habits))
	for _, habit := range habits {
		newHabit := HabitOfDayDto{}
		newHabit.Id = utils.NullInt64ToPtr(habit.ID)
		newHabit.MealName = habit.MealName
		newHabit.Status = HabitStatus(habit.Status)
		habitsDto = append(habitsDto, newHabit)
	}

	return utils.Ok(habitsDto)
}

func (s *HabitsService) RegisterHabit(ctx context.Context, userId int, registerRequest *RegisterHabitDto) utils.ServiceResponse[any] {

	habitDate, err := time.Parse("2006-01-02", registerRequest.HabitDate)
	if err != nil {
		return utils.Error[any]("Formato de fecha incorrecto")
	}

	if err := s.habitsRepository.RegisterHabit(ctx, userId, habitDate, registerRequest.Completed, registerRequest.MealID); err != nil {
		return utils.Error[any]("Error al registrar el habito")
	}

	return utils.Ok[any](nil)
}
