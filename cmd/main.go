package main

import (
	"fmt"
	"log"

	"github.com/yanndr/temperature"
)

type newUnit temperature.Unit

const myUnit = newUnit("Y")

func (newUnit) ToKelvin(v float64) float64 {
	return v + 710
}

func (newUnit) FromKelvin(v float64) temperature.Temperature {
	return temperature.New(v-710, myUnit)
}

func main() {
	c := temperature.New(0, temperature.Celsius)

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
