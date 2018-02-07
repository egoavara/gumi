package main

import (
	"fmt"
	"github.com/iamGreedy/gumi"
)

func main() {
	text := "wfewagawe gregre waefgawegwea"
	fmt.Printf("'%s'\n", text)
	text = gumi.StringControlBackSpace(text)
	fmt.Printf("'%s'\n", text)
	text = gumi.StringControlBackSpace(text)
	fmt.Printf("'%s'\n", text)
	text = gumi.StringControlBackSpace(text)
	fmt.Printf("'%s'\n", text)
	fmt.Println("test"[:4])
	fmt.Println("test"[:0])

}