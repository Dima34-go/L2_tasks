package strategy


func MainStrategy() {
	start := 10
	end := 100
	strategies := []Strategy{ &PublicTransportStrategy{},&RouteStrategy{},&WalkStrategy{}}
	nav := Navigator{}
	for _, strategy := range strategies {
		nav.SetStrategy(strategy)
		nav.Route(start,end)
	}
}