package main

import (
	"fmt"

	"github.com/yanndr/temperature"
)

var myUnit = temperature.Unit{
	Text:       "Y",
	FromKelvin: func(v float64) float64 { return v + 710 },
	ToKelvin:   func(v float64) float64 { return v - 710 },
}

func main() {
	c := temperature.NewCelsius(0)
	k := temperature.NewKelvin(0)
	f := temperature.NewFahrenheit(0)
	n := temperature.NewTemperature(0, myUnit)

	temperature.Convert(c, n)
	temperature.Convert(c, k)
	temperature.Convert(k, f)
	fmt.Println(c)
	fmt.Println(k)
	fmt.Println(f)
	fmt.Println(n)
}
