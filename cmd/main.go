package main

import (
	"fmt"

	"github.com/yanndr/temperature"
)

func main() {
	output := temperature.NewCelsius(0)
	k := temperature.NewKelvin(366)
	temperature.Convert(k, &output)
	n := NewUnit{}

	//temperature.Convert(n, &output)

	fmt.Println(output)
	temperature.Convert(output, &n)
	fmt.Println(output)
}

type NewUnit struct {
	temperature.Temperature
}

func (n NewUnit) ToKelvin() temperature.Kelvin {
	return temperature.NewKelvin(n.Value - 5)
}

func (n *NewUnit) FromKelvin(t temperature.Valuer) {
	n.Value = t.GetValue() + 5
}
