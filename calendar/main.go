package calendar

import (
	"L_2/calendar/internal/Cache"
	"L_2/calendar/internal/httpServer/handlers"
	"L_2/calendar/internal/httpServer/repo"
	"L_2/calendar/internal/httpServer/service"
	"log"
	"net/http"
)

func StartCalendar() {
	//Хранилище
	ch := Cache.NewCache()

	//Архитектура http сервера
	repos := repo.NewChMap(*ch)
	services:= service.NewHttpService(repos)
	handlersHttp := handlers.NewHandlers(services)
	srv := new(Server)
	if err := srv.Run("8080", handlersHttp.InitRouter()); err != nil && err != http.ErrServerClosed {
		log.Fatalf("error running http server: %v", err)
	}
}