package main

import (
	"fmt"
	"github.com/iamGreedy/gumi/glumi"
)

func main() {
	fmt.Println(glumi.DefaultShader.Vertex)
	fmt.Println(glumi.DefaultShader.Fragment)
}