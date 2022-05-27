package codeUnderstanding


import (
	"fmt"
)

func test() (x int) {
	defer func() {
		fmt.Println(&x)
		x++
	}()
	x = 1
	fmt.Println(&x)
	return
}


func anotherTest() int {
	var x int
	defer func() {
		x++
	}()
	x = 1
	fmt.Println(&x)
	return x
}


func MainList2() {
	z1 := test()
	fmt.Println(&z1)
	fmt.Println(z1)
	z2 := anotherTest()
	fmt.Println(&z2)
	fmt.Println(z2)

}
