package unpackage

import (
	"bytes"
	"errors"
	"strconv"
	"unicode"
)

type JustString struct {}

func (j JustString) unpack(str string) (string, error) {
	var num string
	var buf bytes.Buffer
	if len(str) == 0 {
		return "",nil
	}
	if unicode.IsDigit(rune(str[0])) {
		return "",errors.New("Uncorrect str")
	}
	//первый встреченный символ
	ls:=str[0]
	for i:=1;i<len(str);i++ {
		if unicode.IsDigit(rune(str[i])) {
			num+=string(str[i])
		}else{
			n,_:=strconv.Atoi(num)
			if n == 0{ n++ }
			buf.Write(bytes.Repeat([]byte{ls},n))
			num = ""
			ls=str[i]
		}
	}
	n,_:=strconv.Atoi(num)
	if n == 0{ n++ }
	buf.Write(bytes.Repeat([]byte{ls},n))
	return buf.String(),nil
}
