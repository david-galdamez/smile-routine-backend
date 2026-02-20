package appointment

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/david-galdamez/smile-routine-backend/middleware"
	"github.com/david-galdamez/smile-routine-backend/utils"
)

type AppointmentHandler struct {
	service *AppointmentService
}

func NewAppointmentHandler(service *AppointmentService) http.Handler {
	mux := http.NewServeMux()
	appo := &AppointmentHandler{service: service}
	mux.Handle("POST /register", middleware.JWTMiddleware(http.HandlerFunc(appo.RegisterAppointment)))
	mux.Handle("PUT /update/{id}", middleware.JWTMiddleware(http.HandlerFunc(appo.UpdateAppointment)))
	return http.StripPrefix("/api/appointment", mux)
}

func (ah *AppointmentHandler) RegisterAppointment(w http.ResponseWriter, r *http.Request) {

	registerDto := RegisterAppointmentDto{}

	authUser, ok := utils.GetAuthUser(r.Context())
	if !ok {
		utils.RespondWithError(w, http.StatusUnauthorized, "No autorizado")
		return
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&registerDto); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	registerResult := ah.service.CreateAppointment(r.Context(), authUser.UserId, &registerDto)
	if !registerResult.Success || registerResult.Data == nil {
		utils.RespondWithError(w, http.StatusBadRequest, *registerResult.ErrorMessage)
		return
	}

	msg := "Cita registrada exitosamente"
	utils.RespondWithJson(w, http.StatusCreated, utils.ApiResponse[AppointmentDto]{
		Success: true,
		Data:    registerResult.Data,
		Message: &msg,
	})
}

func (ah *AppointmentHandler) UpdateAppointment(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		utils.RespondWithError(w, http.StatusBadRequest, "ID de cita no proporcionado")
		return
	}

	parsedId, err := strconv.Atoi(id)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "ID de cita no válido")
		return
	}

	updateResult := ah.service.UpdateAppointment(r.Context(), true, parsedId)
	if !updateResult.Success || updateResult.Data == nil {
		utils.RespondWithError(w, http.StatusBadRequest, *updateResult.ErrorMessage)
		return
	}

	msg := "Cita actualizada exitosamente"
	utils.RespondWithJson(w, http.StatusOK, utils.ApiResponse[AppointmentDto]{
		Success: true,
		Data:    updateResult.Data,
		Message: &msg,
	})
}
