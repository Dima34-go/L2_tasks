package service

import (
	"L_2/calendar/internal/Cache"
	"L_2/calendar/internal/httpServer"
	"L_2/calendar/internal/httpServer/repo"
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

type HttpService struct {
	repo  repo.Repo
}

func NewHttpService(repo repo.Repo) *HttpService {
	return &HttpService{repo: repo}
}

const layout = "2006-01-02"

func (s *HttpService) CreateEvent(buf Cache.EventBuffer) (*Cache.EventAnswer,*httpServer.CustomError) {
	//Парсинг UserId
	userId, err := strconv.Atoi(buf.UserId)
	if err != nil {
		return nil, httpServer.NewCustomError(400,httpServer.BadRequest,err.Error())
	}
	//Парсинг Date
	date,err :=time.Parse(layout,buf.Date)
	if err != nil {
		return nil, httpServer.NewCustomError(400,httpServer.BadRequest,err.Error())
	}
	//Парсинг тела запроса
	fmt.Println(string(buf.Body))
	var ev Cache.Event
	err = json.Unmarshal(buf.Body,&ev)
	if err != nil {
		return nil, httpServer.NewCustomError(400,httpServer.BadRequest,err.Error())
	}
	ev.Date = date
	//Отправление в БД
	return s.repo.CreateEvent(Cache.EventQuery{
		UserId: userId,
		Date:   date,
		Event:  ev,
	})
}

func (s *HttpService) UpdateEvent(buf Cache.UpEventBuffer) (*Cache.EventAnswer,*httpServer.CustomError) {
	//Парсинг UserId
	userId, err := strconv.Atoi(buf.UserId)
	if err != nil {
		return nil, httpServer.NewCustomError(400,httpServer.BadRequest,err.Error())
	}
	//Парсинг EventId
	evId, err := strconv.Atoi(buf.EventId)
	if err != nil {
		return nil, httpServer.NewCustomError(400,httpServer.BadRequest,err.Error())
	}
	//Парсинг Date
	date,err :=time.Parse(layout,buf.Date)
	if err != nil {
		return nil, httpServer.NewCustomError(400,httpServer.BadRequest,err.Error())
	}
	//Парсинг тела запроса
	fmt.Println(string(buf.Body))
	var ev Cache.Event
	err = json.Unmarshal(buf.Body,&ev)
	if err != nil {
		return nil, httpServer.NewCustomError(400,httpServer.BadRequest,err.Error())
	}
	ev.Date = date
	//Отправление в БД
	return s.repo.UpdateEvent(Cache.UpEventQuery{
		UserId:  userId,
		EventId: evId,
		Date:    date,
		Event:   ev,
	})
}

func (s *HttpService) DeleteEvent(Id, eventId string) (*Cache.EventAnswer,*httpServer.CustomError) {
	//Парсинг UserId
	userId, err := strconv.Atoi(Id)
	if err != nil {
		return nil, httpServer.NewCustomError(400,httpServer.BadRequest,err.Error())
	}
	//Парсинг EventId
	evId, err := strconv.Atoi(eventId)
	if err != nil {
		return nil, httpServer.NewCustomError(400,httpServer.BadRequest,err.Error())
	}
	return s.repo.DeleteEvent(userId,evId)
}

func (s HttpService) GetEventsForDuration(dateStr, user string,dur time.Duration) (*Cache.GetQueryAnswer,*httpServer.CustomError) {
	//Парсинг UserId
	userId, err := strconv.Atoi(user)
	if err != nil {
		return nil, httpServer.NewCustomError(400,httpServer.BadRequest,err.Error())
	}
	//Парсинг Date
	date,err :=time.Parse(layout,dateStr)
	if err != nil {
		return nil, httpServer.NewCustomError(400,httpServer.BadRequest,err.Error())
	}
	return s.repo.GetEventsForDuration(date,userId,dur)
}