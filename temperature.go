//Package temperature provides Temperature struct to deal with different temperature scale.
package temperature

import (
	"fmt"
	"math"
)

//Scale represent a temperature scale.
type Scale struct {
	Name   string
	Symbol string
}

//Predefined scales.
var (
	Kelvin     = Scale{"Kelvin", `K`}
	Celsius    = Scale{"Celsius", `C`}
	Fahrenheit = Scale{"Fahrenheit", `F`}
	Rankine    = Scale{"Rankine", `R`}
	Delisle    = Scale{"Delisle", `D`}
	Reaumur    = Scale{"Reaumur", "Re"}
)

//Temperature represent a temperature with a Value and a scale.
type Temperature struct {
	Value float64
	Scale Scale
}

//New allocate a Temperature with a Value and a scale.
func New(value float64, scale Scale) *Temperature {
	t := &Temperature{Value: value, Scale: scale}
	return t
}

func (t *Temperature) String() string {
	return fmt.Sprintf("%.2f°%s", Round(t.Value, 2), t.Scale.Symbol)
}

//Round returns a round number of a float64.
func Round(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	val := int(num*output + math.Copysign(0.5, num))
	return float64(val) / output
}
