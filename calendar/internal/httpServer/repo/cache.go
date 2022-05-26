package repo

import "C"
import (
	"L_2/calendar/internal/Cache"
	"L_2/calendar/internal/httpServer"
	"time"
)

type ChMap struct {
	C Cache.Cache
}

func NewChMap(cache Cache.Cache) *ChMap {
	return &ChMap{C: cache}
}

func (c *ChMap) CreateEvent(ev Cache.EventQuery) (*Cache.EventAnswer,*httpServer.CustomError) {
	//Получение юзера по ID, если юзера не существует - добавляем его
	user,ok := c.C.UsersEvents[ev.UserId]
	if !ok {
		user = new(Cache.User)
		c.C.UsersEvents[ev.UserId] = user
	}

    //Добавление события в массив
    upUser := c.C.UsersEvents[ev.UserId]
	upUser.AllEvents = append(upUser.AllEvents,ev.Event)
	return &Cache.EventAnswer{
		UserId:  ev.UserId,
		EventId: len(upUser.AllEvents),
		Event:   ev.Event,
	},nil
}

func (c *ChMap) UpdateEvent(ev Cache.UpEventQuery) (*Cache.EventAnswer,*httpServer.CustomError) {
	//Получение юзера по ID, если юзера не существует - выдаем ошибку
	user,ok := c.C.UsersEvents[ev.UserId]
	if !ok {
		return nil,httpServer.NewCustomError(400,"Such user is not exist","Such user is not exist")
	}

	//Добавление события в массив
	if len(user.AllEvents) < ev.EventId {
		return nil,httpServer.NewCustomError(400,"Such event is not exist","Such event is not exist")
	}
	user.AllEvents[ev.EventId-1] = ev.Event
	return &Cache.EventAnswer{
		UserId:  ev.UserId,
		EventId: ev.EventId,
		Event:   ev.Event,
	},nil
}

func (c *ChMap) DeleteEvent(userId, eventId int) (*Cache.EventAnswer,*httpServer.CustomError) {
	//Получение юзера по ID, если юзера не существует - выдаем ошибку
	user,ok := c.C.UsersEvents[userId]
	if !ok {
		return nil,httpServer.NewCustomError(400,"Such user is not exist","Such user is not exist")
	}

	//Удаление события из массива
	if len(user.AllEvents) < eventId {
		return nil,httpServer.NewCustomError(400,"Such event is not exist","Such event is not exist")
	}
	deletedEv:=user.AllEvents[eventId-1]
	user.AllEvents = append(user.AllEvents[:eventId-1],user.AllEvents[eventId:]...)
	return &Cache.EventAnswer{
		UserId:  userId,
		EventId: eventId,
		Event:   deletedEv,
	},nil
}

func (c *ChMap) GetEventsForDuration(date time.Time,userId int,dur time.Duration) (*Cache.GetQueryAnswer,*httpServer.CustomError) {
	ans := new(Cache.GetQueryAnswer)
	user,ok := c.C.UsersEvents[userId]
	if !ok {
		return nil,httpServer.NewCustomError(400,"Such user is not exist","Such user is not exist")
	}
	for _,ev := range user.AllEvents {
		if ev.Date.Sub(date) >=0 && ev.Date.Sub(date.Add(dur)) < 0 {
			ans.Events = append(ans.Events,ev)
		}
	}
	return ans,nil
}