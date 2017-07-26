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
		{name: "Empty temp", temp: temperature.Temperature{}, unit: temperature.Fahrenheit, err: "Unit can't be empty"},
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
