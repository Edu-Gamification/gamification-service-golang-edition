package handler

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

const (
	userPrefix  = "/api/users"
	clanPrefix  = "/api/clans"
	eventPrefix = "/api/events"
)

func InitRouter() *chi.Mux {
	r := chi.NewMux()
	r.Use(middleware.Recoverer)
	return r
}

func (userHandler *UserHandler) InitUserRoutes(r *chi.Mux) {
	r.Get(userPrefix+"/{id}", userHandler.getUserById)
	r.Get(userPrefix+"/", userHandler.getAllUsers)
}

func (clanHandler *ClanHandler) InitClanRoutes(r *chi.Mux) {
	r.Get(clanPrefix, clanHandler.getAllClans)
	r.Get(clanPrefix+"/{id}", clanHandler.getClanById)
}

func (eventHandler *EventHandler) InitEventRoutes(r *chi.Mux) {
	//r.Get(eventPrefix, eventHandler.)
	r.Get(eventPrefix+"/{id}", eventHandler.getEventByID)
}
