package temperature_test

import (
	"testing"

	"github.com/yanndr/temperature"
)

type stringer interface {
	String() string
}

func TestString(t *testing.T) {
	tt := []struct {
		name        string
		temperature stringer
		result      string
	}{
		{name: "Print Celsius", temperature: temperature.NewCelsius(10), result: "10 °C"},
		{name: "Print Fahrenheit", temperature: temperature.NewFahrenheit(10), result: "10 °F"},
		{name: "Print Delisle", temperature: temperature.NewDelisle(10), result: "10 °D"},
		{name: "Print Kelvin", temperature: temperature.NewKelvin(10), result: "10 K"},
		{name: "Print Reaumur", temperature: temperature.NewReaumur(10), result: "10 °Re"},
		{name: "Print Rankine", temperature: temperature.NewRankine(10), result: "10 °Ra"},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			res := tc.temperature.String()
			if res != tc.result {
				t.Fatalf("Expected %v; got %v", tc.result, res)
			}

		})
	}
}
