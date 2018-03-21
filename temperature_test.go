package temperature

import (
	"testing"
)

type stringer interface {
	String() string
}

func TestNewTemperature(t *testing.T) {
	val := 10.0
	unit := Celsius
	temp := NewTemperature(val, unit)

	if val != temp.Value() {
		t.Fatalf("Expected %v; got %v", val, temp.Value())
	}

	if unit.Text != temp.Unit().Text {
		t.Fatalf("Expected %v; got %v", unit.Text, temp.Unit().Text)
	}
}

func TestString(t *testing.T) {
	tt := []struct {
		name        string
		temperature stringer
		result      string
	}{
		{name: "Print Celsius", temperature: NewCelsius(10), result: "10 °C"},
		{name: "Print Fahrenheit", temperature: NewFahrenheit(10), result: "10 °F"},
		{name: "Print Delisle", temperature: NewDelisle(10), result: "10 °D"},
		{name: "Print Kelvin", temperature: NewKelvin(10), result: "10 K"},
		{name: "Print Reaumur", temperature: NewReaumur(10), result: "10 °Re"},
		{name: "Print Rankine", temperature: NewRankine(10), result: "10 °Ra"},
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

func TestEquals(t *testing.T) {
	temp := NewCelsius(45)
	tt := []struct {
		name   string
		a      Temperature
		b      Temperature
		result bool
	}{
		{name: "Same object", a: temp, b: temp, result: true},
		{name: "Equal same unit", a: NewCelsius(10), b: NewCelsius(10), result: true},
		{name: "Not Equal same unit", a: NewCelsius(10), b: NewCelsius(9), result: false},
		{name: "Not Equal different unit", a: NewCelsius(10), b: NewFahrenheit(9), result: false},
		{name: "Equal different unit", a: NewCelsius(0), b: NewFahrenheit(32), result: true},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			res := Equals(tc.a, tc.b)
			if res != tc.result {
				t.Fatalf("Expected %v; got %v", tc.result, res)
			}
		})
	}
}

func TestSetValue(t *testing.T) {
	temp := NewCelsius(10)
	temp.SetValue(5)

	if temp.Value() != 5 {
		t.Fatalf("Expected %v; got %v", 5, temp.Value())
	}
}

func TestSetUnit(t *testing.T) {
	temp := NewCelsius(0)
	temp.SetUnit(Fahrenheit)

	if temp.Unit().Text != Fahrenheit.Text {
		t.Fatalf("Expected %v; got %v", Fahrenheit.Text, temp.Unit())
	}

	if round(temp.Value(), 2) != 32 {
		t.Fatalf("Expected %v; got %v", 32, temp.Value())
	}
}
