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
	assert.Equal(t, temperature.Kelvin, result.Unit)

	temp = temperature.Temperature{Value: 35.0, Unit: temperature.Celsius}
	result = c.ToKelvin(temp)
	excpecteResult = 308.15

	assert.Equal(t, excpecteResult, result.Value)
	assert.Equal(t, temperature.Kelvin, result.Unit)

	temp = temperature.Temperature{Value: 0, Unit: temperature.Kelvin}
	result = c.FromKelvin(temp)
	excpecteResult = -273.15

	assert.Equal(t, excpecteResult, result.Value)
	assert.Equal(t, temperature.Celsius, result.Unit)
}

func TestFahrenheitToKelvinConverter(t *testing.T) {
	c := &fahrenheitKelvinConverter{}

	temp := temperature.Temperature{Value: 0.0, Unit: temperature.Fahrenheit}
	result := c.ToKelvin(temp)
	excpecteResult := 255.37

	assert.Equal(t, excpecteResult, temperature.ToFixed(result.Value, 2))
	assert.Equal(t, temperature.Kelvin, result.Unit)

	temp = temperature.Temperature{Value: 35.0, Unit: temperature.Fahrenheit}
	result = c.ToKelvin(temp)
	excpecteResult = 274.82

	assert.Equal(t, excpecteResult, temperature.ToFixed(result.Value, 2))
	assert.Equal(t, temperature.Kelvin, result.Unit)

	temp = temperature.Temperature{Value: 0, Unit: temperature.Kelvin}
	result = c.FromKelvin(temp)
	excpecteResult = -459.67

	assert.Equal(t, excpecteResult, result.Value)
	assert.Equal(t, temperature.Fahrenheit, result.Unit)
}

func TestRankineToKelvinConverter(t *testing.T) {
	c := &rankineKelvinConverter{}

	temp := temperature.Temperature{Value: 0.0, Unit: temperature.Rankine}
	result := c.ToKelvin(temp)
	excpecteResult := 0.0

	assert.Equal(t, excpecteResult, temperature.ToFixed(result.Value, 2))
	assert.Equal(t, temperature.Kelvin, result.Unit)

	temp = temperature.Temperature{Value: 35.0, Unit: temperature.Rankine}
	result = c.ToKelvin(temp)
	excpecteResult = 19.44

	assert.Equal(t, excpecteResult, temperature.ToFixed(result.Value, 2))
	assert.Equal(t, temperature.Kelvin, result.Unit)

	temp = temperature.Temperature{Value: 0, Unit: temperature.Kelvin}
	result = c.FromKelvin(temp)
	excpecteResult = 0

	assert.Equal(t, excpecteResult, result.Value)
	assert.Equal(t, temperature.Rankine, result.Unit)

	temp = temperature.Temperature{Value: 55, Unit: temperature.Kelvin}
	result = c.FromKelvin(temp)
	excpecteResult = 99

	assert.Equal(t, excpecteResult, result.Value)
	assert.Equal(t, temperature.Rankine, result.Unit)
}

func TestDelisleToKelvinConverter(t *testing.T) {
	c := &delisleKelvinConverter{}

	temp := temperature.Temperature{Value: 0.0, Unit: temperature.Delisle}
	result := c.ToKelvin(temp)
	excpecteResult := 373.15

	assert.Equal(t, excpecteResult, temperature.ToFixed(result.Value, 2))
	assert.Equal(t, temperature.Kelvin, result.Unit)

	temp = temperature.Temperature{Value: 35.0, Unit: temperature.Delisle}
	result = c.ToKelvin(temp)
	excpecteResult = 349.82

	assert.Equal(t, excpecteResult, temperature.ToFixed(result.Value, 2))
	assert.Equal(t, temperature.Kelvin, result.Unit)

	temp = temperature.Temperature{Value: 0, Unit: temperature.Kelvin}
	result = c.FromKelvin(temp)
	excpecteResult = 559.72

	assert.Equal(t, excpecteResult, temperature.ToFixed(result.Value, 2))
	assert.Equal(t, temperature.Delisle, result.Unit)

	temp = temperature.Temperature{Value: 55, Unit: temperature.Kelvin}
	result = c.FromKelvin(temp)
	excpecteResult = 477.23

	assert.Equal(t, excpecteResult, temperature.ToFixed(result.Value, 2))
	assert.Equal(t, temperature.Delisle, result.Unit)
}
