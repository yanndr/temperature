# temperature
[![Build Status](https://travis-ci.org/yanndr/temperature.svg?branch=master)](https://travis-ci.org/yanndr/temperature) [![Go Report Card](https://goreportcard.com/badge/github.com/yanndr/temperature)](https://goreportcard.com/report/github.com/yanndr/temperature)
[![GoDoc](https://godoc.org/github.com/yanndr/temperature?status.svg)](https://godoc.org/github.com/yanndr/temperature)

This is Go library for temperature conversion.
It allow you to convert Kelvin, Celsius, Fahrenheit, Rankine, Reaumur and Delisle unit so far.
You can use your own unit as long at the unit is implementing the interfaces ToKelvin and FromKelvin and Valuer.

## Installing

```
go get github.com/yanndr/temperature
```

## Usage
 
```
c := temperature.NewCelsius(30)

result,_ := temperature.Convert(c, temperature.Fahrenheit)


fmt.Println(f)
// output 86 °F
```

If you want to use your a unit that I didn't implemented:

```
var myUnit = temperature.Unit{
	Text:       "°Y",
	FromKelvin: func(v float64) float64 { return v + 710 },
	ToKelvin:   func(v float64) float64 { return v - 710 },
}

func main() {
	c := temperature.NewCelsius(0)

	result,_ := temperature.Convert(c, myUnit)
	fmt.Println(c)
	fmt.Println(result)
}

```
