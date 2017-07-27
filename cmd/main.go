package main

import (
	"fmt"

	"github.com/yanndr/temperature"
)

func main() {
	c := temperature.NewCelsius(0)
	k := temperature.NewKelvin(366)
	n := newUnit{}

	temperature.Convert(c, &n)
	temperature.Convert(c, &k)
	fmt.Println(c)
	fmt.Println(k)
	fmt.Println(n)
}

type newUnit struct {
	temperature.Temperature
}

func (n newUnit) ToKelvin() temperature.Kelvin {
	return temperature.NewKelvin(n.Value - 5)
}

func (n *newUnit) FromKelvin(t temperature.Valuer) {
	n.Value = t.GetValue() + 5
}

func (n newUnit) String() string {
	return fmt.Sprintf("%v °Y", temperature.Round(n.Value, 2))
}
