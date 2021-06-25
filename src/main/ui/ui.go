package ui

import (
	"bytes"
	"fmt"
)

const (
	TAB string = "<TAB>"
)

func Render(uiArray []string) {
	for _, v := range uiArray {
		if v == TAB {
			var buffer bytes.Buffer
			buffer.WriteString(" + ")
			for i := 0; i < 200; i++ {
				buffer.WriteString("=")
			}
			buffer.WriteString(" + ")
			fmt.Println(buffer.String())
		} else {
			var buffer bytes.Buffer
			buffer.WriteString(" | ")
			buffer.WriteString(v)
			for i := 0; i < 200-len(v); i++ {
				buffer.WriteString(" ")
			}
			buffer.WriteString(" | ")
			fmt.Println(buffer.String())
		}
	}
}
