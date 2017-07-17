//Package temperature provides Temperature struct to deal with different temperature unit.
package temperature

import (
	"fmt"
	"math"
)

//Unit represent a temperature unit.
type Unit struct {
	Name    string
	Symbole string
}

//Enum for Unit.
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
	return fmt.Sprintf("%.2f°%s", ToFixed(t.Value, 2), t.Unit.Symbole)
}

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func ToFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}
