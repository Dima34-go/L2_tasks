package unpackage

import (
	"fmt"
)


type Unpacker interface {
	unpack(s string) (string,error)
}
func MainUnPackage2(){
	var str string
	fmt.Scan(&str)
	unpackEscape(str)
}

