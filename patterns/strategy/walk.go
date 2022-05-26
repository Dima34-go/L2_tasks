package strategy

import "fmt"

type WalkStrategy struct {}

func (r *WalkStrategy) Route(startPoint, endPoint int) {
	avgSpeed := 4
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
