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
		output Unit
		value  float64
		err    error
	}{
		{name: "CtoF 0", input: NewCelsius(0), output: Fahrenheit, value: 32},
		{name: "FtoC 32", input: NewFahrenheit(32), output: Celsius, value: 0},
		{name: "KtoC 32", input: NewKelvin(0), output: Celsius, value: -273.15},
		{name: "CtoK -273.15", input: NewCelsius(-273.15), output: Kelvin, value: 0},

		{name: "RatoRa 50", input: NewRankine(50), output: Rankine, value: 50},
		{name: "RatoC 50", input: NewRankine(50), output: Celsius, value: -245.37},
		{name: "RatoF 50", input: NewRankine(50), output: Fahrenheit, value: -409.67},
		{name: "RatoK 50", input: NewRankine(50), output: Kelvin, value: 27.78},
		{name: "RatoD 50", input: NewRankine(50), output: Delisle, value: 518.06},
		{name: "RatoRe 50", input: NewRankine(50), output: Reaumur, value: -196.3},

		{name: "FtoC 50", input: NewFahrenheit(50), output: Celsius, value: 10},
		{name: "FtoF 50", input: NewFahrenheit(50), output: Fahrenheit, value: 50},
		{name: "FtoK 50", input: NewFahrenheit(50), output: Kelvin, value: 283.15},
		{name: "FtoRa 50", input: NewFahrenheit(50), output: Rankine, value: 509.67},
		{name: "FtoD 50", input: NewFahrenheit(50), output: Delisle, value: 135},
		{name: "FtoRe 50", input: NewFahrenheit(50), output: Reaumur, value: 8},

		{name: "CtoC 50", input: NewCelsius(50), output: Celsius, value: 50},
		{name: "CtoK 50", input: NewCelsius(50), output: Kelvin, value: 323.15},
		{name: "CtoF 50", input: NewCelsius(50), output: Fahrenheit, value: 122},
		{name: "CtoRa 50", input: NewCelsius(50), output: Rankine, value: 581.67},
		{name: "CtoD 50", input: NewCelsius(50), output: Delisle, value: 75},
		{name: "CtoRe 50", input: NewCelsius(50), output: Reaumur, value: 40},

		{name: "KtoC 50", input: NewKelvin(50), output: Celsius, value: -223.15},
		{name: "KtoF 50", input: NewKelvin(50), output: Fahrenheit, value: -369.67},
		{name: "KtoK 50", input: NewKelvin(50), output: Kelvin, value: 50},
		{name: "KtoRa 50", input: NewKelvin(50), output: Rankine, value: 90},
		{name: "KtoD 50", input: NewKelvin(50), output: Delisle, value: 484.73},
		{name: "KtoRe 50", input: NewKelvin(50), output: Reaumur, value: -178.52},

		{name: "DtoD 50", input: NewDelisle(50), output: Delisle, value: 50},
		{name: "DtoC 50", input: NewDelisle(50), output: Celsius, value: 66.67},
		{name: "DtoF 50", input: NewDelisle(50), output: Fahrenheit, value: 152},
		{name: "DtoK 50", input: NewDelisle(50), output: Kelvin, value: 339.82},
		{name: "DtoRa 50", input: NewDelisle(50), output: Rankine, value: 611.67},
		{name: "DtoRe 50", input: NewDelisle(50), output: Reaumur, value: 53.33},

		{name: "RetoD 50", input: NewReaumur(50), output: Delisle, value: 56.25},
		{name: "RetoC 50", input: NewReaumur(50), output: Celsius, value: 62.5},
		{name: "RetoF 50", input: NewReaumur(50), output: Fahrenheit, value: 144.5},
		{name: "RetoK 50", input: NewReaumur(50), output: Kelvin, value: 335.65},
		{name: "RetoRa 50", input: NewReaumur(50), output: Rankine, value: 604.17},
		{name: "RetoRe 50", input: NewReaumur(50), output: Reaumur, value: 50},

		{name: "Convert Nil input", input: nil, output: Celsius, err: ErrNilInputTemperature},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			log.Println(tc.name)
			var result Temperature
			var err error
			if result, err = Convert(tc.input, tc.output); err != nil {
				if err != tc.err {
					t.Fatalf("Expected error %s; got %s.", tc.err, err.Error())
				}
				return
			}

			val := round(result.Value(), 2)
			if val != tc.value {

				t.Fatalf("Expected %v; got %v", tc.value, val)
			}
		})
	}
}

func ExampleConvert() {
	c := NewCelsius(30)

	result, err := Convert(c, Fahrenheit)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)
	//Output: 86 °F
}

func ExampleSetTemperature() {
	c := NewCelsius(30)
	f := NewFahrenheit(0)

	f.SetTemperature(c)
	fmt.Println(f)
	//Output: 86 °F
}
