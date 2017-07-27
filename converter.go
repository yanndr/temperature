package temperature

import "fmt"

func (c Celsius) ToKelvin() Kelvin {
	return Kelvin{Temperature: Temperature{Value: c.GetValue() + 273.15}}
}

func (c *Celsius) FromKelvin(v Valuer) {
	c.Value = v.GetValue() - 273.15
}

func (c Celsius) String() string {
	return fmt.Sprintf("%v °C", Round(c.Value, 2))
}

func (k Kelvin) ToKelvin() Kelvin {
	return k
}

func (k *Kelvin) FromKelvin(v Valuer) {
	k.Value = v.GetValue()
}

func (k Kelvin) String() string {
	return fmt.Sprintf("%v K", Round(k.Value, 2))
}

func (f Fahrenheit) ToKelvin() Kelvin {
	return NewKelvin((f.Value + 459.67) * 5 / 9)
}

func (f *Fahrenheit) FromKelvin(v Valuer) {
	f.Value = v.GetValue()*9/5 - 459.67
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%v °F", Round(f.Value, 2))
}

func (r Rankine) ToKelvin() Kelvin {
	return NewKelvin(r.Value * 5 / 9)
}

func (r *Rankine) FromKelvin(v Valuer) {
	r.Value = v.GetValue() * 9 / 5
}

func (r Rankine) String() string {
	return fmt.Sprintf("%v °Ra", Round(r.Value, 2))
}

func (d Delisle) ToKelvin() Kelvin {
	return NewKelvin(373.15 - d.Value*2/3)
}

func (d *Delisle) FromKelvin(v Valuer) {
	d.Value = (373.15 - v.GetValue()) * 3 / 2
}

func (d Delisle) String() string {
	return fmt.Sprintf("%v °D", Round(d.Value, 2))
}

func (r Reaumur) ToKelvin() Kelvin {
	return NewKelvin(r.Value*5/4 + 273.15)
}

func (r *Reaumur) FromKelvin(v Valuer) {
	r.Value = (v.GetValue() - 273.15) * 4 / 5
}

func (r Reaumur) String() string {
	return fmt.Sprintf("%v °Re", Round(r.Value, 2))
}
