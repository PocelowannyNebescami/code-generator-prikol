package main

import (
	"code-generator/generator"
	"fmt"
)

func main() {
	config := generator.Config{Length: 24}
	generator, err := generator.Init(config)
	if err != nil {
		fmt.Println("Generator is not set properly:", err)
	}

	fmt.Println("code:", generator.Generate())
}
