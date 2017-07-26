package converter

import (
	"testing"

	"github.com/yanndr/temperature"
)

type tempTestCase struct {
	name  string
	temp  temperature.Temperature
	value float64
}

func TestCelsuisToKelvinConverter(t *testing.T) {

	tt := []tempTestCase{
		{name: "zero celsius to kelvin", temp: temperature.Temperature{Value: 0.0, Unit: temperature.Celsius}, value: 273.15},
		{name: "thirty five celsius to kelvin", temp: temperature.Temperature{Value: 35.0, Unit: temperature.Celsius}, value: 308.15},
	}
	c := &celsiusKelvinConverter{}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			res := c.ToKelvin(tc.temp)
			if res.Value != tc.value {
				t.Fatalf("Expected %v; got %v", tc.value, res.Value)
			}
			if res.Unit != temperature.Kelvin {
				t.Fatalf("Expected Kelvin; got %v", res.Unit)
			}
		})
	}
}

func TestKelvinToCelsiusConverter(t *testing.T) {

	tt := []tempTestCase{
		{name: "zero kelvin to celsius", temp: temperature.Temperature{Value: 0.0, Unit: temperature.Kelvin}, value: -273.15},
	}
	c := &celsiusKelvinConverter{}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			res := c.FromKelvin(tc.temp)
			if res.Value != tc.value {
				t.Fatalf("Expected %v; got %v", tc.value, res.Value)
			}
			if res.Unit != temperature.Celsius {
				t.Fatalf("Expected Celsius; got %v", res.Unit)
			}
		})
	}
}

func TestFahrenheitToKelvinConverter(t *testing.T) {

	tt := []tempTestCase{
		{name: "zero fahrenheit to kelvin", temp: temperature.Temperature{Value: 0.0, Unit: temperature.Fahrenheit}, value: 255.37},
		{name: "thirty five fahrenheit to kelvin", temp: temperature.Temperature{Value: 35.0, Unit: temperature.Fahrenheit}, value: 274.82},
	}
	c := &fahrenheitKelvinConverter{}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			res := c.ToKelvin(tc.temp)
			val := temperature.ToFixed(res.Value, 2)
			if val != tc.value {
				t.Fatalf("Expected %v; got %v", tc.value, val)
			}
			if res.Unit != temperature.Kelvin {
				t.Fatalf("Expected Kelvin; got %v", res.Unit)
			}
		})
	}
}

func TestKelvinToFahrenheitConverter(t *testing.T) {

	tt := []tempTestCase{
		{name: "zero kelvin to fahrenheit", temp: temperature.Temperature{Value: 0.0, Unit: temperature.Fahrenheit}, value: -459.67},
		{name: "thirty five kelvin to fahrenheit", temp: temperature.Temperature{Value: 35.0, Unit: temperature.Fahrenheit}, value: -396.67},
	}
	c := &fahrenheitKelvinConverter{}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			res := c.FromKelvin(tc.temp)
			val := temperature.ToFixed(res.Value, 2)
			if val != tc.value {
				t.Fatalf("Expected %v; got %v", tc.value, val)
			}
			if res.Unit != temperature.Fahrenheit {
				t.Fatalf("Expected Fahrenheit; got %v", res.Unit)
			}
		})
	}
}

func TestRankineToKelvinConverter(t *testing.T) {

	tt := []tempTestCase{
		{name: "zero Rankine to Kelvin", temp: temperature.Temperature{Value: 0.0, Unit: temperature.Fahrenheit}, value: 0.0},
		{name: "thirty five Rankine to Kelvin", temp: temperature.Temperature{Value: 35.0, Unit: temperature.Fahrenheit}, value: 19.44},
	}
	c := rankineKelvinConverter{}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			res := c.ToKelvin(tc.temp)
			val := temperature.ToFixed(res.Value, 2)
			if val != tc.value {
				t.Fatalf("Expected %v; got %v", tc.value, val)
			}
			if res.Unit != temperature.Kelvin {
				t.Fatalf("Expected Kelvin; got %v", res.Unit)
			}
		})
	}
}

func TestKelvinToRankineConverter(t *testing.T) {

	tt := []tempTestCase{
		{name: "zero Kelvin to Rankine", temp: temperature.Temperature{Value: 0.0, Unit: temperature.Kelvin}, value: 0.0},
		{name: "fifty five Kelvin to Rankine", temp: temperature.Temperature{Value: 55.0, Unit: temperature.Kelvin}, value: 99},
	}
	c := rankineKelvinConverter{}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			res := c.FromKelvin(tc.temp)
			val := temperature.ToFixed(res.Value, 2)
			if val != tc.value {
				t.Fatalf("Expected %v; got %v", tc.value, val)
			}
			if res.Unit != temperature.Rankine {
				t.Fatalf("Expected Rankine; got %v", res.Unit)
			}
		})
	}
}

func TestDelisleToKelvinConverter(t *testing.T) {

	tt := []tempTestCase{
		{name: "zero Delisle to Kelvin", temp: temperature.Temperature{Value: 0.0, Unit: temperature.Delisle}, value: 373.15},
		{name: "thirty five Delisle to Kelvin", temp: temperature.Temperature{Value: 35.0, Unit: temperature.Delisle}, value: 349.82},
	}
	c := delisleKelvinConverter{}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			res := c.ToKelvin(tc.temp)
			val := temperature.ToFixed(res.Value, 2)
			if val != tc.value {
				t.Fatalf("Expected %v; got %v", tc.value, val)
			}
			if res.Unit != temperature.Kelvin {
				t.Fatalf("Expected Kelvin; got %v", res.Unit)
			}
		})
	}
}

func TestKelvinToDelisleConverter(t *testing.T) {

	tt := []tempTestCase{
		{name: "zero Kelvin to Delisle", temp: temperature.Temperature{Value: 0.0, Unit: temperature.Kelvin}, value: 559.72},
		{name: "fifty five Kelvin to Delisle", temp: temperature.Temperature{Value: 55.0, Unit: temperature.Kelvin}, value: 477.23},
	}
	c := delisleKelvinConverter{}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			res := c.FromKelvin(tc.temp)
			val := temperature.ToFixed(res.Value, 2)
			if val != tc.value {
				t.Fatalf("Expected %v; got %v", tc.value, val)
			}
			if res.Unit != temperature.Delisle {
				t.Fatalf("Expected Delisle; got %v", res.Unit)
			}
		})
	}
}
