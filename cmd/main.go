package main

import (
	"fmt"
	"log"

	"github.com/yanndr/temperature"
)

var myUnit = temperature.Unit{
	Text:       "Y",
	FromKelvin: func(v float64) float64 { return v + 710 },
	ToKelvin:   func(v float64) float64 { return v - 710 },
}

func main() {
	c := temperature.NewCelsius(0)

	n, err := temperature.Convert(c, myUnit)
	if err != nil {
		log.Fatal(err)
	}
	k, err := temperature.Convert(c, temperature.Kelvin)
	if err != nil {
		log.Fatal(err)
	}
	f, err := temperature.Convert(k, temperature.Fahrenheit)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(c)
	fmt.Println(k)
	fmt.Println(f)
	fmt.Println(n)
}
