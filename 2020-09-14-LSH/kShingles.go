package main

import (
	"fmt"
	"strings"
)

func main() {
	input := "The quick fox jumped over the lazy dog";
	input = strings.Replace(input, " ", "_", -1)
	shingleLen := 5
	for i, j:=0, shingleLen; j< len(input); i, j = i+1,j+1 {
		if(i == 0) {
			fmt.Print("{")
		} else {
			fmt.Print(",")
		}
		fmt.Printf("\"%s\"",input[i:j])
	}
	fmt.Print("}")
}
