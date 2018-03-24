package temperature

type (
	//Unit is a temperature unit type.
	Unit           string
	CelsiusUnit    Unit
	KelvinUnit     Unit
	FahrenheitUnit Unit
	RankineUnit    Unit
	DelisleUnit    Unit
	ReaumurUnit    Unit

	//Convertible provides function to convert from and to Kelvin unit.
	Convertible interface {
		ToKelvin(float64) float64
		FromKelvin(float64) Temperature
	}
)

const (
	//Celsius temperature unit
	Celsius = CelsiusUnit("°C")
	//Kelvin temperature unit
	Kelvin = KelvinUnit("K")
	//Fahrenheit temperature unit
	Fahrenheit = FahrenheitUnit("°F")
	//Rankine temperature unit
	Rankine = RankineUnit("°Ra")
	//Delisle temperature unit
	Delisle = DelisleUnit("°D")
	//Reaumur temperature unit
	Reaumur = ReaumurUnit("°Re")
)

func (CelsiusUnit) ToKelvin(v float64) float64 {
	return v + 273.15
}

func (CelsiusUnit) FromKelvin(v float64) Temperature {
	return &temperature{v: v - 273.15, unit: Celsius}
}

func (KelvinUnit) ToKelvin(v float64) float64 {
	return v
}

func (KelvinUnit) FromKelvin(v float64) Temperature {
	return &temperature{v: v, unit: Kelvin}
}

func (FahrenheitUnit) ToKelvin(v float64) float64 {
	return (v + 459.67) * 5 / 9
}

func (FahrenheitUnit) FromKelvin(v float64) Temperature {
	return &temperature{v: v*9/5 - 459.67, unit: Fahrenheit}
}

func (RankineUnit) ToKelvin(v float64) float64 {
	return v * 5 / 9
}

func (RankineUnit) FromKelvin(v float64) Temperature {
	return &temperature{v: v * 9 / 5, unit: Rankine}
}

func (DelisleUnit) ToKelvin(v float64) float64 {
	return 373.15 - v*2/3
}

func (DelisleUnit) FromKelvin(v float64) Temperature {
	return &temperature{v: (373.15 - v) * 3 / 2, unit: Delisle}
}

func (ReaumurUnit) ToKelvin(v float64) float64 {
	return v*5/4 + 273.15
}

func (ReaumurUnit) FromKelvin(v float64) Temperature {
	return &temperature{v: (v - 273.15) * 4 / 5, unit: Reaumur}
}
