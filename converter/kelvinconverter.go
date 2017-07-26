package converter

// KelvinConverter is the interface that define the Kelvin conversion.
//
// Each Converters use by the temperature converter must implement these methods to convert the unit to Kelvin and from Kelvin.
type KelvinConverter interface {
	FromKelvin(float64) float64
	ToKelvin(float64) float64
}

type celsiusKelvinConverter struct {
}

func (*celsiusKelvinConverter) FromKelvin(t float64) float64 {
	return t - 273.15
}

func (*celsiusKelvinConverter) ToKelvin(t float64) float64 {
	return t + 273.15
}

type fahrenheitKelvinConverter struct {
}

func (*fahrenheitKelvinConverter) FromKelvin(t float64) float64 {
	return t*9/5 - 459.67
}

func (*fahrenheitKelvinConverter) ToKelvin(t float64) float64 {
	return (t + 459.67) * 5 / 9
}

type rankineKelvinConverter struct {
}

func (*rankineKelvinConverter) FromKelvin(t float64) float64 {
	return t * 9 / 5
}

func (*rankineKelvinConverter) ToKelvin(t float64) float64 {
	return (t) * 5 / 9
}

type delisleKelvinConverter struct {
}

func (*delisleKelvinConverter) FromKelvin(t float64) float64 {
	return (373.15 - t) * 3 / 2
}

func (*delisleKelvinConverter) ToKelvin(t float64) float64 {
	return 373.15 - t*2/3
}

type reaumurKelvinConverter struct {
}

func (*reaumurKelvinConverter) FromKelvin(t float64) float64 {
	return (t - 273.15) * 4 / 5
}

func (*reaumurKelvinConverter) ToKelvin(t float64) float64 {
	return t*5/4 + 273.15
}
