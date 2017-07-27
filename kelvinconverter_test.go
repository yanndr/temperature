package temperature

// import (
// 	"testing"
// )

// type tempTestCase struct {
// 	name  string
// 	temp  float64
// 	value float64
// }

// func TestCelsuisToKelvinConverter(t *testing.T) {

// 	tt := []tempTestCase{
// 		{name: "zero celsius to kelvin", temp: 0.0, value: 273.15},
// 		{name: "thirty five celsius to kelvin", temp: 35.0, value: 308.15},
// 	}
// 	c := &celsiusKelvinConverter{}

// 	for _, tc := range tt {
// 		t.Run(tc.name, func(t *testing.T) {
// 			res := c.ToKelvin(tc.temp)
// 			if res != tc.value {
// 				t.Fatalf("Expected %v; got %v", tc.value, res)
// 			}
// 		})
// 	}
// }

// func TestKelvinToCelsiusConverter(t *testing.T) {

// 	tt := []tempTestCase{
// 		{name: "zero kelvin to celsius", temp: 0.0, value: -273.15},
// 	}
// 	c := &celsiusKelvinConverter{}

// 	for _, tc := range tt {
// 		t.Run(tc.name, func(t *testing.T) {
// 			res := c.FromKelvin(tc.temp)
// 			if res != tc.value {
// 				t.Fatalf("Expected %v; got %v", tc.value, res)
// 			}
// 		})
// 	}
// }

// func TestFpackage temperature

// import (
// 	"testing"
// )

// type tempTestCase struct {
// 	name  string
// 	temp  float64
// 	value float64
// }

// func TestCelsuisToKelvinConverter(t *testing.T) {

// 	tt := []tempTestCase{
// 		{name: "zero celsius to kelvin", temp: 0.0, value: 273.15},
// 		{name: "thirty five celsius to kelvin", temp: 35.0, value: 308.15},
// 	}
// 	c := &celsiusKelvinConverter{}

// 	for _, tc := range tt {
// 		t.Run(tc.name, func(t *testing.T) {
// 			res := c.ToKelvin(tc.temp)
// 			if res != tc.value {
// 				t.Fatalf("Expected %v; got %v", tc.value, res)
// 			}
// 		})
// 	}
// }

// func TestKelvinToCelsiusConverter(t *testing.T) {

// 	tt := []tempTestCase{
// 		{name: "zero kelvin to celsius", temp: 0.0, value: -273.15},
// 	}
// 	c := &celsiusKelvinConverter{}

// 	for _, tc := range tt {
// 		t.Run(tc.name, func(t *testing.T) {
// 			res := c.FromKelvin(tc.temp)
// 			if res != tc.value {
// 				t.Fatalf("Expected %v; got %v", tc.value, res)
// 			}
// 		})
// 	}
// }

// func TestFahrenheitToKelvinConverter(t *testing.T) {

// 	tt := []tempTestCase{
// 		{name: "zero fahrenheit to kelvin", temp: 0.0, value: 255.37},
// 		{name: "thirty five fahrenheit to kelvin", temp: 35.0, value: 274.82},
// 	}
// 	c := &fahrenheitKelvinConverter{}

// 	for _, tc := range tt {
// 		t.Run(tc.name, func(t *testing.T) {
// 			res := c.ToKelvin(tc.temp)
// 			val := Round(res, 2)
// 			if val != tc.value {
// 				t.Fatalf("Expected %v; got %v", tc.value, val)
// 			}
// 		})
// 	}
// }

// func TestKelvinToFahrenheitConverter(t *testing.T) {

// 	tt := []tempTestCase{
// 		{name: "zero kelvin to fahrenheit", temp: 0.0, value: -459.67},
// 		{name: "thirty five kelvin to fahrenheit", temp: 35.0, value: -396.67},
// 	}
// 	c := &fahrenheitKelvinConverter{}

// 	for _, tc := range tt {
// 		t.Run(tc.name, func(t *testing.T) {
// 			res := c.FromKelvin(tc.temp)
// 			val := Round(res, 2)
// 			if val != tc.value {
// 				t.Fatalf("Expected %v; got %v", tc.value, val)
// 			}
// 		})
// 	}
// }

// func TestRankineToKelvinConverter(t *testing.T) {

// 	tt := []tempTestCase{
// 		{name: "zero Rankine to Kelvin", temp: 0.0, value: 0.0},
// 		{name: "thirty five Rankine to Kelvin", temp: 35.0, value: 19.44},
// 	}
// 	c := rankineKelvinConverter{}

