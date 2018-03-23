package temperature

import (
	"testing"
)

func TestNew(t *testing.T) {
	val := 10.0
	unit := Celsius
	temp := New(val, unit)

	if val != temp.Value() {
		t.Fatalf("Expected %v; got %v", val, temp.Value())
	}

	if unit != temp.Unit() {
		t.Fatalf("Expected %v; got %v", unit, temp.Unit())
	}
}

func TestNewWithHandler(t *testing.T) {
	val := 10.0
	changeVal := 15.0
	unit := Celsius
	handlerCalled := false

	f := func(temp Temperature) {
		handlerCalled = true

		if changeVal != temp.Value() {
			t.Fatalf("Expected %v; got %v", val, temp.Value())
		}

		if unit != temp.Unit() {
			t.Fatalf("Expected %v; got %v", unit, temp.Unit())
		}
	}

	temp := NewWithHandler(val, unit, f)

	if val != temp.Value() {
		t.Fatalf("Expected %v; got %v", val, temp.Value())
	}

	if unit != temp.Unit() {
		t.Fatalf("Expected %v; got %v", unit, temp.Unit())
	}

	temp.SetTemperature(New(changeVal, Celsius))

	if !handlerCalled {
		t.Fatalf("Handler not called Expected %v; got %v", true, handlerCalled)
	}
}

func TestSetTemperateChangeHandler(t *testing.T) {
	val := 10.0
	changeVal := 15.0
	unit := Celsius
	handlerCalled := false

	f := func(temp Temperature) {
		handlerCalled = true

		if changeVal != temp.Value() {
			t.Fatalf("Expected %v; got %v", val, temp.Value())
		}

		if unit != temp.Unit() {
			t.Fatalf("Expected %v; got %v", unit, temp.Unit())
		}
	}

	temp := New(val, unit)

	if val != temp.Value() {
		t.Fatalf("Expected %v; got %v", val, temp.Value())
	}

	if unit != temp.Unit() {
		t.Fatalf("Expected %v; got %v", unit, temp.Unit())
	}

	temp.SetTemperateChangeHandler(f)
	temp.SetTemperature(New(changeVal, Celsius))

	if !handlerCalled {
		t.Fatalf("Handler not called Expected %v; got %v", true, handlerCalled)
	}
}
func TestString(t *testing.T) {
	tt := []struct {
		name        string
		temperature Stringer
		result      string
	}{
		{name: "Print Celsius", temperature: New(10, Celsius), result: "10 °C"},
		{name: "Print Fahrenheit", temperature: New(10, Fahrenheit), result: "10 °F"},
		{name: "Print Delisle", temperature: New(10, Delisle), result: "10 °D"},
		{name: "Print Kelvin", temperature: New(10, Kelvin), result: "10 K"},
		{name: "Print Reaumur", temperature: New(10, Reaumur), result: "10 °Re"},
		{name: "Print Rankine", temperature: New(10, Rankine), result: "10 °Ra"},
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
	temp := New(45, Celsius)
	tt := []struct {
		name   string
		a      Temperature
		b      Temperature
		result bool
	}{
		{name: "Same object", a: temp, b: temp, result: true},
		{name: "Equal same unit", a: New(10, Celsius), b: New(10, Celsius), result: true},
		{name: "Not Equal same unit", a: New(10, Celsius), b: New(9, Celsius), result: false},
		{name: "Not Equal different unit", a: New(10, Celsius), b: New(9, Kelvin), result: false},
		{name: "Equal different unit", a: New(0, Celsius), b: New(273.15, Kelvin), result: true},
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
	temp := New(10, Celsius)
	temp.SetValue(5)

	if temp.Value() != 5 {
		t.Fatalf("Expected %v; got %v", 5, temp.Value())
	}
}

func TestSetUnit(t *testing.T) {
	temp := New(0, Celsius)
	temp.SetUnit(Fahrenheit)

	if temp.Unit() != Fahrenheit {
		t.Fatalf("Expected %v; got %v", Fahrenheit, temp.Unit())
	}

	if round(temp.Value(), 2) != 32 {
		t.Fatalf("Expected %v; got %v", 32, temp.Value())
	}
}
