package strategy

import "fmt"

type PublicTransportStrategy struct {}

func (r *PublicTransportStrategy) Route(startPoint, endPoint int) {
	avgSpeed := 40
	total := endPoint - startPoint
	totalTime := total / avgSpeed
	fmt.Printf(
		"Road A:[%d], B:[%d] Avg speed: [%d] Total: [%d] Total time: [%d] min \n",
		startPoint,
		endPoint,
		avgSpeed,
		total,
		totalTime,
	)
}