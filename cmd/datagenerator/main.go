package main

import (
	"fmt"
	"github.com/asaphin/all-databases-go/internal/datagenerator"
)

func main() {
	fmt.Println(datagenerator.New().VR().UnitedStatesAddress())
}
