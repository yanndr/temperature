package temperature

//Converter is the default implementation of a temperature converter.
type Converter struct {
	// Converters map[Scale]KelvinConverter
}

//NewConverter allocate a new Converter with all the default scales converters.
// func NewConverter() *Converter {
// 	tc := &Converter{}
// 	tc.Converters = make(map[Scale]KelvinConverter)
// 	tc.Converters[Celsius] = &celsiusKelvinConverter{}
// 	tc.Converters[Fahrenheit] = &fahrenheitKelvinConverter{}
// 	tc.Converters[Rankine] = &rankineKelvinConverter{}
// 	tc.Converters[Delisle] = &delisleKelvinConverter{}
// 	tc.Converters[Reaumur] = &reaumurKelvinConverter{}
// 	return tc
// }

// type noConverterFoundError struct {
// 	scaleName string
// }

// func (e *noConverterFoundError) Error() string {
// 	return fmt.Sprintf("No converter found for Scale %s", e.scaleName)
// }

// var errEmptyScale = errors.New("Scale can't be empty")

//Convert is the implementation of the convert method for TemperatureConverter.
func (tc *Converter) Convert(t TemperatureConvertible, u TemperatureConvertible) TemperatureConvertible {
	return u.FromKelvin(t.ToKelvin())
}
