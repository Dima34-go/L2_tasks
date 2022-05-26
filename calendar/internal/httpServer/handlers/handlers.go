package handlers

import (
	"L_2/calendar/internal/httpServer/middleware"
	"L_2/calendar/internal/httpServer/service"
	"net/http"
)

type Handlers struct {
	service service.Service
}

func NewHandlers(service service.Service) *Handlers {
	return &Handlers{service: service}
}

func (h *Handlers) InitRouter() *http.Handler{
	//Router
	eventMux:=http.NewServeMux()
	eventMux.HandleFunc("/create_event",h.CreateEvent)
	eventMux.HandleFunc("/update_event",h.UpdateEvent)
	eventMux.HandleFunc("/delete_event",h.DeleteEvent)
	eventMux.HandleFunc("/events_for_day",h.GetEventsForDay)
	eventMux.HandleFunc("/events_for_week",h.GetEventsForWeek)
	eventMux.HandleFunc("/events_for_month",h.GetEventsForMonth)
	//Добавление Middleware
	siteHandler := middleware.LogMiddleWare(eventMux)
	siteHandler = middleware.PanicMiddleWare(siteHandler)
	return &siteHandler
}