package temperature

// Unit represents a temperature unit
type Unit struct {
	Text       string
	ToKelvin   func(float64) float64
	FromKelvin func(float64) float64
}

var (
	// Celsius temperature unit.
	Celsius = Unit{
		Text:       "°C",
		ToKelvin:   func(v float64) float64 { return v + 273.15 },
		FromKelvin: func(v float64) float64 { return v - 273.15 },
	}
	//Kelvin temperature unit.
	Kelvin = Unit{
		Text:       "K",
		ToKelvin:   func(v float64) float64 { return v },
		FromKelvin: func(v float64) float64 { return v },
	}
	// Fahrenheit temperature unit.
	Fahrenheit = Unit{
		Text:       "°F",
		ToKelvin:   func(v float64) float64 { return (v + 459.67) * 5 / 9 },
		FromKelvin: func(v float64) float64 { return v*9/5 - 459.67 },
	}
	// Rankine temperature unit.
	Rankine = Unit{
		Text:       "°Ra",
		ToKelvin:   func(v float64) float64 { return v * 5 / 9 },
		FromKelvin: func(v float64) float64 { return v * 9 / 5 },
	}
	// Delisle temperature unit.
	Delisle = Unit{
		Text:       "°D",
		ToKelvin:   func(v float64) float64 { return 373.15 - v*2/3 },
		FromKelvin: func(v float64) float64 { return (373.15 - v) * 3 / 2 },
	}
	// Reaumur temperature unit.
	Reaumur = Unit{
		Text:       "°Re",
		ToKelvin:   func(v float64) float64 { return v*5/4 + 273.15 },
		FromKelvin: func(v float64) float64 { return (v - 273.15) * 4 / 5 },
	}
)

// NewCelsius returns a Celsius temperature initialized with the value.
func NewCelsius(value float64) Temperature {
	return &temperature{value: value, unit: Celsius}
}

// NewKelvin returns a Kelvin temperature initialized with the value.
func NewKelvin(value float64) Temperature {
	return &temperature{value: value, unit: Kelvin}
}

// NewFahrenheit returns a Fahrenheit temperature initialized with the value.
func NewFahrenheit(value float64) Temperature {
	return &temperature{value: value, unit: Fahrenheit}
}

// NewRankine returns a Rankine Temperature nitialized with the value.
func NewRankine(value float64) Temperature {
	return &temperature{value: value, unit: Rankine}
}

// NewDelisle returns a Delisle temperature initialized with the value.
func NewDelisle(value float64) Temperature {
	return &temperature{value: value, unit: Delisle}
}

// NewReaumur returns a Reaumur temperature initialized with the value.
func NewReaumur(value float64) Temperature {
	return &temperature{value: value, unit: Reaumur}
}
