package temperature

import (
	"errors"
	"fmt"
)

//Converter is the default implementation of a temperature converter.
type Converter struct {
	Converters map[Scale]KelvinConverter
}

//NewConverter allocate a new Converter with all the default scales converters.
func NewConverter() *Converter {
	tc := &Converter{}
	tc.Converters = make(map[Scale]KelvinConverter)
	tc.Converters[Celsius] = &celsiusKelvinConverter{}
	tc.Converters[Fahrenheit] = &fahrenheitKelvinConverter{}
	tc.Converters[Rankine] = &rankineKelvinConverter{}
	tc.Converters[Delisle] = &delisleKelvinConverter{}
	tc.Converters[Reaumur] = &reaumurKelvinConverter{}
	return tc
}

type noConverterFoundError struct {
	scaleName string
}

func (e *noConverterFoundError) Error() string {
	return fmt.Sprintf("No converter found for Scale %s", e.scaleName)
}

var errEmptyScale = errors.New("Scale can't be empty")

//Convert is the implementation of the convert method for TemperatureConverter.
func (tc *Converter) Convert(t Temperature, u Scale) (*Temperature, error) {
	if t.Scale.Name == "" || u.Name == "" {
		return nil, errEmptyScale
	}

	if t.Scale == u {
		return &t, nil
	}

	var tempInKelvin float64
	if t.Scale == Kelvin {
		tempInKelvin = t.Value
	} else {
		if tc.Converters[t.Scale] == nil {
			return nil, &noConverterFoundError{t.Scale.Name}
		}

		tempInKelvin = tc.Converters[t.Scale].ToKelvin(t.Value)
	}

	if u == Kelvin {
		return &Temperature{Value: tempInKelvin, Scale: u}, nil
	}

	if tc.Converters[u] == nil {
		return nil, &noConverterFoundError{u.Name}
	}

	return &Temperature{Value: tc.Converters[u].FromKelvin(tempInKelvin), Scale: u}, nil
}
