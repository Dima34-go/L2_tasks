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
	js := EscapeSequence{}
	newStr,err :=js.unpack(str)
	if err != nil {
		fmt.Println(err.Error())
	}else {
		fmt.Println(newStr)
	}
}

