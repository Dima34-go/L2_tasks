package handlers

import (
	"L_2/calendar/internal/Cache"
	"L_2/calendar/internal/httpServer"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)


func (h *Handlers) CreateEvent(w http.ResponseWriter,r *http.Request)  {
	//Получение user_id
	id := r.URL.Query().Get("user_id")
	//Получение date
	date := r.URL.Query().Get("date")
	//Получение информации о событии
	buf,err := ioutil.ReadAll(r.Body)
	if err != nil {
		httpServer.WriteError(w,&httpServer.CustomError{
			Code:    500,
			Message: httpServer.InternalError,
			ErrDesc: err.Error(),
		})
		return
	}
	ev, cErr := h.service.CreateEvent(Cache.EventBuffer{
		UserId: id,
		Date:   date,
		Body:   buf,
	})
	if cErr != nil {
		httpServer.WriteError(w,cErr)
		return
	}
	ans,err := json.Marshal(map[string]interface{}{
		"result": ev,
	})
	if err != nil {
		httpServer.WriteError(w,&httpServer.CustomError{
			Code:    500,
			Message: httpServer.InternalError,
			ErrDesc: err.Error(),
		})
		return
	}
	fmt.Fprintf(w,"%s",string(ans))
}

func (h *Handlers) UpdateEvent(w http.ResponseWriter,r *http.Request) {
	id := r.URL.Query().Get("user_id")
	date := r.URL.Query().Get("date")
	evId := r.URL.Query().Get("event_id")
	//Получение информации о событии
	buf,err := ioutil.ReadAll(r.Body)
	if err != nil {
		httpServer.WriteError(w,&httpServer.CustomError{
			Code:    500,
			Message: httpServer.InternalError,
			ErrDesc: err.Error(),
		})
		return
	}
	ev, cErr := h.service.UpdateEvent(Cache.UpEventBuffer{
		UserId:  id,
		EventId: evId,
		Date:    date,
		Body:    buf,
	})
	if cErr != nil {
		httpServer.WriteError(w,cErr)
		return
	}
	ans,err := json.Marshal(map[string]interface{}{
		"result": ev,
	})
	if err != nil {
		httpServer.WriteError(w,&httpServer.CustomError{
			Code:    500,
			Message: httpServer.InternalError,
			ErrDesc: err.Error(),
		})
		return
	}
	fmt.Fprintf(w,"%s",string(ans))
}

func (h *Handlers) DeleteEvent(w http.ResponseWriter,r *http.Request) {
	//Получение user_id
	id := r.URL.Query().Get("user_id")
	evId := r.URL.Query().Get("event_id")
	ev, cErr := h.service.DeleteEvent(id,evId)
	if cErr != nil {
		httpServer.WriteError(w,cErr)
		return
	}
	ans,err := json.Marshal(map[string]interface{}{
		"result": ev,
	})
	if err != nil {
		httpServer.WriteError(w,&httpServer.CustomError{
			Code:    500,
			Message: httpServer.InternalError,
			ErrDesc: err.Error(),
		})
		return
	}
	fmt.Fprintf(w,"%s",string(ans))
}

func (h *Handlers) GetEventsForDay(w http.ResponseWriter,r *http.Request) {
	id := r.URL.Query().Get("user_id")
	date := r.URL.Query().Get("date")
	ev,cErr := h.service.GetEventsForDuration(date,id,24*time.Hour)
	if cErr != nil {
		httpServer.WriteError(w,cErr)
		return
	}
	ans,err := json.Marshal(map[string]interface{}{
		"result": ev,
	})
	if err != nil {
		httpServer.WriteError(w,&httpServer.CustomError{
			Code:    500,
			Message: httpServer.InternalError,
			ErrDesc: err.Error(),
		})
		return
	}
	fmt.Fprintf(w,"%s",string(ans))
}

func (h *Handlers) GetEventsForWeek(w http.ResponseWriter,r *http.Request) {
	id := r.URL.Query().Get("user_id")
	date := r.URL.Query().Get("date")
	ev,cErr := h.service.GetEventsForDuration(date,id,7*24*time.Hour)
	if cErr != nil {
		httpServer.WriteError(w,cErr)
		return
	}
	ans,err := json.Marshal(map[string]interface{}{
		"result": ev,
	})
	if err != nil {
		httpServer.WriteError(w,&httpServer.CustomError{
			Code:    500,
			Message: httpServer.InternalError,
			ErrDesc: err.Error(),
		})
		return
	}
	fmt.Fprintf(w,"%s",string(ans))
}

func (h *Handlers) GetEventsForMonth(w http.ResponseWriter,r *http.Request) {
	id := r.URL.Query().Get("user_id")
	date := r.URL.Query().Get("date")
	ev,cErr := h.service.GetEventsForDuration(date,id,30*24*time.Hour)
	if cErr != nil {
		httpServer.WriteError(w,cErr)
		return
	}
	ans,err := json.Marshal(map[string]interface{}{
		"result": ev,
	})
	if err != nil {
		httpServer.WriteError(w,&httpServer.CustomError{
			Code:    500,
			Message: httpServer.InternalError,
			ErrDesc: err.Error(),
		})
		return
	}
	fmt.Fprintf(w,"%s",string(ans))
}
