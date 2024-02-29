package main

import (
	"code-generator/generator"
	"fmt"
)

func main() {
	generator := generator.Generator{}
	fmt.Println(generator.Generate())
}
