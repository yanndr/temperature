package converter

import (
	"testing"

	"github.com/yanndr/temperature"
)

func TestTemperatureConvert(t *testing.T) {

	tt := []struct {
		name  string
		temp  temperature.Temperature
		unit  temperature.Unit
		value float64
		err   string
	}{
		{name: "CtoF zero", temp: temperature.Temperature{Value: 0.0, Unit: temperature.Celsius}, unit: temperature.Fahrenheit, value: 32.0},
		{name: "FtoC 32", temp: temperature.Temperature{Value: 0.0, Unit: temperature.Celsius}, unit: temperature.Fahrenheit, value: 32.0},
		{name: "KtoK 0", temp: temperature.Temperature{Value: 0.0, Unit: temperature.Kelvin}, unit: temperature.Kelvin, value: 0.0},
		{name: "KtoC 0", temp: temperature.Temperature{Value: 0.0, Unit: temperature.Kelvin}, unit: temperature.Celsius, value: -273.15},
		{name: "CtoC 0", temp: temperature.Temperature{Value: 0.0, Unit: temperature.Celsius}, unit: temperature.Celsius, value: 0.0},
		{name: "Ctok 0", temp: temperature.Temperature{Value: -273.15, Unit: temperature.Celsius}, unit: temperature.Kelvin, value: 0.0},
		{name: "RtoR 50", temp: temperature.Temperature{Value: 50, Unit: temperature.Rankine}, unit: temperature.Rankine, value: 50},
		{name: "RtoC 50", temp: temperature.Temperature{Value: 50, Unit: temperature.Rankine}, unit: temperature.Celsius, value: -245.37},
		{name: "RtoF 50", temp: temperature.Temperature{Value: 50, Unit: temperature.Rankine}, unit: temperature.Fahrenheit, value: -409.67},
		{name: "RtoK 50", temp: temperature.Temperature{Value: 50, Unit: temperature.Rankine}, unit: temperature.Kelvin, value: 27.78},
		{name: "FtoC 50", temp: temperature.Temperature{Value: 50, Unit: temperature.Fahrenheit}, unit: temperature.Celsius, value: 10},
		{name: "FtoF 50", temp: temperature.Temperature{Value: 50, Unit: temperature.Fahrenheit}, unit: temperature.Fahrenheit, value: 50},
		{name: "FtoK 50", temp: temperature.Temperature{Value: 50, Unit: temperature.Fahrenheit}, unit: temperature.Kelvin, value: 283.15},
		{name: "FtoR 50", temp: temperature.Temperature{Value: 50, Unit: temperature.Fahrenheit}, unit: temperature.Rankine, value: 509.67},
		{name: "CtoC 50", temp: temperature.Temperature{Value: 50, Unit: temperature.Celsius}, unit: temperature.Celsius, value: 50},
		{name: "CtoF 50", temp: temperature.Temperature{Value: 50, Unit: temperature.Celsius}, unit: temperature.Fahrenheit, value: 122},
		{name: "CtoK 50", temp: temperature.Temperature{Value: 50, Unit: temperature.Celsius}, unit: temperature.Kelvin, value: 323.15},
		{name: "CtoR 50", temp: temperature.Temperature{Value: 50, Unit: temperature.Celsius}, unit: temperature.Rankine, value: 581.67},
		{name: "FtoC 50", temp: temperature.Temperature{Value: 50, Unit: temperature.Fahrenheit}, unit: temperature.Celsius, value: 10},
		{name: "FtoF 50", temp: temperature.Temperature{Value: 50, Unit: temperature.Fahrenheit}, unit: temperature.Fahrenheit, value: 50},
		{name: "FtoK 50", temp: temperature.Temperature{Value: 50, Unit: temperature.Fahrenheit}, unit: temperature.Kelvin, value: 283.15},
		{name: "FtoR 50", temp: temperature.Temperature{Value: 50, Unit: temperature.Fahrenheit}, unit: temperature.Rankine, value: 509.67},
		{name: "Empty temp", unit: temperature.Fahrenheit, err: "Unit can't be empty"},
		{name: "Empty unit", temp: temperature.Temperature{Value: 0.0, Unit: temperature.Celsius}, err: "Unit can't be empty"},
		{name: "No input converter", temp: temperature.Temperature{Value: 0.0, Unit: temperature.Unit{Name: "Test", Symbol: "T"}}, unit: temperature.Celsius, err: "No converter found for Unit Test"},
		{name: "No output converter", temp: temperature.Temperature{Value: 0.0, Unit: temperature.Celsius}, unit: temperature.Unit{Name: "Test", Symbol: "T"}, err: "No converter found for Unit Test"},
	}

	c := NewTemperatureConverter()

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			res, err := c.Convert(tc.temp, tc.unit)

			if err != nil {
				if err.Error() != tc.err {
					t.Fatalf("Expected error %s; got %s.", tc.err, err.Error())
				}
				return
			}

			val := temperature.Round(res.Value, 2)
			if val != tc.value {

				t.Fatalf("Expected %v; got %v", tc.value, val)
			}

			if res.Unit != tc.unit {
				t.Fatalf("Expected %v; got %v", tc.unit, res.Unit)
			}
		})
	}

}
