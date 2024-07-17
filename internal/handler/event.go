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

type EventService interface {
	GetEventAuthors(eventId int64) ([]domain.User, error)
	GetEventParticipants(eventId int64) ([]domain.User, error)
	GetById(id int64) (domain.Event, error)
}

type EventHandler struct {
	eventService EventService
}

func NewEventHandler(service EventService) *EventHandler {
	return &EventHandler{eventService: service}
}

func (eventHandler *EventHandler) getEventByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
		util.LogError(err)
		return
	}

	event, err := eventHandler.eventService.GetById(id)
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
	json.NewEncoder(w).Encode(event)
}
