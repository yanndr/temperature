package temperature

type (
	//Unit is a temperature unit type.
	Unit       string
	celsius    Unit
	kelvin     Unit
	fahrenheit Unit
	rankine    Unit
	delisle    Unit
	reaumur    Unit

	//Convertible provides function to convert from and to Kelvin unit.
	Convertible interface {
		ToKelvin(float64) float64
		FromKelvin(float64) Temperature
	}
)

const (
	//Celsius temperature unit
	Celsius = celsius("°C")
	//Kelvin temperature unit
	Kelvin = kelvin("K")
	//Fahrenheit temperature unit
	Fahrenheit = fahrenheit("°F")
	//Rankine temperature unit
	Rankine = rankine("°Ra")
	//Delisle temperature unit
	Delisle = delisle("°D")
	//Reaumur temperature unit
	Reaumur = reaumur("°Re")
)

func (celsius) ToKelvin(v float64) float64 {
	return v + 273.15
}

func (celsius) FromKelvin(v float64) Temperature {
	return &temperature{v: v - 273.15, unit: Celsius}
}

func (kelvin) ToKelvin(v float64) float64 {
	return v
}

func (kelvin) FromKelvin(v float64) Temperature {
	return &temperature{v: v, unit: Kelvin}
}

func (fahrenheit) ToKelvin(v float64) float64 {
	return (v + 459.67) * 5 / 9
}

func (fahrenheit) FromKelvin(v float64) Temperature {
	return &temperature{v: v*9/5 - 459.67, unit: Fahrenheit}
}

func (rankine) ToKelvin(v float64) float64 {
	return v * 5 / 9
}

func (rankine) FromKelvin(v float64) Temperature {
	return &temperature{v: v * 9 / 5, unit: Rankine}
}

func (delisle) ToKelvin(v float64) float64 {
	return 373.15 - v*2/3
}

func (delisle) FromKelvin(v float64) Temperature {
	return &temperature{v: (373.15 - v) * 3 / 2, unit: Delisle}
}

func (reaumur) ToKelvin(v float64) float64 {
	return v*5/4 + 273.15
}

func (reaumur) FromKelvin(v float64) Temperature {
	return &temperature{v: (v - 273.15) * 4 / 5, unit: Reaumur}
}
