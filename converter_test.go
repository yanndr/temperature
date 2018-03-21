package temperature

import (
	"fmt"
	"log"
	"testing"
)

func TestConverter(t *testing.T) {

	tt := []struct {
		name   string
		input  Temperature
		output Temperature
		value  float64
		err    error
	}{
		{name: "CtoF 0", input: NewCelsius(0), output: NewFahrenheit(0), value: 32},
		{name: "FtoC 32", input: NewFahrenheit(32), output: NewCelsius(0), value: 0},
		{name: "KtoC 32", input: NewKelvin(0), output: NewCelsius(0), value: -273.15},
		{name: "CtoK -273.15", input: NewCelsius(-273.15), output: NewKelvin(0), value: 0},

		{name: "RatoRa 50", input: NewRankine(50), output: NewRankine(0), value: 50},
		{name: "RatoC 50", input: NewRankine(50), output: NewCelsius(0), value: -245.37},
		{name: "RatoF 50", input: NewRankine(50), output: NewFahrenheit(0), value: -409.67},
		{name: "RatoK 50", input: NewRankine(50), output: NewKelvin(0), value: 27.78},
		{name: "RatoD 50", input: NewRankine(50), output: NewDelisle(0), value: 518.06},
		{name: "RatoRe 50", input: NewRankine(50), output: NewReaumur(0), value: -196.3},

		{name: "FtoC 50", input: NewFahrenheit(50), output: NewCelsius(0), value: 10},
		{name: "FtoF 50", input: NewFahrenheit(50), output: NewFahrenheit(0), value: 50},
		{name: "FtoK 50", input: NewFahrenheit(50), output: NewKelvin(0), value: 283.15},
		{name: "FtoRa 50", input: NewFahrenheit(50), output: NewRankine(0), value: 509.67},
		{name: "FtoD 50", input: NewFahrenheit(50), output: NewDelisle(0), value: 135},
		{name: "FtoRe 50", input: NewFahrenheit(50), output: NewReaumur(0), value: 8},

		{name: "CtoC 50", input: NewCelsius(50), output: NewCelsius(0), value: 50},
		{name: "CtoK 50", input: NewCelsius(50), output: NewKelvin(0), value: 323.15},
		{name: "CtoF 50", input: NewCelsius(50), output: NewFahrenheit(0), value: 122},
		{name: "CtoRa 50", input: NewCelsius(50), output: NewRankine(0), value: 581.67},
		{name: "CtoD 50", input: NewCelsius(50), output: NewDelisle(0), value: 75},
		{name: "CtoRe 50", input: NewCelsius(50), output: NewReaumur(0), value: 40},

		{name: "KtoC 50", input: NewKelvin(50), output: NewCelsius(0), value: -223.15},
		{name: "KtoF 50", input: NewKelvin(50), output: NewFahrenheit(0), value: -369.67},
		{name: "KtoK 50", input: NewKelvin(50), output: NewKelvin(0), value: 50},
		{name: "KtoRa 50", input: NewKelvin(50), output: NewRankine(0), value: 90},
		{name: "KtoD 50", input: NewKelvin(50), output: NewDelisle(0), value: 484.73},
		{name: "KtoRe 50", input: NewKelvin(50), output: NewReaumur(0), value: -178.52},

		{name: "DtoD 50", input: NewDelisle(50), output: NewDelisle(0), value: 50},
		{name: "DtoC 50", input: NewDelisle(50), output: NewCelsius(0), value: 66.67},
		{name: "DtoF 50", input: NewDelisle(50), output: NewFahrenheit(0), value: 152},
		{name: "DtoK 50", input: NewDelisle(50), output: NewKelvin(0), value: 339.82},
		{name: "DtoRa 50", input: NewDelisle(50), output: NewRankine(0), value: 611.67},
		{name: "DtoRe 50", input: NewDelisle(50), output: NewReaumur(0), value: 53.33},

		{name: "RetoD 50", input: NewReaumur(50), output: NewDelisle(0), value: 56.25},
		{name: "RetoC 50", input: NewReaumur(50), output: NewCelsius(0), value: 62.5},
		{name: "RetoF 50", input: NewReaumur(50), output: NewFahrenheit(0), value: 144.5},
		{name: "RetoK 50", input: NewReaumur(50), output: NewKelvin(0), value: 335.65},
		{name: "RetoRa 50", input: NewReaumur(50), output: NewRankine(0), value: 604.17},
		{name: "RetoRe 50", input: NewReaumur(50), output: NewReaumur(0), value: 50},

		{name: "Convert Nil outupt", input: NewReaumur(50), output: nil, err: ErrNilOutputTemperature},
		{name: "Convert Nil input", input: nil, output: NewCelsius(0), err: ErrNilInputTemperature},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			log.Println(tc.name)
			if err := Convert(tc.input, tc.output); err != nil {
				if err != tc.err {
					t.Fatalf("Expected error %s; got %s.", tc.err, err.Error())
				}
				return
			}

			val := round(tc.output.Value(), 2)
			if val != tc.value {

				t.Fatalf("Expected %v; got %v", tc.value, val)
			}
		})
	}
}

func ExampleConvert() {
	c := NewCelsius(30)
	f := NewFahrenheit(0)

	Convert(c, f)

	fmt.Println(f)
	//Output: 86 °F
}

func ExampleSetTemperature() {
	c := NewCelsius(30)
	f := NewFahrenheit(0)

	f.SetTemperature(c)
	fmt.Println(f)
	//Output: 86 °F
}
