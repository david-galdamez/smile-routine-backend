package users

import (
	"encoding/json"
	"net/http"

	"github.com/david-galdamez/smile-routine-backend/middleware"
	"github.com/david-galdamez/smile-routine-backend/utils"
)

func NewUserHandler(us *UserService) http.Handler {
	mux := http.NewServeMux()
	uh := &UserHandler{userService: us}

	mux.HandleFunc("POST /register", uh.RegisterUser)
	mux.HandleFunc("POST /login", uh.LoginUser)
	mux.Handle("POST /logout", middleware.JWTMiddleware(http.HandlerFunc(uh.LogoutUser)))
	mux.Handle("GET /me", middleware.JWTMiddleware(http.HandlerFunc(uh.GetUser)))
	mux.Handle("PUT /update", middleware.JWTMiddleware(http.HandlerFunc(uh.UpdateUser)))
	mux.Handle("PATCH /update-password", middleware.JWTMiddleware(http.HandlerFunc(uh.UpdatePassword)))

	return http.StripPrefix("/api/users", mux)
}

type UserHandler struct {
	userService *UserService
}

func (uh *UserHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	registerUser := RegisterUserDto{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&registerUser); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	if valid := registerUser.Gender.IsValid(); !valid {
		utils.RespondWithError(w, http.StatusBadRequest, "Genero invalido")
		return
	}

	isEmailRegistered := uh.userService.IsEmailRegistered(r.Context(), registerUser.Email)
	if isEmailRegistered {
		utils.RespondWithError(w, http.StatusConflict, "Email ya registrado")
		return
	}

	registerResult := uh.userService.RegisterUser(r.Context(), &registerUser)
	if !registerResult.Success || registerResult.Data == nil {
		utils.RespondWithError(w, http.StatusBadRequest, *registerResult.ErrorMessage)
		return
	}

	message := "Usuario registrado con exito."
	utils.RespondWithJson(w, http.StatusCreated, utils.ApiResponse[LoginResponseDto]{
		Success: true,
		Message: &message,
		Data:    registerResult.Data,
	})
}

func (uh *UserHandler) LoginUser(w http.ResponseWriter, r *http.Request) {
	loginUser := LoginUserDto{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&loginUser); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	isEmailRegistered := uh.userService.IsEmailRegistered(r.Context(), loginUser.Email)
	if !isEmailRegistered {
		utils.RespondWithError(w, http.StatusNotFound, "Email no registrado")
		return
	}

	loginResult := uh.userService.LoginUser(r.Context(), &loginUser)
	if !loginResult.Success || loginResult.Data == nil {
		utils.RespondWithError(w, http.StatusBadRequest, *loginResult.ErrorMessage)
		return
	}

	message := "Usuario logueado con exito."
	utils.RespondWithJson(w, http.StatusOK, utils.ApiResponse[LoginResponseDto]{
		Success: true,
		Message: &message,
		Data:    loginResult.Data,
	})
}

func (uh *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	authUser, ok := utils.GetAuthUser(r.Context())
	if !ok {
		utils.RespondWithError(w, http.StatusUnauthorized, "Usuario no autorizado")
		return
	}

	userExists := uh.userService.DoesUserExist(r.Context(), authUser.UserId)
	if !userExists {
		utils.RespondWithError(w, http.StatusNotFound, "Usuario no encontrado")
		return
	}

	user := uh.userService.GetUserById(r.Context(), authUser.UserId)
	if !user.Success || user.Data == nil {
		utils.RespondWithError(w, http.StatusNotFound, *user.ErrorMessage)
		return
	}

	utils.RespondWithJson(w, http.StatusOK, utils.ApiResponse[UserDto]{
		Success: true,
		Data:    user.Data,
	})
}

func (uh *UserHandler) LogoutUser(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
}

func (uh *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	updateUser := UpdateUserDto{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&updateUser); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	authUser, ok := utils.GetAuthUser(r.Context())
	if !ok {
		utils.RespondWithError(w, http.StatusUnauthorized, "Usuario no autorizado")
		return
	}

	userExists := uh.userService.DoesUserExist(r.Context(), authUser.UserId)
	if !userExists {
		utils.RespondWithError(w, http.StatusNotFound, "Usuario no encontrado")
		return
	}

	updateResult := uh.userService.UpdateUser(r.Context(), authUser.UserId, &updateUser)
	if !updateResult.Success || updateResult.Data == nil {
		utils.RespondWithError(w, http.StatusBadRequest, *updateResult.ErrorMessage)
		return
	}

	message := "Usuario actualizado correctamente"
	utils.RespondWithJson(w, http.StatusOK, utils.ApiResponse[UserDto]{
		Success: true,
		Message: &message,
		Data:    updateResult.Data,
	})
}

func (uh *UserHandler) UpdatePassword(w http.ResponseWriter, r *http.Request) {
	updatePassword := UpdatePasswordDto{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&updatePassword); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	authUser, ok := utils.GetAuthUser(r.Context())
	if !ok {
		utils.RespondWithError(w, http.StatusUnauthorized, "Usuario no autorizado")
		return
	}

	userExists := uh.userService.DoesUserExist(r.Context(), authUser.UserId)
	if !userExists {
		utils.RespondWithError(w, http.StatusNotFound, "Usuario no encontrado")
		return
	}

	updateResult := uh.userService.UpdatePassword(r.Context(), authUser.UserId, &updatePassword)
	if !updateResult.Success {
		utils.RespondWithError(w, http.StatusBadRequest, *updateResult.ErrorMessage)
		return
	}

	message := "Contraseña actualizada correctamente"
	utils.RespondWithJson(w, http.StatusOK, utils.ApiResponse[string]{
		Success: true,
		Message: &message,
	})
}
