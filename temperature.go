//Package temperature provides Temperature struct to deal with different temperature unit.
package temperature

import (
	"fmt"
	"math"
)

//Unit represent a temperature unit.
type Unit struct {
	Name   string
	Symbol string
}

//Predefined Unit.
var (
	Kelvin     = Unit{"Kelvin", `K`}
	Celsius    = Unit{"Celsius", `C`}
	Fahrenheit = Unit{"Fahrenheit", `F`}
	Rankine    = Unit{"Rankine", `R`}
	Delisle    = Unit{"Delisle", `D`}
)

//Temperature represent a temperature with a Value and a Unit.
type Temperature struct {
	Value float64
	Unit  Unit
}

//NewTemperature allocate a Temperature with a Value and a Unit.
func NewTemperature(value float64, unit Unit) *Temperature {
	t := &Temperature{Value: value, Unit: unit}
	return t
}

func (t *Temperature) String() string {
	return fmt.Sprintf("%.2f°%s", Round(t.Value, 2), t.Unit.Symbol)
}

//Round returns a round number of a float64.
func Round(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	val := int(num*output + math.Copysign(0.5, num))
	return float64(val) / output
}
