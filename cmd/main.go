package main

import (
	"fmt"

	"github.com/yanndr/temperature"
	"github.com/yanndr/temperature/converter"
)

func main() {
	t := temperature.NewTemperature(0, temperature.Celsius)

	tc := converter.NewTemperatureConverter()

	//fmt.Println(tc.Convert(*t, temperature.Kelvin))

	fmt.Println(tc.Convert(*t, temperature.Fahrenheit))
}
