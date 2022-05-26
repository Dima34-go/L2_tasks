package strategy

type Strategy interface {
	Route(startPoint, endPoint int)
}