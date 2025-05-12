package main

import (
	"fmt"

	"github.com/strategycomplex/go-geocodio"
)

func main() {

	gc, err := geocodio.New("YOUR_API_KEY")
	if err != nil {
		panic(err)
	}

	result, err := gc.Geocode("42370 Bob Hope Dr, Rancho Mirage, CA")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Geocode Result %v", result)

	resultReverse, err := gc.Reverse(38.9002898, -76.9990361)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Reverse Geocode Result %v", resultReverse)
}
