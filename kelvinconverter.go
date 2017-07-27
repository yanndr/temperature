package temperature

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
