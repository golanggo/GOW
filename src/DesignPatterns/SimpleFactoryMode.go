package DesignPatterns

import (
	"fmt"
)

type IProduct interface {
	Show()
}

type ProductA struct {
}

type ProductB struct {
}

func (P *ProductA) Show() {
	fmt.Println("ProductA")
}

func (P *ProductB) Show() {
	fmt.Println("ProductB")
}

type Factory struct {
}

type IFactory interface {
	GetP(patter string) IProduct
}

func (F *Factory) GetP(patter string) IProduct {
	switch patter {
	case "A":
		return &ProductA{}
	case "B":
		return &ProductB{}
	default:
		return nil
	}
}
