package user_settings

import (
	"encoding/json"
	"net/http"

	"github.com/david-galdamez/smile-routine-backend/middleware"
	"github.com/david-galdamez/smile-routine-backend/utils"
)

type UserSettingsHandler struct {
	userSettingService *UserSettingService
}

func NewUserSettingsHandler(userSettingService *UserSettingService) http.Handler {
	mux := http.NewServeMux()
	ush := &UserSettingsHandler{
		userSettingService: userSettingService,
	}
	mux.Handle("PUT /", middleware.JWTMiddleware(http.HandlerFunc(ush.UpdateUserSettings)))

	return http.StripPrefix("/api/user-settings", mux)
}

func (ush *UserSettingsHandler) UpdateUserSettings(w http.ResponseWriter, r *http.Request) {
	updateRequest := UserSettingsUpdateDto{}

	authUser, ok := utils.GetAuthUser(r.Context())
	if !ok {
		utils.RespondWithError(w, http.StatusUnauthorized, "No estas autorizado")
		return
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&updateRequest); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	updateResponse := ush.userSettingService.UpdateUserSetting(r.Context(), authUser.UserId, updateRequest.Minutes)
	if !updateResponse.Success {
		utils.RespondWithError(w, http.StatusInternalServerError, *updateResponse.ErrorMessage)
		return
	}

	msg := "Configuracion actualizada correctamente"
	utils.RespondWithJson(w, http.StatusOK, utils.ApiResponse[UserSettingsUpdateDto]{
		Success: true,
		Message: &msg,
		Data:    &updateRequest,
	})
}
