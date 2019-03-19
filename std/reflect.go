package main

import (
	"fmt"
	"reflect"
)

type tempController struct{}

func (this *tempController) Name(name string) {

	fmt.Printf("tempController method Name , %s \n", name)

}

func main() {

	className := "tempController"

	classType := reflect.TypeOf(className)

	fmt.Printf("111111, %v\n", classType)

}
