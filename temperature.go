//Package temperature provides Temperature struct to deal with different temperature scale.
package temperature

import (
	"math"
)

type TemperatureConvertible interface {
	FromKelvin(TemperatureConvertible) TemperatureConvertible
	ToKelvin() TemperatureConvertible
	Value() float64
}

type Temperature struct {
	value float64
}

type (
	celsius struct{ Temperature }
	kelvin  struct{ Temperature }
)

func (c celsius) ToKelvin() TemperatureConvertible {
	return kelvin{Temperature: Temperature{value: c.Value() + 273.15}}
}
func (c celsius) FromKelvin(v TemperatureConvertible) TemperatureConvertible {
	return celsius{Temperature: Temperature{value: v.Value() - 273.15}}
}
func (t Temperature) Value() float64 {
	return t.value
}

func (k kelvin) ToKelvin() TemperatureConvertible {
	return k
}
func (k kelvin) FromKelvin(v TemperatureConvertible) TemperatureConvertible {
	return kelvin{Temperature: Temperature{value: v.Value()}}
}

//Scale represent a temperature scale.
type Scale struct {
	Name   string
	Symbol string
}

// Predefined scales.
var (
	Kelvin     = Scale{"Kelvin", `K`}
	Celsius    = Scale{"Celsius", `C`}
	Fahrenheit = Scale{"Fahrenheit", `F`}
	Rankine    = Scale{"Rankine", `R`}
	Delisle    = Scale{"Delisle", `D`}
	Reaumur    = Scale{"Reaumur", "Re"}
)

//Temperature represent a temperature with a Value and a scale.
// type Temperature struct {
// Predefined scales.
// 	Scale Scale
// }

//New allocate a Temperature with a Value and a scale.
// func New(value float64, scale Scale) *Temperature {
// 	t := &Temperature{Value: value, Scale: scale}
// 	return t
// }

// func (t *Temperature) String() string {
// 	return fmt.Sprintf("%v°%s", Round(t.Value, 2), t.Scale.Symbol)
// }

//Round returns a round number of a float64.
func Round(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	val := int(num*output + math.Copysign(0.5, num))
	return float64(val) / output
}
