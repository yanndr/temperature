package converter

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yanndr/temperature"
)

func TestCeslsuisToKelvinConverter(t *testing.T) {
	c := &celsiusKelvinConverter{}

	temp := temperature.Temperature{Value: 0.0, Unit: temperature.Celsius}
	result := c.ToKelvin(temp)
	excpecteResult := 273.15

	assert.Equal(t, excpecteResult, result.Value)

	temp = temperature.Temperature{Value: 35.0, Unit: temperature.Celsius}
	result = c.ToKelvin(temp)
	excpecteResult = 308.15

	assert.Equal(t, excpecteResult, result.Value)

	temp = temperature.Temperature{Value: 0, Unit: temperature.Kelvin}
	result = c.FromKelvin(temp)
	excpecteResult = -273.15

	assert.Equal(t, excpecteResult, result.Value)
}

func TestFahrenheitToKelvinConverter(t *testing.T) {
	c := &fahrenheitKelvinConverter{}

	temp := temperature.Temperature{Value: 0.0, Unit: temperature.Fahrenheit}
	result := c.ToKelvin(temp)
	excpecteResult := 255.37

	assert.Equal(t, excpecteResult, temperature.ToFixed(result.Value, 2))

	temp = temperature.Temperature{Value: 35.0, Unit: temperature.Fahrenheit}
	result = c.ToKelvin(temp)
	excpecteResult = 274.82

	assert.Equal(t, excpecteResult, temperature.ToFixed(result.Value, 2))

	temp = temperature.Temperature{Value: 0, Unit: temperature.Kelvin}
	result = c.FromKelvin(temp)
	excpecteResult = -459.67

	assert.Equal(t, excpecteResult, result.Value)
}
