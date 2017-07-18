//Package converter provides converter to convert temperature units.
package converter

import (
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
	Converters map[temperature.Unit]kelvinConverter
}

//NewTemperatureConverter allocate a new TemperatureConverter.
func NewTemperatureConverter() *TemperatureConverter {
	tc := &TemperatureConverter{}
	tc.Converters = make(map[temperature.Unit]kelvinConverter)
	tc.Converters[temperature.Celsius] = &celsiusKelvinConverter{}
	tc.Converters[temperature.Fahrenheit] = &fahrenheitKelvinConverter{}
	tc.Converters[temperature.Rankine] = &rankineKelvinConverter{}
	tc.Converters[temperature.Delisle] = &delisleKelvinConverter{}
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
			return nil, fmt.Errorf("no converter found for unit %s", t.Unit.Name)
		}

		tempInKelvin = tc.Converters[t.Unit].ToKelvin(t)
	}

	if u == temperature.Kelvin {
		return tempInKelvin, nil
	}

	if tc.Converters[u] == nil {
		return nil, fmt.Errorf("no converter found for unit %s", u.Name)
	}

	return tc.Converters[u].FromKelvin(*tempInKelvin), nil

}

type kelvinConverter interface {
	FromKelvin(temperature.Temperature) *temperature.Temperature
	ToKelvin(temperature.Temperature) *temperature.Temperature
}

type celsiusKelvinConverter struct {
}

func (*celsiusKelvinConverter) FromKelvin(t temperature.Temperature) *temperature.Temperature {
	return temperature.NewTemperature(t.Value-273.15, temperature.Celsius)
}

func (*celsiusKelvinConverter) ToKelvin(t temperature.Temperature) *temperature.Temperature {
	return temperature.NewTemperature(t.Value+273.15, temperature.Kelvin)
}

type fahrenheitKelvinConverter struct {
}

func (*fahrenheitKelvinConverter) FromKelvin(t temperature.Temperature) *temperature.Temperature {
	return temperature.NewTemperature(t.Value*9/5-459.67, temperature.Fahrenheit)
}

func (*fahrenheitKelvinConverter) ToKelvin(t temperature.Temperature) *temperature.Temperature {
	return temperature.NewTemperature((t.Value+459.67)*5/9, temperature.Kelvin)
}

type rankineKelvinConverter struct {
}

func (*rankineKelvinConverter) FromKelvin(t temperature.Temperature) *temperature.Temperature {
	return temperature.NewTemperature(t.Value*9/5, temperature.Rankine)
}

func (*rankineKelvinConverter) ToKelvin(t temperature.Temperature) *temperature.Temperature {
	return temperature.NewTemperature((t.Value)*5/9, temperature.Kelvin)
}

type delisleKelvinConverter struct {
}

func (*delisleKelvinConverter) FromKelvin(t temperature.Temperature) *temperature.Temperature {
	return temperature.NewTemperature((373.15-t.Value)*3/2, temperature.Delisle)
}

func (*delisleKelvinConverter) ToKelvin(t temperature.Temperature) *temperature.Temperature {
	return temperature.NewTemperature(373.15-t.Value*2/3, temperature.Kelvin)
}
