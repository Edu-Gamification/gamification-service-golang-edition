package main

import (
	_ "GamificationEducation/cmd/docs"
	"GamificationEducation/internal/config"
	"GamificationEducation/internal/db"
	"GamificationEducation/internal/handler"
	"GamificationEducation/internal/util/inits"
	httpSwagger "github.com/swaggo/http-swagger"
	"log"
	"net/http"
)

// @title Gamification Education App API
// @version 1.0
// @description API for Gamification App

// @host localhost:8080
// @BasePath /

func main() {
	cfg, err := config.ParseFromYaml()
	if err != nil {
		log.Fatal("Error parsing config")
	}

	db, err := db.NewDB(&cfg.DB)
	if err != nil {
		log.Fatal("Error create database connection")
	}
	defer db.Close()

	router := handler.InitRouter()

	inits.UserInit(db, router)
	inits.ClanInit(db, router)
	router.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json")))

	err = http.ListenAndServe(":8080", router)
	if err != nil {
		return
	}
}
