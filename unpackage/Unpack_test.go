package unpackage

import (
	"fmt"
	"testing"
)

type TestStruct struct {
	input  string
	output string
}
func TestUnpack(t *testing.T)  {
	testTable :=[]TestStruct{
		{
			input:  "qwe\\4\\5",
			output: "qwe45",
		},
		{
			input:  "qwe\\45",
			output: "qwe44444",
		},
		{
			input:  "qwe\\\\5",
			output: "qwe\\\\\\\\\\",
		},
		{
			input:  "qwe\\\\\\",
			output: "qwe\\\\",
		},
	}
	es := EscapeSequence{}
	for i, test := range testTable {
		str,err := es.unpack(test.input)
		if str == test.output && err == nil {
			fmt.Printf("Test %d  was successfully passed\n",i)
		}else{
			t.Errorf("Test %d failed, expected %v, %v received %v, %v",i,test.output,nil,str,err)
		}
	}
}