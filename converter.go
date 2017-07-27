package temperature

import "fmt"

func (c Celsius) ToKelvin() Kelvin {
	return Kelvin{Temperature: Temperature{Value: c.GetValue() + 273.15}}
}

func (c *Celsius) FromKelvin(v Valuer) {
	c.Value = v.GetValue() - 273.15
}

func (c Celsius) String() string {
	return fmt.Sprintf("%v°C", Round(c.Value, 2))
}

func (k Kelvin) ToKelvin() Kelvin {
	return k
}

func (k *Kelvin) FromKelvin(v Valuer) {
	k.Value = v.GetValue()
}

func (k Kelvin) String() string {
	return fmt.Sprintf("%vK", Round(k.Value, 2))
}

func (f Fahrenheit) ToKelvin() Kelvin {
	return NewKelvin((f.Value + 459.67) * 5 / 9)
}

func (f *Fahrenheit) FromKelvin(v Valuer) {
	f.Value = v.GetValue()*9/5 - 459.67
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%v°F", Round(f.Value, 2))
}
