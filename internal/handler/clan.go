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

type ClanService interface {
	GetAll() ([]domain.Clan, error)
	GetByName(name string) (domain.Clan, error)
	GetMembers(id int64) ([]domain.User, error)
	GetById(id int64) (domain.Clan, error)
	GetMinClan() (domain.Clan, error)
}

type ClanHandler struct {
	clanService ClanService
}

func NewClanHandler(service ClanService) *ClanHandler {
	return &ClanHandler{clanService: service}
}

func (clanHandler *ClanHandler) getAllClans(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	clans, err := clanHandler.clanService.GetAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		util.LogError(err)
		return
	}
	response, err := json.Marshal(clans)
	if err != nil {
		util.LogError(err)
	}
	w.Write(response)
}

func (clanHandler *ClanHandler) getClanById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
		util.LogError(err)
		return
	}
	clan, err := clanHandler.clanService.GetById(id)
	if err != nil {
		if errors.Is(err, errors2.ErrNotFound) {
			w.WriteHeader(http.StatusNotFound)
			response, err := json.Marshal(ErrorResponse{Message: errors2.ErrNotFound.Error()})
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
	json.NewEncoder(w).Encode(clan)
}
