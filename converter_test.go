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
		output Convertible
		value  float64
		err    error
	}{
		{name: "CtoF 0", input: New(0, Celsius), output: Fahrenheit, value: 32},
		{name: "FtoC 32", input: New(32, Fahrenheit), output: Celsius, value: 0},
		{name: "KtoC 0", input: New(0, Kelvin), output: Celsius, value: -273.15},
		{name: "CtoK -273.15", input: New(-273.15, Celsius), output: Kelvin, value: 0},

		{name: "RatoRa 50", input: New(50, Rankine), output: Rankine, value: 50},
		{name: "RatoC 50", input: New(50, Rankine), output: Celsius, value: -245.37},
		{name: "RatoF 50", input: New(50, Rankine), output: Fahrenheit, value: -409.67},
		{name: "RatoK 50", input: New(50, Rankine), output: Kelvin, value: 27.78},
		{name: "RatoD 50", input: New(50, Rankine), output: Delisle, value: 518.06},
		{name: "RatoRe 50", input: New(50, Rankine), output: Reaumur, value: -196.3},

		{name: "FtoC 50", input: New(50, Fahrenheit), output: Celsius, value: 10},
		{name: "FtoF 50", input: New(50, Fahrenheit), output: Fahrenheit, value: 50},
		{name: "FtoK 50", input: New(50, Fahrenheit), output: Kelvin, value: 283.15},
		{name: "FtoRa 50", input: New(50, Fahrenheit), output: Rankine, value: 509.67},
		{name: "FtoD 50", input: New(50, Fahrenheit), output: Delisle, value: 135},
		{name: "FtoRe 50", input: New(50, Fahrenheit), output: Reaumur, value: 8},

		{name: "CtoC 50", input: New(50, Celsius), output: Celsius, value: 50},
		{name: "CtoK 50", input: New(50, Celsius), output: Kelvin, value: 323.15},
		{name: "CtoF 50", input: New(50, Celsius), output: Fahrenheit, value: 122},
		{name: "CtoRa 50", input: New(50, Celsius), output: Rankine, value: 581.67},
		{name: "CtoD 50", input: New(50, Celsius), output: Delisle, value: 75},
		{name: "CtoRe 50", input: New(50, Celsius), output: Reaumur, value: 40},

		{name: "KtoC 50", input: New(50, Kelvin), output: Celsius, value: -223.15},
		{name: "KtoF 50", input: New(50, Kelvin), output: Fahrenheit, value: -369.67},
		{name: "KtoK 50", input: New(50, Kelvin), output: Kelvin, value: 50},
		{name: "KtoRa 50", input: New(50, Kelvin), output: Rankine, value: 90},
		{name: "KtoD 50", input: New(50, Kelvin), output: Delisle, value: 484.73},
		{name: "KtoRe 50", input: New(50, Kelvin), output: Reaumur, value: -178.52},

		{name: "DtoD 50", input: New(50, Delisle), output: Delisle, value: 50},
		{name: "DtoC 50", input: New(50, Delisle), output: Celsius, value: 66.67},
		{name: "DtoF 50", input: New(50, Delisle), output: Fahrenheit, value: 152},
		{name: "DtoK 50", input: New(50, Delisle), output: Kelvin, value: 339.82},
		{name: "DtoRa 50", input: New(50, Delisle), output: Rankine, value: 611.67},
		{name: "DtoRe 50", input: New(50, Delisle), output: Reaumur, value: 53.33},

		{name: "RetoD 50", input: New(50, Reaumur), output: Delisle, value: 56.25},
		{name: "RetoC 50", input: New(50, Reaumur), output: Celsius, value: 62.5},
		{name: "RetoF 50", input: New(50, Reaumur), output: Fahrenheit, value: 144.5},
		{name: "RetoK 50", input: New(50, Reaumur), output: Kelvin, value: 335.65},
		{name: "RetoRa 50", input: New(50, Reaumur), output: Rankine, value: 604.17},
		{name: "RetoRe 50", input: New(50, Reaumur), output: Reaumur, value: 50},

		{name: "Convert Nil input", input: nil, output: Celsius, err: ErrNilArgument},
		{name: "Convert Nil outup", input: New(0, Celsius), output: nil, err: ErrNilArgument},
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
	c := New(30, Celsius)

	result, err := Convert(c, Fahrenheit)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)
	//Output: 86 °F
}

// func ExampleSetTemperature() {
// 	c := NewCelsius(30)
// 	f := NewFahrenheit(0)

// 	f.SetTemperature(c)
// 	fmt.Println(f)
// 	//Output: 86 °F
// }
