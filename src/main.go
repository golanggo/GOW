package main

import (
	"./DataStruct"
	"./DesignPatterns"
	"fmt"
	"os"
)

func main() {

	//字定义错误
	f, err := os.Open("C:/Users/Administrator/Desktop/aaa.txt")
	fmt.Println(err)
	if err != nil {
		fmt.Println(err.Error())
	}
	f.Close()
	//fmt.Println(f.Name())

	//数组，切片，映射map
	ary := []int{} //空数组
	fmt.Println(ary)

	var nilSic []int       // nil 切片
	empt := make([]int, 2) //空切片

	if nilSic == nil {
		fmt.Println("var nilSic is nil")
	}

	if empt == nil {
		fmt.Println("var empt is nil")
		//不是nil 切片
	}

	fmt.Println(nilSic)
	fmt.Println(empt)
	//sic := make([]int,0) //fmt.Println(mary)

	//线性存储-顺序存储
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

	//简单工厂模式
	factory := DesignPatterns.Factory{}
	factory.GetP("B").Show()
	fmt.Println("test")
}
