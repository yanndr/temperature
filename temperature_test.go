package temperature_test

// import "testing"
// import "github.com/yanndr/temperature"

// func TestNew(t *testing.T) {

// 	tt := []struct {
// 		name  string
// 		scale temperature.Scale
// 		value float64
// 	}{
// 		{"New Celsius", temperature.Celsius, 10},
// 		{"New Fahrenheit", temperature.Fahrenheit, 30.4},
// 		{"New Delisle", temperature.Delisle, -10.543},
// 		{"New Kelvin", temperature.Kelvin, 100.56},
// 		{"New Reaumur", temperature.Reaumur, 10000},
// 		{"New Rankine", temperature.Rankine, 0.67},
// 	}

// 	for _, tc := range tt {
// 		t.Run(tc.name, func(t *testing.T) {
// 			temp := temperature.New(tc.value, tc.scale)

// 			if temp == nil {
// 				t.Fatal("Temperature is not supposed to be nil.")
// 			}

// 			if temp.Scale != tc.scale {
// 				t.Fatalf("Expected scale %v; got %v", temperature.Celsius, temp.Scale)
// 			}

// 			if temp.Value != tc.value {
// 				t.Fatalf("Expected value %v; got %v", tc.value, temp.Value)
// 			}
// 		})
// 	}
// }

// func TestString(t *testing.T) {
// 	tt := []struct {
// 		name        string
// 		temperature temperature.Temperature
// 		result      string
// 	}{
// 		{name: "Print Celsius", temperature: temperature.Temperature{Value: 10, Scale: temperature.Celsius}, result: "10°C"},
// 		{name: "Print Fahrenheit", temperature: temperature.Temperature{Value: 10, Scale: temperature.Fahrenheit}, result: "10°F"},
// 		{name: "Print Delisle", temperature: temperature.Temperature{Value: 10, Scale: temperature.Delisle}, result: "10°D"},
// 		{name: "Print Kelvin", temperature: temperature.Temperature{Value: 10, Scale: temperature.Kelvin}, result: "10°K"},
// 		{name: "Print Reaumur", temperature: temperature.Temperature{Value: 10, Scale: temperature.Reaumur}, result: "10°Re"},
// 		{name: "Print Rankine", temperature: temperature.Temperature{Value: 10, Scale: temperature.Rankine}, result: "10°R"},
// 	}

// 	for _, tc := range tt {
// 		t.Run(tc.name, func(t *testing.T) {
// 			res := tc.temperature.String()
// 			if res != tc.result {
// 				t.Fatalf("Expected %v; got %v", tc.result, res)
// 			}

// 		})
// 	}
// }
