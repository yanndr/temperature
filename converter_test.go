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
		text   string
		err    error
	}{
		{name: "CtoF 0", input: New(0, Celsius), output: Fahrenheit, value: 32, text: "32 °F"},
		{name: "FtoC 32", input: New(32, Fahrenheit), output: Celsius, value: 0, text: "0 °C"},
		{name: "KtoC 0", input: New(0, Kelvin), output: Celsius, value: -273.15, text: "-273.15 °C"},
		{name: "CtoK -273.15", input: New(-273.15, Celsius), output: Kelvin, value: 0, text: "0 K"},

		{name: "RatoRa 50", input: New(50, Rankine), output: Rankine, value: 50, text: "50 °Ra"},
		{name: "RatoC 50", input: New(50, Rankine), output: Celsius, value: -245.37, text: "-245.37 °C"},
		{name: "RatoF 50", input: New(50, Rankine), output: Fahrenheit, value: -409.67, text: "-409.67 °F"},
		{name: "RatoK 50", input: New(50, Rankine), output: Kelvin, value: 27.78, text: "27.78 K"},
		{name: "RatoD 50", input: New(50, Rankine), output: Delisle, value: 518.06, text: "518.06 °D"},
		{name: "RatoRe 50", input: New(50, Rankine), output: Reaumur, value: -196.3, text: "-196.3 °Re"},

		{name: "FtoC 50", input: New(50, Fahrenheit), output: Celsius, value: 10, text: "10 °C"},
		{name: "FtoF 50", input: New(50, Fahrenheit), output: Fahrenheit, value: 50, text: "50 °F"},
		{name: "FtoK 50", input: New(50, Fahrenheit), output: Kelvin, value: 283.15, text: "283.15 K"},
		{name: "FtoRa 50", input: New(50, Fahrenheit), output: Rankine, value: 509.67, text: "509.67 °Ra"},
		{name: "FtoD 50", input: New(50, Fahrenheit), output: Delisle, value: 135, text: "135 °D"},
		{name: "FtoRe 50", input: New(50, Fahrenheit), output: Reaumur, value: 8, text: "8 °Re"},

		{name: "CtoC 50", input: New(50, Celsius), output: Celsius, value: 50, text: "50 °C"},
		{name: "CtoK 50", input: New(50, Celsius), output: Kelvin, value: 323.15, text: "323.15 K"},
		{name: "CtoF 50", input: New(50, Celsius), output: Fahrenheit, value: 122, text: "122 °F"},
		{name: "CtoRa 50", input: New(50, Celsius), output: Rankine, value: 581.67, text: "581.67 °Ra"},
		{name: "CtoD 50", input: New(50, Celsius), output: Delisle, value: 75, text: "75 °D"},
		{name: "CtoRe 50", input: New(50, Celsius), output: Reaumur, value: 40, text: "40 °Re"},

		{name: "KtoC 50", input: New(50, Kelvin), output: Celsius, value: -223.15, text: "-223.15 °C"},
		{name: "KtoF 50", input: New(50, Kelvin), output: Fahrenheit, value: -369.67, text: "-369.67 °F"},
		{name: "KtoK 50", input: New(50, Kelvin), output: Kelvin, value: 50, text: "50 K"},
		{name: "KtoRa 50", input: New(50, Kelvin), output: Rankine, value: 90, text: "90 °Ra"},
		{name: "KtoD 50", input: New(50, Kelvin), output: Delisle, value: 484.73, text: "484.73 °D"},
		{name: "KtoRe 50", input: New(50, Kelvin), output: Reaumur, value: -178.52, text: "-178.52 °Re"},

		{name: "DtoD 50", input: New(50, Delisle), output: Delisle, value: 50, text: "50 °D"},
		{name: "DtoC 50", input: New(50, Delisle), output: Celsius, value: 66.67, text: "66.67 °C"},
		{name: "DtoF 50", input: New(50, Delisle), output: Fahrenheit, value: 152, text: "152 °F"},
		{name: "DtoK 50", input: New(50, Delisle), output: Kelvin, value: 339.82, text: "339.82 K"},
		{name: "DtoRa 50", input: New(50, Delisle), output: Rankine, value: 611.67, text: "611.67 °Ra"},
		{name: "DtoRe 50", input: New(50, Delisle), output: Reaumur, value: 53.33, text: "53.33 °Re"},

		{name: "RetoD 50", input: New(50, Reaumur), output: Delisle, value: 56.25, text: "56.25 °D"},
		{name: "RetoC 50", input: New(50, Reaumur), output: Celsius, value: 62.5, text: "62.5 °C"},
		{name: "RetoF 50", input: New(50, Reaumur), output: Fahrenheit, value: 144.5, text: "144.5 °F"},
		{name: "RetoK 50", input: New(50, Reaumur), output: Kelvin, value: 335.65, text: "335.65 K"},
		{name: "RetoRa 50", input: New(50, Reaumur), output: Rankine, value: 604.17, text: "604.17 °Ra"},
		{name: "RetoRe 50", input: New(50, Reaumur), output: Reaumur, value: 50, text: "50 °Re"},

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

			if result.String() != tc.text {
				t.Fatalf("Expected %v; got %v", tc.text, result.String())
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
