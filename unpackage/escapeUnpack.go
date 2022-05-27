package unpackage

import (
	"bytes"
	"strconv"
	"unicode"
)

type EscapeSequence struct {}

func (e EscapeSequence) unpack(str string) (string, error) {
	var buf bytes.Buffer
	for i:=0;i<len(str);i++ {
		if str[i] != '\\' {
			buf.Write([]byte{str[i]})
		} else {
			//попали на / , считываем следующий номер
			var sumb byte
			if i+1<len(str){
				sumb = str[i+1]
			}else{
				buf.Write([]byte{str[i]})
				break
			}
			//считываем число следующее за ним
			num := ""
			for i = i+2;i<len(str);i++{
				if unicode.IsDigit(rune(str[i])) {
					num+=string(str[i])
				}else{
					break
				}
			}
			i--
			n,_ :=strconv.Atoi(num)
			if n == 0 {
				n++
			}
			buf.Write(bytes.Repeat([]byte{sumb},n))

		}
	}
	return buf.String(),nil
}