// 	for _, tc := range tt {
// 		t.Run(tc.name, func(t *testing.T) {
// 			res := c.ToKelvin(tc.temp)
// 			val := Round(res, 2)
// 			if val != tc.value {
// 				t.Fatalf("Expected %v; got %v", tc.value, val)
// 			}
// 		})
// 	}
// }

// func TestKelvinToRankineConverter(t *testing.T) {

// 	tt := []tempTestCase{
// 		{name: "zero Kelvin to Rankine", temp: 0.0, value: 0.0},
// 		{name: "fifty five Kelvin to Rankine", temp: 55.0, value: 99},
// 	}
// 	c := rankineKelvinConverter{}

// 	for _, tc := range tt {
// 		t.Run(tc.name, func(t *testing.T) {
// 			res := c.FromKelvin(tc.temp)
// 			val := Round(res, 2)
// 			if val != tc.value {
// 				t.Fatalf("Expected %v; got %v", tc.value, val)
// 			}

// 		})
// 	}
// }

// func TestDelisleToKelvinConverter(t *testing.T) {

// 	tt := []tempTestCase{
// 		{name: "zero Delisle to Kelvin", temp: 0.0, value: 373.15},
// 		{name: "thirty five Delisle to Kelvin", temp: 35.0, value: 349.82},
// 	}
// 	c := delisleKelvinConverter{}

// 	for _, tc := range tt {
// 		t.Run(tc.name, func(t *testing.T) {
// 			res := c.ToKelvin(tc.temp)
// 			val := Round(res, 2)
// 			if val != tc.value {
// 				t.Fatalf("Expected %v; got %v", tc.value, val)
// 			}
// 		})
// 	}
// }

// func TestKelvinToDelisleConverter(t *testing.T) {

// 	tt := []tempTestCase{
// 		{name: "zero Kelvin to Delisle", temp: 0.0, value: 559.72},
// 		{name: "fifty five Kelvin to Delisle", temp: 55.0, value: 477.23},
// 	}
// 	c := delisleKelvinConverter{}

// 	for _, tc := range tt {
// 		t.Run(tc.name, func(t *testing.T) {
// 			res := c.FromKelvin(tc.temp)
// 			val := Round(res, 2)
// 			if val != tc.value {
// 				t.Fatalf("Expected %v; got %v", tc.value, val)
// 			}
// 		})
// 	}
// }

// func TestReaumurToKelvinConverter(t *testing.T) {

// 	tt := []tempTestCase{
// 		{name: "zero Reaumur to Kelvin", temp: 0.0, value: 273.15},
// 		{name: "thirty five Reaumur to Kelvin", temp: 35.0, value: 316.9},
// 	}
// 	c := reaumurKelvinConverter{}

// 	for _, tc := range tt {
// 		t.Run(tc.name, func(t *testing.T) {
// 			res := c.ToKelvin(tc.temp)
// 			val := Round(res, 2)
// 			if val != tc.value {
// 				t.Fatalf("Expected %v; got %v", tc.value, val)
// 			}
// 		})
// 	}
// }

// func TestKelvinToReamurConverter(t *testing.T) {

// 	tt := []tempTestCase{
// 		{name: "zero Kelvin to Reaumur", temp: 0.0, value: -218.52},
// 		{name: "fifty five Kelvin to Reaumur", temp: 55.0, value: -174.52},
// 	}
// 	c := reaumurKelvinConverter{}

// 	for _, tc := range tt {
// 		t.Run(tc.name, func(t *testing.T) {
// 			res := c.FromKelvin(tc.temp)
// 			val := Round(res, 2)
// 			if val != tc.value {
// 				t.Fatalf("Expected %v; got %v", tc.value, val)
// 			}
// 		})
// 	}
// }
// ahrenheitToKelvinConverter(t *testing.T) {

// 	tt := []tempTestCase{
// 		{name: "zero fahrenheit to kelvin", temp: 0.0, value: 255.37},
// 		{name: "thirty five fahrenheit to kelvin", temp: 35.0, value: 274.82},
// 	}
// 	c := &fahrenheitKelvinConverter{}

// 	for _, tc := range tt {
// 		t.Run(tc.name, func(t *testing.T) {
// 			res := c.ToKelvin(tc.temp)
// 			val := Round(res, 2)
// 			if val != tc.value {
// 				t.Fatalf("Expected %v; got %v", tc.value, val)
// 			}
// 		})
// 	}
// }

// func TestKelvinToFahrenheitConverter(t *testing.T) {

