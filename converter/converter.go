//Package converter provides converter to convert temperature units.
package converter

import (
	"errors"
	"fmt"

	"github.com/yanndr/temperature"
)

//Converter is the interface implemented by temperature converter.
type Converter interface {
	//Convert  return a temperature with the new temperature unit.
	Convert(temperature.Temperature, temperature.Unit) (*temperature.Temperature, error)
}

//TemperatureConverter is the default implemtation of a temperature converter.
type TemperatureConverter struct {
	Converters map[temperature.Unit]KelvinConverter
}

//NewTemperatureConverter allocate a new TemperatureConverter.
func NewTemperatureConverter() *TemperatureConverter {
	tc := &TemperatureConverter{}
	tc.Converters = make(map[temperature.Unit]KelvinConverter)
	tc.Converters[temperature.Celsius] = &celsiusKelvinConverter{}
	tc.Converters[temperature.Fahrenheit] = &fahrenheitKelvinConverter{}
	tc.Converters[temperature.Rankine] = &rankineKelvinConverter{}
	tc.Converters[temperature.Delisle] = &delisleKelvinConverter{}
	return tc
}

type noConverterFoundError struct {
	unitName string
}

func (e *noConverterFoundError) Error() string {
	return fmt.Sprintf("No converter found for Unit %s", e.unitName)
}

var errEmptyUnit = errors.New("Unit can't be empty")

//Convert is the implementation of the convert method for TemperatureConverter.
func (tc *TemperatureConverter) Convert(t temperature.Temperature, u temperature.Unit) (*temperature.Temperature, error) {
	if t.Unit.Name == "" || u.Name == "" {
		return nil, errEmptyUnit
	}

	if t.Unit == u {
		return &t, nil
	}

	var tempInKelvin float64
	if t.Unit == temperature.Kelvin {
		tempInKelvin = t.Value
	} else {
		if tc.Converters[t.Unit] == nil {
			return nil, &noConverterFoundError{t.Unit.Name}
		}

		tempInKelvin = tc.Converters[t.Unit].ToKelvin(t.Value)
	}

	if u == temperature.Kelvin {
		return &temperature.Temperature{Value: tempInKelvin, Unit: u}, nil
	}

	if tc.Converters[u] == nil {
		return nil, &noConverterFoundError{t.Unit.Name}
	}

	return &temperature.Temperature{Value: tc.Converters[u].FromKelvin(tempInKelvin), Unit: u}, nil
}
