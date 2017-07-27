package temperature

import (
	"testing"
)

func TestConverter(t *testing.T) {

	tt := []struct {
		name  string
		temp  Temperature
		scale Scale
		value float64
		err   string
	}{
		{name: "CtoF zero", temp: Temperature{Value: 0.0, Scale: Celsius}, scale: Fahrenheit, value: 32.0},
		{name: "FtoC 32", temp: Temperature{Value: 0.0, Scale: Celsius}, scale: Fahrenheit, value: 32.0},
		{name: "KtoK 0", temp: Temperature{Value: 0.0, Scale: Kelvin}, scale: Kelvin, value: 0.0},
		{name: "KtoC 0", temp: Temperature{Value: 0.0, Scale: Kelvin}, scale: Celsius, value: -273.15},
		{name: "CtoC 0", temp: Temperature{Value: 0.0, Scale: Celsius}, scale: Celsius, value: 0.0},
		{name: "Ctok 0", temp: Temperature{Value: -273.15, Scale: Celsius}, scale: Kelvin, value: 0.0},

		{name: "RtoR 50", temp: Temperature{Value: 50, Scale: Rankine}, scale: Rankine, value: 50},
		{name: "RtoC 50", temp: Temperature{Value: 50, Scale: Rankine}, scale: Celsius, value: -245.37},
		{name: "RtoF 50", temp: Temperature{Value: 50, Scale: Rankine}, scale: Fahrenheit, value: -409.67},
		{name: "RtoK 50", temp: Temperature{Value: 50, Scale: Rankine}, scale: Kelvin, value: 27.78},
		{name: "RtoD 50", temp: Temperature{Value: 50, Scale: Rankine}, scale: Delisle, value: 518.06},
		{name: "RtoRe 50", temp: Temperature{Value: 50, Scale: Rankine}, scale: Reaumur, value: -196.3},

		{name: "FtoC 50", temp: Temperature{Value: 50, Scale: Fahrenheit}, scale: Celsius, value: 10},
		{name: "FtoF 50", temp: Temperature{Value: 50, Scale: Fahrenheit}, scale: Fahrenheit, value: 50},
		{name: "FtoK 50", temp: Temperature{Value: 50, Scale: Fahrenheit}, scale: Kelvin, value: 283.15},
		{name: "FtoR 50", temp: Temperature{Value: 50, Scale: Fahrenheit}, scale: Rankine, value: 509.67},
		{name: "FtoD 50", temp: Temperature{Value: 50, Scale: Fahrenheit}, scale: Delisle, value: 135},
		{name: "FtoRe 50", temp: Temperature{Value: 50, Scale: Fahrenheit}, scale: Reaumur, value: 8},

		{name: "CtoC 50", temp: Temperature{Value: 50, Scale: Celsius}, scale: Celsius, value: 50},
		{name: "CtoF 50", temp: Temperature{Value: 50, Scale: Celsius}, scale: Fahrenheit, value: 122},
		{name: "CtoK 50", temp: Temperature{Value: 50, Scale: Celsius}, scale: Kelvin, value: 323.15},
		{name: "CtoR 50", temp: Temperature{Value: 50, Scale: Celsius}, scale: Rankine, value: 581.67},
		{name: "CtoD 50", temp: Temperature{Value: 50, Scale: Celsius}, scale: Delisle, value: 75},
		{name: "CtoRe 50", temp: Temperature{Value: 50, Scale: Celsius}, scale: Reaumur, value: 40},

		{name: "KtoC 50", temp: Temperature{Value: 50, Scale: Kelvin}, scale: Celsius, value: -223.15},
		{name: "KtoF 50", temp: Temperature{Value: 50, Scale: Kelvin}, scale: Fahrenheit, value: -369.67},
		{name: "KtoK 50", temp: Temperature{Value: 50, Scale: Kelvin}, scale: Kelvin, value: 50},
		{name: "KtoR 50", temp: Temperature{Value: 50, Scale: Kelvin}, scale: Rankine, value: 90},
		{name: "KtoD 50", temp: Temperature{Value: 50, Scale: Kelvin}, scale: Delisle, value: 484.73},
		{name: "KtoRe 50", temp: Temperature{Value: 50, Scale: Kelvin}, scale: Reaumur, value: -178.52},

		{name: "DtoD 50", temp: Temperature{Value: 50, Scale: Delisle}, scale: Delisle, value: 50},
		{name: "DtoC 50", temp: Temperature{Value: 50, Scale: Delisle}, scale: Celsius, value: 66.67},
		{name: "DtoF 50", temp: Temperature{Value: 50, Scale: Delisle}, scale: Fahrenheit, value: 152},
		{name: "DtoK 50", temp: Temperature{Value: 50, Scale: Delisle}, scale: Kelvin, value: 339.82},
		{name: "DtoR 50", temp: Temperature{Value: 50, Scale: Delisle}, scale: Rankine, value: 611.67},
		{name: "DtoRe 50", temp: Temperature{Value: 50, Scale: Delisle}, scale: Reaumur, value: 53.33},

		{name: "RetoD 50", temp: Temperature{Value: 50, Scale: Reaumur}, scale: Delisle, value: 56.25},
		{name: "RetoC 50", temp: Temperature{Value: 50, Scale: Reaumur}, scale: Celsius, value: 62.5},
		{name: "RetoF 50", temp: Temperature{Value: 50, Scale: Reaumur}, scale: Fahrenheit, value: 144.5},
		{name: "RetoK 50", temp: Temperature{Value: 50, Scale: Reaumur}, scale: Kelvin, value: 335.65},
		{name: "RetoR 50", temp: Temperature{Value: 50, Scale: Reaumur}, scale: Rankine, value: 604.17},
		{name: "RetoRr 50", temp: Temperature{Value: 50, Scale: Reaumur}, scale: Reaumur, value: 50},

		{name: "Empty temp", scale: Fahrenheit, err: "Scale can't be empty"},
		{name: "Empty scale", temp: Temperature{Value: 0.0, Scale: Celsius}, err: "Scale can't be empty"},
		{name: "No input converter", temp: Temperature{Value: 0.0, Scale: Scale{Name: "Test", Symbol: "T"}}, scale: Celsius, err: "No converter found for Scale Test"},
		{name: "No output converter", temp: Temperature{Value: 0.0, Scale: Celsius}, scale: Scale{Name: "Test", Symbol: "T"}, err: "No converter found for Scale Test"},
	}

	c := NewConverter()

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			res, err := c.Convert(tc.temp, tc.scale)

			if err != nil {
				if err.Error() != tc.err {
					t.Fatalf("Expected error %s; got %s.", tc.err, err.Error())
				}
				return
			}

			val := Round(res.Value, 2)
			if val != tc.value {

				t.Fatalf("Expected %v; got %v", tc.value, val)
			}

			if res.Scale != tc.scale {
				t.Fatalf("Expected %v; got %v", tc.scale, res.Scale)
			}
		})
	}

}
