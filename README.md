# temperature
[![Build Status](https://travis-ci.org/yanndr/temperature.svg?branch=master)](https://travis-ci.org/yanndr/temperature) [![Go Report Card](https://goreportcard.com/badge/github.com/yanndr/temperature)](https://goreportcard.com/report/github.com/yanndr/temperature)
[![GoDoc](https://godoc.org/github.com/yanndr/temperature?status.svg)](https://godoc.org/github.com/yanndr/temperature)

This is Go library for temperature conversion.
It allow you to convert Kelvin, Celsius, Fahrenheit, Rankine, Reaumur and Delisle unit so far.
You can use your own unit as long at the unit is implementing the interfaces ToKelvin and FromKelvin and Valuer.

# Usage
 
```
c := temperature.NewCelsius(30)
f := temperature.Fahrenheit{}

temperature.Convert(c, &f)

fmt.Println(f)t
// output 86 °F
```

If you want to use your a unit that I didn't implemented:

```
func main() {
	c := temperature.NewCelsius(0)
	k := temperature.NewKelvin(366)
	n := newUnit{}

	temperature.Convert(c, &n)
	temperature.Convert(c, &k)
	fmt.Println(c)
	fmt.Println(k)
	fmt.Println(n)
  //output:
  //0 °C
  //273.15 K
  //278.15 °Y
}

type newUnit struct {
	temperature.Temperature
}

func (n newUnit) ToKelvin() temperature.Kelvin {
	return temperature.NewKelvin(n.Value - 5)
}

func (n *newUnit) FromKelvin(t temperature.Valuer) {
	n.Value = t.GetValue() + 5
}

func (n newUnit) String() string {
	return fmt.Sprintf("%v °Y", temperature.Round(n.Value, 2))
}
```
