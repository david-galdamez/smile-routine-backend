package habits

import (
	"encoding/json"
	"net/http"

	"github.com/david-galdamez/smile-routine-backend/middleware"
	"github.com/david-galdamez/smile-routine-backend/utils"
)

type HabitsHandler struct {
	habitsService *HabitsService
}

func NewHabitsHandler(habitsService *HabitsService) http.Handler {
	mux := http.NewServeMux()
	hh := &HabitsHandler{
		habitsService: habitsService,
	}

	mux.Handle("GET /", middleware.JWTMiddleware(http.HandlerFunc(hh.GetHabits)))
	mux.Handle("POST /register", middleware.JWTMiddleware(http.HandlerFunc(hh.RegisterHabit)))
	return http.StripPrefix("/api/habits", mux)
}

func (hh *HabitsHandler) GetHabits(w http.ResponseWriter, r *http.Request) {
	year := r.URL.Query().Get("year")
	month := r.URL.Query().Get("month")
	if year == "" || month == "" {
		utils.RespondWithError(w, http.StatusBadRequest, "El año y el mes son requeridos")
		return
	}

	authUser, ok := utils.GetAuthUser(r.Context())
	if !ok {
		utils.RespondWithError(w, http.StatusUnauthorized, "No autorizado")
		return
	}

	habitsResult := hh.habitsService.GetHabits(r.Context(), authUser.UserId, year, month)
	if !habitsResult.Success || habitsResult.Data == nil {
		utils.RespondWithError(w, http.StatusInternalServerError, *habitsResult.ErrorMessage)
		return
	}

	utils.RespondWithJson(w, http.StatusOK, utils.ApiResponse[[]HabitDto]{
		Success: true,
		Data:    habitsResult.Data,
	})
}

func (hh *HabitsHandler) RegisterHabit(w http.ResponseWriter, r *http.Request) {
	registerRequest := RegisterHabitDto{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&registerRequest); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	authUser, ok := utils.GetAuthUser(r.Context())
	if !ok {
		utils.RespondWithError(w, http.StatusUnauthorized, "No autorizado")
		return
	}

	registerResult := hh.habitsService.RegisterHabit(r.Context(), authUser.UserId, &registerRequest)
	if !registerResult.Success {
		utils.RespondWithError(w, http.StatusInternalServerError, *registerResult.ErrorMessage)
		return
	}

	msg := "Habito registrado exitosamente"
	utils.RespondWithJson(w, http.StatusCreated, utils.ApiResponse[any]{
		Success: true,
		Message: &msg,
	})
}
