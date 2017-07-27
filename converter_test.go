package temperature_test

import (
	"fmt"
	"testing"

	"github.com/yanndr/temperature"
)

func TestConverter(t *testing.T) {

	tt := []struct {
		name   string
		input  temperature.ToKelvinConvertible
		output temperature.FromKelvinConvertible
		value  float64
		err    error
	}{
		{name: "CtoF 0", input: temperature.NewCelsius(0), output: &temperature.Fahrenheit{}, value: 32},
		{name: "FtoC 32", input: temperature.NewFahrenheit(32), output: &temperature.Celsius{}, value: 0},
		{name: "KtoC 32", input: temperature.NewKelvin(0), output: &temperature.Celsius{}, value: -273.15},
		{name: "CtoK -273.15", input: temperature.NewCelsius(-273.15), output: &temperature.Kelvin{}, value: 0},

		{name: "RtoR 50", input: temperature.NewRankine(50), output: &temperature.Rankine{}, value: 50},
		{name: "RtoC 50", input: temperature.NewRankine(50), output: &temperature.Celsius{}, value: -245.37},
		{name: "RtoF 50", input: temperature.NewRankine(50), output: &temperature.Fahrenheit{}, value: -409.67},
		{name: "RtoK 50", input: temperature.NewRankine(50), output: &temperature.Kelvin{}, value: 27.78},
		{name: "RtoD 50", input: temperature.NewRankine(50), output: &temperature.Delisle{}, value: 518.06},
		// {name: "RtoRe 50", temp: temperature.Temperature{Value: 50, Scale: temperature.Rankine}, scale: temperature.Reaumur, value: -196.3},

		{name: "FtoC 50", input: temperature.NewFahrenheit(50), output: &temperature.Celsius{}, value: 10},
		{name: "FtoF 50", input: temperature.NewFahrenheit(50), output: &temperature.Fahrenheit{}, value: 50},
		{name: "FtoK 50", input: temperature.NewFahrenheit(50), output: &temperature.Kelvin{}, value: 283.15},
		{name: "FtoRa 50", input: temperature.NewFahrenheit(50), output: &temperature.Rankine{}, value: 509.67},
		{name: "FtoD 50", input: temperature.NewFahrenheit(50), output: &temperature.Delisle{}, value: 135},
		// {name: "FtoRe 50", temp: temperature.Temperature{Value: 50, Scale: temperature.Fahrenheit}, scale: temperature.Reaumur, value: 8},

		{name: "CtoC 50", input: temperature.NewCelsius(50), output: &temperature.Celsius{}, value: 50},
		{name: "CtoK 50", input: temperature.NewCelsius(50), output: &temperature.Kelvin{}, value: 323.15},
		{name: "CtoF 50", input: temperature.NewCelsius(50), output: &temperature.Fahrenheit{}, value: 122},
		{name: "CtoRa 50", input: temperature.NewCelsius(50), output: &temperature.Rankine{}, value: 581.67},
		{name: "CtoD 50", input: temperature.NewCelsius(50), output: &temperature.Delisle{}, value: 75},
		// {name: "CtoRe 50", temp: temperature.Temperature{Value: 50, Scale: temperature.Celsius}, scale: temperature.Reaumur, value: 40},

		{name: "KtoC 50", input: temperature.NewKelvin(50), output: &temperature.Celsius{}, value: -223.15},
		{name: "KtoF 50", input: temperature.NewKelvin(50), output: &temperature.Fahrenheit{}, value: -369.67},
		{name: "KtoK 50", input: temperature.NewKelvin(50), output: &temperature.Kelvin{}, value: 50},
		{name: "KtoRa 50", input: temperature.NewKelvin(50), output: &temperature.Rankine{}, value: 90},
		{name: "KtoD 50", input: temperature.NewKelvin(50), output: &temperature.Delisle{}, value: 484.73},
		// {name: "KtoRe 50", temp: temperature.Temperature{Value: 50, Scale: temperature.Kelvin}, scale: temperature.Reaumur, value: -178.52},

		{name: "DtoD 50", input: temperature.NewDelisle(50), output: &temperature.Delisle{}, value: 50},
		{name: "DtoC 50", input: temperature.NewDelisle(50), output: &temperature.Celsius{}, value: 66.67},
		{name: "DtoF 50", input: temperature.NewDelisle(50), output: &temperature.Fahrenheit{}, value: 152},
		{name: "DtoK 50", input: temperature.NewDelisle(50), output: &temperature.Kelvin{}, value: 339.82},
		{name: "DtoRa 50", input: temperature.NewDelisle(50), output: &temperature.Rankine{}, value: 611.67},
		// {name: "DtoRe 50", temp: temperature.Temperature{Value: 50, Scale: temperature.Delisle}, scale: temperature.Reaumur, value: 53.33},

		{name: "RetoD 50", input: temperature.NewReaumur(50), output: &temperature.Delisle{}, value: 56.25},
		{name: "RetoC 50", input: temperature.NewReaumur(50), output: &temperature.Celsius{}, value: 62.5},
		{name: "RetoF 50", input: temperature.NewReaumur(50), output: &temperature.Fahrenheit{}, value: 144.5},
		{name: "RetoK 50", input: temperature.NewReaumur(50), output: &temperature.Kelvin{}, value: 335.65},
		{name: "RetoRa 50", input: temperature.NewReaumur(50), output: &temperature.Rankine{}, value: 604.17},
		{name: "RetoRe 50", input: temperature.NewReaumur(50), output: &temperature.Reaumur{}, value: 50},

		{name: "Convert Nil outupt", input: temperature.NewReaumur(50), output: nil, err: temperature.ErrNilOutputTemperature},
		{name: "Convert Nil input", input: nil, output: &temperature.Celsius{}, err: temperature.ErrNilInputTemperature},
		// {name: "Empty temp", scale: temperature.Fahrenheit, err: "Scale can't be empty"},
		// {name: "Empty scale", temp: temperature.Temperature{Value: 0.0, Scale: temperature.Celsius}, err: "Scale can't be empty"},
		// {name: "No input converter", temp: temperature.Temperature{Value: 0.0, Scale: temperature.Scale{Name: "Test", Symbol: "T"}}, scale: temperature.Celsius, err: "No converter found for Scale Test"},
		// {name: "No output converter", temp: temperature.Temperature{Value: 0.0, Scale: temperature.Celsius}, scale: temperature.Scale{Name: "Test", Symbol: "T"}, err: "No converter found for Scale Test"},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			if err := temperature.Convert(tc.input, tc.output); err != nil {
				if err != tc.err {
					t.Fatalf("Expected error %s; got %s.", tc.err, err.Error())
				}
				return
			}

			valuer, ok := tc.output.(temperature.Valuer)
			if !ok {
				return
			}
			val := temperature.Round(valuer.GetValue(), 2)
			if val != tc.value {

				t.Fatalf("Expected %v; got %v", tc.value, val)
			}
		})
	}
}

func ExampleConvert() {
	c := temperature.NewCelsius(30)
	f := temperature.Fahrenheit{}

	temperature.Convert(c, &f)

	fmt.Println(f)
	//Output: 86 °F
}
