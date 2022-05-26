package Cache

import "time"

//Событие

type Event struct {
	Date        time.Time `json:"time,omitempty"`
	Name        string    `json:"name,omitempty"`
	Description string    `json:"description,omitempty"`
}

type EventQuery struct {
	UserId int
	Date   time.Time
	Event  Event
}

type UpEventQuery struct {
	UserId  int
	EventId int
	Date    time.Time
	Event   Event
}

type EventBuffer struct {
	UserId string
	Date   string
	Body   []byte
}

type UpEventBuffer struct {
	UserId  string
	EventId string
	Date    string
	Body    []byte
}

type EventAnswer struct {
	UserId  int
	EventId int
	Event   Event
}
type GetQueryAnswer struct {
	Events []Event
}
//Все события, относящиеся к одному Пользователю

type User struct {
	AllEvents []Event
}

type Cache struct {
	UsersEvents map[int]*User
}
