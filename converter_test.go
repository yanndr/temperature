package temperature

import (
	"testing"
)

func TestConverter(t *testing.T) {

	tt := []struct {
		name   string
		input  TemperatureConvertible
		output TemperatureConvertible
		value  float64
		err    string
	}{
		// {name: "CtoF zero", temp: temperature.Temperature{Value: 0.0, Scale: temperature.Celsius}, scale: temperature.Fahrenheit, value: 32.0},
		// {name: "FtoC 32", temp: temperature.Temperature{Value: 0.0, Scale: temperature.Celsius}, scale: temperature.Fahrenheit, value: 32.0},
		// {name: "KtoK 0", temp: temperature.Temperature{Value: 0.0, Scale: temperature.Kelvin}, scale: temperature.Kelvin, value: 0.0},
		// {name: "KtoC 0", temp: temperature.Temperature{Value: 0.0, Scale: temperature.Kelvin}, scale: temperature.Celsius, value: -273.15},
		// {name: "CtoC 0", temp: temperature.Temperature{Value: 0.0, Scale: temperature.Celsius}, scale: temperature.Celsius, value: 0.0},
		// {name: "Ctok 0", temp: temperature.Temperature{Value: -273.15, Scale: temperature.Celsius}, scale: temperature.Kelvin, value: 0.0},

		// {name: "RtoR 50", temp: temperature.Temperature{Value: 50, Scale: temperature.Rankine}, scale: temperature.Rankine, value: 50},
		// {name: "RtoC 50", temp: temperature.Temperature{Value: 50, Scale: temperature.Rankine}, scale: temperature.Celsius, value: -245.37},
		// {name: "RtoF 50", temp: temperature.Temperature{Value: 50, Scale: temperature.Rankine}, scale: temperature.Fahrenheit, value: -409.67},
		// {name: "RtoK 50", temp: temperature.Temperature{Value: 50, Scale: temperature.Rankine}, scale: temperature.Kelvin, value: 27.78},
		// {name: "RtoD 50", temp: temperature.Temperature{Value: 50, Scale: temperature.Rankine}, scale: temperature.Delisle, value: 518.06},
		// {name: "RtoRe 50", temp: temperature.Temperature{Value: 50, Scale: temperature.Rankine}, scale: temperature.Reaumur, value: -196.3},

		// {name: "FtoC 50", temp: temperature.Temperature{Value: 50, Scale: temperature.Fahrenheit}, scale: temperature.Celsius, value: 10},
		// {name: "FtoF 50", temp: temperature.Temperature{Value: 50, Scale: temperature.Fahrenheit}, scale: temperature.Fahrenheit, value: 50},
		// {name: "FtoK 50", temp: temperature.Temperature{Value: 50, Scale: temperature.Fahrenheit}, scale: temperature.Kelvin, value: 283.15},
		// {name: "FtoR 50", temp: temperature.Temperature{Value: 50, Scale: temperature.Fahrenheit}, scale: temperature.Rankine, value: 509.67},
		// {name: "FtoD 50", temp: temperature.Temperature{Value: 50, Scale: temperature.Fahrenheit}, scale: temperature.Delisle, value: 135},
		// {name: "FtoRe 50", temp: temperature.Temperature{Value: 50, Scale: temperature.Fahrenheit}, scale: temperature.Reaumur, value: 8},

		{name: "CtoC 50", input: celsius{Temperature: Temperature{value: 50}}, output: celsius{}, value: 50},
		{name: "CtoK 50", input: celsius{Temperature: Temperature{value: 50}}, output: kelvin{}, value: 323.15},
		// {name: "CtoF 50", temp: temperature.Temperature{Value: 50, Scale: temperature.Celsius}, scale: temperature.Fahrenheit, value: 122},
		// {name: "CtoK 50", temp: temperature.Temperature{Value: 50, Scale: temperature.Celsius}, scale: temperature.Kelvin, value: 323.15},
		// {name: "CtoR 50", temp: temperature.Temperature{Value: 50, Scale: temperature.Celsius}, scale: temperature.Rankine, value: 581.67},
		// {name: "CtoD 50", temp: temperature.Temperature{Value: 50, Scale: temperature.Celsius}, scale: temperature.Delisle, value: 75},
		// {name: "CtoRe 50", temp: temperature.Temperature{Value: 50, Scale: temperature.Celsius}, scale: temperature.Reaumur, value: 40},

		// {name: "KtoC 50", temp: temperature.Temperature{Value: 50, Scale: temperature.Kelvin}, scale: temperature.Celsius, value: -223.15},
		// {name: "KtoF 50", temp: temperature.Temperature{Value: 50, Scale: temperature.Kelvin}, scale: temperature.Fahrenheit, value: -369.67},
		// {name: "KtoK 50", temp: temperature.Temperature{Value: 50, Scale: temperature.Kelvin}, scale: temperature.Kelvin, value: 50},
		// {name: "KtoR 50", temp: temperature.Temperature{Value: 50, Scale: temperature.Kelvin}, scale: temperature.Rankine, value: 90},
		// {name: "KtoD 50", temp: temperature.Temperature{Value: 50, Scale: temperature.Kelvin}, scale: temperature.Delisle, value: 484.73},
		// {name: "KtoRe 50", temp: temperature.Temperature{Value: 50, Scale: temperature.Kelvin}, scale: temperature.Reaumur, value: -178.52},

		// {name: "DtoD 50", temp: temperature.Temperature{Value: 50, Scale: temperature.Delisle}, scale: temperature.Delisle, value: 50},
		// {name: "DtoC 50", temp: temperature.Temperature{Value: 50, Scale: temperature.Delisle}, scale: temperature.Celsius, value: 66.67},
		// {name: "DtoF 50", temp: temperature.Temperature{Value: 50, Scale: temperature.Delisle}, scale: temperature.Fahrenheit, value: 152},
		// {name: "DtoK 50", temp: temperature.Temperature{Value: 50, Scale: temperature.Delisle}, scale: temperature.Kelvin, value: 339.82},
		// {name: "DtoR 50", temp: temperature.Temperature{Value: 50, Scale: temperature.Delisle}, scale: temperature.Rankine, value: 611.67},
		// {name: "DtoRe 50", temp: temperature.Temperature{Value: 50, Scale: temperature.Delisle}, scale: temperature.Reaumur, value: 53.33},

		// {name: "RetoD 50", temp: temperature.Temperature{Value: 50, Scale: temperature.Reaumur}, scale: temperature.Delisle, value: 56.25},
		// {name: "RetoC 50", temp: temperature.Temperature{Value: 50, Scale: temperature.Reaumur}, scale: temperature.Celsius, value: 62.5},
		// {name: "RetoF 50", temp: temperature.Temperature{Value: 50, Scale: temperature.Reaumur}, scale: temperature.Fahrenheit, value: 144.5},
		// {name: "RetoK 50", temp: temperature.Temperature{Value: 50, Scale: temperature.Reaumur}, scale: temperature.Kelvin, value: 335.65},
		// {name: "RetoR 50", temp: temperature.Temperature{Value: 50, Scale: temperature.Reaumur}, scale: temperature.Rankine, value: 604.17},
		// {name: "RetoRr 50", temp: temperature.Temperature{Value: 50, Scale: temperature.Reaumur}, scale: temperature.Reaumur, value: 50},

		// {name: "Empty temp", scale: temperature.Fahrenheit, err: "Scale can't be empty"},
		// {name: "Empty scale", temp: temperature.Temperature{Value: 0.0, Scale: temperature.Celsius}, err: "Scale can't be empty"},
		// {name: "No input converter", temp: temperature.Temperature{Value: 0.0, Scale: temperature.Scale{Name: "Test", Symbol: "T"}}, scale: temperature.Celsius, err: "No converter found for Scale Test"},
		// {name: "No output converter", temp: temperature.Temperature{Value: 0.0, Scale: temperature.Celsius}, scale: temperature.Scale{Name: "Test", Symbol: "T"}, err: "No converter found for Scale Test"},
	}

	c := Converter{}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			res := c.Convert(tc.input, tc.output)

			// if err != nil {
			// 	if err.Error() != tc.err {
			// 		t.Fatalf("Expected error %s; got %s.", tc.err, err.Error())
			// 	}
			// 	return
			// }

			val := Round(res.Value(), 2)
			if val != tc.value {

				t.Fatalf("Expected %v; got %v", tc.value, val)
			}

			// if res.Scale != tc.scale {
			// 	t.Fatalf("Expected %v; got %v", tc.scale, res.Scale)
			// }
		})
	}
}

// func ExampleConverter() {
// 	c := temperature.NewConverter()
// 	t := temperature.New(50.0, temperature.Celsius)

// 	r, err := c.Convert(*t, temperature.Fahrenheit)
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	fmt.Println(r)
// 	//Output: 122°F
// }
