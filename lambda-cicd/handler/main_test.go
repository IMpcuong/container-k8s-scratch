package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestTypeCompare(t *testing.T) {
	input := "Dude"
	strType := reflect.TypeOf("string")
	fmt.Println(strType)

	if kind := reflect.TypeOf(input); kind == strType {
		fmt.Println([]rune(fmt.Sprint(input)))
		return
	}
	t.Fail()
}
