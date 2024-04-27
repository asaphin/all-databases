package main

import (
	"fmt"
	"github.com/asaphin/all-databases-go/internal/datagenerator"
)

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(datagenerator.New().VR().Address().String())
	}
}
