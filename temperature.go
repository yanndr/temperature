//Package temperature provides Temperature struct and a convert method to deal with different temperature scale.
package temperature

import (
	"errors"
	"math"
)

// ToKelvinConvertible is an interface necessary for temperature to be converted in Kelvin.
type ToKelvinConvertible interface {
	ToKelvin() Kelvin
}

// FromKelvinConvertible is an interface necessary for temperature to be converted from Kelvin.
type FromKelvinConvertible interface {
	FromKelvin(Valuer)
}

// Valuer is an interface of a temperature who exposes a value.
type Valuer interface {
	GetValue() float64
}

// Temperature is a default struct of a temperature.
type Temperature struct {
	Value float64
}

// GetValue returns the value of the temperature.
func (t Temperature) GetValue() float64 {
	return t.Value
}

type (
	// Celsius is a temperature unit.
	Celsius struct{ Temperature }
	// Kelvin is a temperature unit.
	Kelvin struct{ Temperature }
	// Fahrenheit is a temperature unit.
	Fahrenheit struct{ Temperature }
	// Rankine is a temperature unit.
	Rankine struct{ Temperature }
	// Delisle is a temperature unit.
	Delisle struct{ Temperature }
	// Reaumur is a temperature unit.
	Reaumur struct{ Temperature }
)

// NewCelsius returns a Celsius initialized with the value.
func NewCelsius(value float64) Celsius {
	return Celsius{
		Temperature: Temperature{Value: value},
	}
}

// NewKelvin returns a Kelvin initialized with the value.
func NewKelvin(value float64) Kelvin {
	return Kelvin{
		Temperature: Temperature{Value: value},
	}
}

// NewFahrenheit returns a Fahrenheit initialized with the value.
func NewFahrenheit(value float64) Fahrenheit {
	return Fahrenheit{
		Temperature: Temperature{Value: value},
	}
}

// NewRankine returns a Rankine initialized with the value.
func NewRankine(value float64) Rankine {
	return Rankine{
		Temperature: Temperature{Value: value},
	}
}

// NewDelisle returns a Delisle initialized with the value.
func NewDelisle(value float64) Delisle {
	return Delisle{
		Temperature: Temperature{Value: value},
	}
}

// NewReaumur returns a Reaumur initialized with the value.
func NewReaumur(value float64) Reaumur {
	return Reaumur{
		Temperature: Temperature{Value: value},
	}
}

//Round returns a round number of a float64.
func Round(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	val := int(num*output + math.Copysign(0.5, num))
	return float64(val) / output
}

// ErrNilOutputTemperature is an error when the output temperature is nil.
var ErrNilOutputTemperature = errors.New("Output temperature can't be nil")

// ErrNilInputTemperature is an error when the input temperature is nil.
var ErrNilInputTemperature = errors.New("Input temperature can't be nil")

// Convert a temperature to different unit. The output must be a pointer as its value will be updated.
func Convert(input ToKelvinConvertible, output FromKelvinConvertible) error {
	if output == nil {
		return ErrNilOutputTemperature
	}

	if input == nil {
		return ErrNilInputTemperature
	}

	output.FromKelvin(input.ToKelvin())
	return nil
}
