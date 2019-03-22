package main

import (
	"DataStruct"
	"fmt"
)

func main() {
	var list DataStruct.List
	list.Insert(0, 15)
	list.Insert(1, 12)
	list.Insert(2, 16)
	a := list.Get(0)
	b := list.Get(1)
	c := list.Get(2)

	fmt.Println(a, b, c, list.Lenght)

	//list.Del(3)
	fmt.Println(list, list.Lenght)

	list.Sort()
	fmt.Println(list)
	fmt.Println(list)
	fmt.Println(list)
	fmt.Println(list)

}
