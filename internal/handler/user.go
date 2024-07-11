package handler

import (
	"GamificationEducation/internal/domain"
	errors2 "GamificationEducation/internal/errors"
	"GamificationEducation/internal/util"
	"encoding/json"
	"errors"
	"github.com/go-chi/chi"
	"net/http"
	"strconv"
)

type UserService interface {
	GetById(id int64) (domain.User, error)
	GetAll() ([]domain.User, error)
}

type UserHandler struct {
	userService UserService
}

func NewUserHandler(service UserService) *UserHandler {
	return &UserHandler{userService: service}
}

func (userHandler *UserHandler) getUserById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
		util.LogError(err)
		return
	}

	user, err := userHandler.userService.GetById(id)
	if err != nil {
		if errors.Is(err, errors2.ErrNotFound) {
			w.WriteHeader(http.StatusNotFound)
			response, err := json.Marshal(ErrorResponse{Message: "user not found"})
			if err != nil {
				util.LogError(err)
			}
			w.Write(response)
		} else {
			http.Error(w, "", http.StatusBadRequest)
			util.LogError(err)
		}
		return
	}
	json.NewEncoder(w).Encode(user)
}

func (userHandler *UserHandler) getAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	//clan := r.URL.Query().Get("clan")
	//email := r.URL.Query().Get("email")

	users, err := userHandler.userService.GetAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response, err := json.Marshal(ErrorResponse{Message: err.Error()})
		if err != nil {
			util.LogError(err)
		}
		w.Write(response)
		return
	}
	response, err := json.Marshal(users)
	if err != nil {
		util.LogError(err)
	}
	w.Write(response)
}
