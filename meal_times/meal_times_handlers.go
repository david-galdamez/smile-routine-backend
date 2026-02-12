package mealtimes

import (
	"encoding/json"
	"net/http"

	"github.com/david-galdamez/smile-routine-backend/middleware"
	"github.com/david-galdamez/smile-routine-backend/utils"
)

type MealTimeHandler struct {
	mts *MealTimeService
}

func NewMealTimeHandler(mts *MealTimeService) http.Handler {
	mux := http.NewServeMux()
	mth := &MealTimeHandler{
		mts: mts,
	}

	mux.Handle("GET /", middleware.JWTMiddleware(http.HandlerFunc(mth.GetMealTimes)))
	mux.Handle("PUT /update", middleware.JWTMiddleware(http.HandlerFunc(mth.UpdateMealTimes)))

	return http.StripPrefix("/api/meal-time", mux)
}

func (mth *MealTimeHandler) GetMealTimes(w http.ResponseWriter, r *http.Request) {
	authUser, ok := utils.GetAuthUser(r.Context())
	if !ok {
		utils.RespondWithError(w, http.StatusUnauthorized, "No Autorizado")
		return
	}

	mealTimeResult := mth.mts.GetMealTimes(r.Context(), authUser.UserId)
	if !mealTimeResult.Success {
		utils.RespondWithError(w, http.StatusInternalServerError, *mealTimeResult.ErrorMessage)
		return
	}

	utils.RespondWithJson(w, http.StatusOK, utils.ApiResponse[MealTimeDto]{
		Success: true,
		Data:    mealTimeResult.Data,
	})
}

func (mth *MealTimeHandler) UpdateMealTimes(w http.ResponseWriter, r *http.Request) {
	updateRequst := MealTimeDto{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&updateRequst); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	authUser, ok := utils.GetAuthUser(r.Context())
	if !ok {
		utils.RespondWithError(w, http.StatusUnauthorized, "No Autorizado")
		return
	}

	updateResult := mth.mts.UpdateMealTime(r.Context(), authUser.UserId, &updateRequst)
	if !updateResult.Success {
		utils.RespondWithError(w, http.StatusInternalServerError, *updateResult.ErrorMessage)
		return
	}

	msg := "Hora de comidas actualizada con éxito"
	utils.RespondWithJson(w, http.StatusOK, utils.ApiResponse[MealTimeDto]{
		Success: true,
		Message: &msg,
		Data:    updateResult.Data,
	})
}
