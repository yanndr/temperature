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
}

var NewUnit = temperature.Unit{"NewUnit", `Y`}

type NewUnitConverter struct {
}

func (*NewUnitConverter) ToKelvin(t temperature.Temperature) *temperature.Temperature {
	return temperature.NewTemperature(t.Value+200, temperature.Kelvin)
}

func (*NewUnitConverter) FromKelvin(t temperature.Temperature) *temperature.Temperature {
	return temperature.NewTemperature(t.Value-200, NewUnit)
}
