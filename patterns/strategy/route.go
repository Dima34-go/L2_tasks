package strategy

import "fmt"

type RouteStrategy struct {}

func (r *RouteStrategy) Route(startPoint, endPoint int) {
	avgSpeed := 30
	trafficJam := 2
	total := endPoint - startPoint
	totalTime := trafficJam * total / avgSpeed
	fmt.Printf(
		"Road A:[%d], B:[%d] Avg speed: [%d] Traffic jam: [%d] Total: [%d] Total time: [%d] min \n",
		startPoint,
		endPoint,
		avgSpeed,
		trafficJam,
		total,
		totalTime,
	)
}
