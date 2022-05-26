package service

import (
	"L_2/calendar/internal/Cache"
	"L_2/calendar/internal/httpServer"
	"time"
)

type Service interface {
	CreateEvent(buf Cache.EventBuffer) (*Cache.EventAnswer,*httpServer.CustomError)
	UpdateEvent(buf Cache.UpEventBuffer) (*Cache.EventAnswer,*httpServer.CustomError)
	DeleteEvent(userId, eventId string) (*Cache.EventAnswer,*httpServer.CustomError)
	GetEventsForDuration(dateStr,user string,duration time.Duration) (*Cache.GetQueryAnswer,*httpServer.CustomError)
}
