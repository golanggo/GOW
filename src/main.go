package main

import (
	"./DesignPatterns"
	"fmt"
)

func main() {

	factory := DesignPatterns.Factory{}
	factory.GetP("B").Show()
	fmt.Println("test")
}
