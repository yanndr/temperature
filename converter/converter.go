//Package converter provides converter to convert temperature units.
package converter

import (
	"errors"

	"github.com/yanndr/temperature"
)

//Converter is the interface implemented by temperature converter.
type Converter interface {
	//Convert  return a temperature with the new temperature unit.
	Convert(temperature.Temperature, temperature.Unit) (*temperature.Temperature, error)
}

//TemperatureConverter is the default implemtation of a temperature converter.
type TemperatureConverter struct {
	Converters map[temperature.Unit]kelvinConverter
}

//NewTemperatureConverter allocate a new TemperatureConverter and add the Celsuis and Fahrenheit converter to Kelvin.
func NewTemperatureConverter() *TemperatureConverter {
	tc := &TemperatureConverter{}
	tc.Converters = make(map[temperature.Unit]kelvinConverter)
	tc.Converters[temperature.Celsius] = &celsuisConverter{}
	tc.Converters[temperature.Fahrenheit] = &fahrenheitConverter{}
	return tc
}

//Convert is the implementation of the convert method for TemperatureConverter.
func (tc *TemperatureConverter) Convert(t temperature.Temperature, u temperature.Unit) (*temperature.Temperature, error) {
	if t.Unit == u {
		return &t, nil
	}

	var tempInKelvin *temperature.Temperature
	if t.Unit == temperature.Kelvin {
		tempInKelvin = &t
	} else {
		if tc.Converters[t.Unit] == nil {
			return nil, errors.New("no converter found")
		}

		tempInKelvin = tc.Converters[t.Unit].ToKelvin(t)
	}

	if u == temperature.Kelvin {
		return tempInKelvin, nil
	}

	if tc.Converters[u] == nil {
		return nil, errors.New("no converter found")
	}

	return tc.Converters[u].FromKelvin(*tempInKelvin), nil

}

type kelvinConverter interface {
	FromKelvin(temperature.Temperature) *temperature.Temperature
	ToKelvin(temperature.Temperature) *temperature.Temperature
}

type celsuisConverter struct {
}

func (*celsuisConverter) FromKelvin(t temperature.Temperature) *temperature.Temperature {
	return temperature.NewTemperature(t.Value-273.15, temperature.Celsius)
}

func (*celsuisConverter) ToKelvin(t temperature.Temperature) *temperature.Temperature {
	return temperature.NewTemperature(t.Value+273.15, temperature.Kelvin)
}

type fahrenheitConverter struct {
}

func (*fahrenheitConverter) FromKelvin(t temperature.Temperature) *temperature.Temperature {
	return temperature.NewTemperature(t.Value*9/5-459.67, temperature.Fahrenheit)
}

func (*fahrenheitConverter) ToKelvin(t temperature.Temperature) *temperature.Temperature {
	return temperature.NewTemperature((t.Value+459.67)*5/9, temperature.Kelvin)
}
