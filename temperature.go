//Package temperature provides Temperature struct to deal with different temperature unit.
package temperature

import (
	"fmt"
)

//Unit represent a temperature unit.
type Unit string

//Enum for Unit.
const (
	Kelvin     Unit = "K"
	Celsius         = "C"
	Fahrenheit      = "F"
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
	return fmt.Sprintf("%.2f° %s", round(t.Value, 0.05), t.Unit)
}

func round(x, unit float64) float64 {
	return float64(int64(x/unit+0.5)) * unit
}
