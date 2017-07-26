package main

import (
	"fmt"

	"github.com/yanndr/temperature"
	"github.com/yanndr/temperature/converter"
)

func main() {
	t := temperature.NewTemperature(0, temperature.Celsius)

	tc := converter.NewTemperatureConverter()

	tc.Converters[NewUnit] = &NewUnitConverter{}

	temp, _ := tc.Convert(*t, NewUnit)
	fmt.Println(temp)

	t = temperature.NewTemperature(25.3, NewUnit)

	temp, _ = tc.Convert(*t, temperature.Fahrenheit)
	fmt.Println(temp)
}

var NewUnit = temperature.Unit{"NewUnit", `Y`}

type NewUnitConverter struct {
}

func (*NewUnitConverter) ToKelvin(t float64) float64 {
	return t + 200
}

func (*NewUnitConverter) FromKelvin(t float64) float64 {
	return t - 200
}
