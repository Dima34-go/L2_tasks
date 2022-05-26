package repo

import (
	"L_2/calendar/internal/Cache"
	"L_2/calendar/internal/httpServer"
	"time"
)

type Repo interface {
	CreateEvent(ev Cache.EventQuery) (*Cache.EventAnswer,*httpServer.CustomError)
	UpdateEvent(ev Cache.UpEventQuery) (*Cache.EventAnswer,*httpServer.CustomError)
	DeleteEvent(userId, eventId int) (*Cache.EventAnswer,*httpServer.CustomError)
	GetEventsForDuration(date time.Time,userId int,dur time.Duration) (*Cache.GetQueryAnswer,*httpServer.CustomError)
}