// 	tt := []tempTestCase{
// 		{name: "zero kelvin to fahrenheit", temp: 0.0, value: -459.67},
// 		{name: "thirty five kelvin to fahrenheit", temp: 35.0, value: -396.67},
// 	}
// 	c := &fahrenheitKelvinConverter{}

// 	for _, tc := range tt {
// 		t.Run(tc.name, func(t *testing.T) {
// 			res := c.FromKelvin(tc.temp)
// 			val := Round(res, 2)
// 			if val != tc.value {
// 				t.Fatalf("Expected %v; got %v", tc.value, val)
// 			}
// 		})
// 	}
// }

// func TestRankineToKelvinConverter(t *testing.T) {

// 	tt := []tempTestCase{
// 		{name: "zero Rankine to Kelvin", temp: 0.0, value: 0.0},
// 		{name: "thirty five Rankine to Kelvin", temp: 35.0, value: 19.44},
// 	}
// 	c := rankineKelvinConverter{}

// 	for _, tc := range tt {
// 		t.Run(tc.name, func(t *testing.T) {
// 			res := c.ToKelvin(tc.temp)
// 			val := Round(res, 2)
// 			if val != tc.value {
// 				t.Fatalf("Expected %v; got %v", tc.value, val)
// 			}
// 		})
// 	}
// }

// func TestKelvinToRankineConverter(t *testing.T) {

// 	tt := []tempTestCase{
// 		{name: "zero Kelvin to Rankine", temp: 0.0, value: 0.0},
// 		{name: "fifty five Kelvin to Rankine", temp: 55.0, value: 99},
// 	}
// 	c := rankineKelvinConverter{}

// 	for _, tc := range tt {
// 		t.Run(tc.name, func(t *testing.T) {
// 			res := c.FromKelvin(tc.temp)
// 			val := Round(res, 2)
// 			if val != tc.value {
// 				t.Fatalf("Expected %v; got %v", tc.value, val)
// 			}

// 		})
// 	}
// }

// func TestDelisleToKelvinConverter(t *testing.T) {

// 	tt := []tempTestCase{
// 		{name: "zero Delisle to Kelvin", temp: 0.0, value: 373.15},
// 		{name: "thirty five Delisle to Kelvin", temp: 35.0, value: 349.82},
// 	}
// 	c := delisleKelvinConverter{}

// 	for _, tc := range tt {
// 		t.Run(tc.name, func(t *testing.T) {
// 			res := c.ToKelvin(tc.temp)
// 			val := Round(res, 2)
// 			if val != tc.value {
// 				t.Fatalf("Expected %v; got %v", tc.value, val)
// 			}
// 		})
// 	}
// }

// func TestKelvinToDelisleConverter(t *testing.T) {

// 	tt := []tempTestCase{
// 		{name: "zero Kelvin to Delisle", temp: 0.0, value: 559.72},
// 		{name: "fifty five Kelvin to Delisle", temp: 55.0, value: 477.23},
// 	}
// 	c := delisleKelvinConverter{}

// 	for _, tc := range tt {
// 		t.Run(tc.name, func(t *testing.T) {
// 			res := c.FromKelvin(tc.temp)
// 			val := Round(res, 2)
// 			if val != tc.value {
// 				t.Fatalf("Expected %v; got %v", tc.value, val)
// 			}
// 		})
// 	}
// }

// func TestReaumurToKelvinConverter(t *testing.T) {

// 	tt := []tempTestCase{
// 		{name: "zero Reaumur to Kelvin", temp: 0.0, value: 273.15},
// 		{name: "thirty five Reaumur to Kelvin", temp: 35.0, value: 316.9},
// 	}
// 	c := reaumurKelvinConverter{}

// 	for _, tc := range tt {
// 		t.Run(tc.name, func(t *testing.T) {
// 			res := c.ToKelvin(tc.temp)
// 			val := Round(res, 2)
// 			if val != tc.value {
// 				t.Fatalf("Expected %v; got %v", tc.value, val)
// 			}
// 		})
// 	}
// }

// func TestKelvinToReamurConverter(t *testing.T) {

// 	tt := []tempTestCase{
// 		{name: "zero Kelvin to Reaumur", temp: 0.0, value: -218.52},
// 		{name: "fifty five Kelvin to Reaumur", temp: 55.0, value: -174.52},
// 	}
// 	c := reaumurKelvinConverter{}

// 	for _, tc := range tt {
// 		t.Run(tc.name, func(t *testing.T) {
// 			res := c.FromKelvin(tc.temp)
// 			val := Round(res, 2)
// 			if val != tc.value {
// 				t.Fatalf("Expected %v; got %v", tc.value, val)
// 			}
// 		})
// 	}
// }
