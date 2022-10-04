package main

import (
	"context"
	"fmt"
	"reflect"

	"github.com/aws/aws-lambda-go/lambda"
)

type Human struct {
	Name        string `json:"name"`
	Age         uint8  `json:"age"`
	Description string `json:"desc"`
}

// NOTE: `Strum` is equal to `T`. This is 2 methods to declare type paramneter.
type Strum interface {
	~uint | ~string
}

func runeConverter[T ~uint | ~string](input T) []rune {
	strType := reflect.TypeOf("string")
	if kind := reflect.TypeOf(input); kind == strType {
		return []rune(fmt.Sprint(input))
	}

	// Equals: `return []rune{}`
	return *new([]rune)
}

func HandleReq(ctx context.Context, human Human) ([][]rune, error) {
	name := runeConverter(human.Name)
	age := []rune(string(human.Age))
	desc := runeConverter(human.Description)

	runeArr := [][]rune{name, age, desc}
	return runeArr, nil
}

func main() {
	lambda.Start(HandleReq)
}
