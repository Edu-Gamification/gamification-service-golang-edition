package inits

import (
	"GamificationEducation/internal/handler"
	"GamificationEducation/internal/repository"
	"GamificationEducation/internal/service"
	"database/sql"
	"github.com/go-chi/chi"
)

func UserInit(db *sql.DB, router *chi.Mux) {
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)
	userHandler.InitUserRoutes(router)
}

func ClanInit(db *sql.DB, router *chi.Mux) {
	clanRepository := repository.NewClanRepository(db)
	clanService := service.NewClanService(clanRepository)
	clanHandler := handler.NewClanHandler(clanService)
	clanHandler.InitClanRoutes(router)
}
