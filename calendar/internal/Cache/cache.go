package Cache


func NewCache() *Cache {
	mp := make(map[int]*User)
	allEv:=make( []Event,0)
	allEv = append(allEv,Event{
		Name:        "Day Z",
		Description: "The best day",
	})
	mp[1] = &User{AllEvents: allEv}
	return &Cache{UsersEvents: mp}
}