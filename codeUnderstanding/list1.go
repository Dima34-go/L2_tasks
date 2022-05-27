package codeUnderstanding

import (
"fmt"
)

func MainList1() {
	a := [5]int{76, 77, 78, 79, 80}
	var b []int = a[0:5]
	b = append(b,5)
	a[0] = 4
	fmt.Println(b)
	fmt.Println(a)
}
