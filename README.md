# temperature
[![Build Status](https://travis-ci.org/yanndr/temperature.svg?branch=master)](https://travis-ci.org/yanndr/temperature) [![Go Report Card](https://goreportcard.com/badge/github.com/yanndr/temperature)](https://goreportcard.com/report/github.com/yanndr/temperature)
[![GoDoc](https://godoc.org/github.com/yanndr/temperature?status.svg)](https://godoc.org/github.com/yanndr/temperature)

This is Go library for temperature conversion.
So far, it allows you to convert Kelvin, Celsius, Fahrenheit, Rankine, Reaumur and Delisle unit.
You can also use your own unit as long at the unit is implementing the interfaces ToKelvin and FromKelvin.

## Installing

```
go get github.com/yanndr/temperature
```

## Usage
 
```
c := temperature.New(30,Celsius)

result,_ := temperature.Convert(c, temperature.Fahrenheit)


fmt.Println(result)
// output 86 °F
```

If you want to use a unit that I didn't implement:

```
type newUnit temperature.Unit

const myUnit = newUnit("Y")

func (newUnit) ToKelvin(v float64) float64 {
	return v + 710
}

func (newUnit) FromKelvin(v float64) temperature.Temperature {
	return temperature.New(v-710, myUnit)
}

func main() {
	c := temperature.New(0,temperature.Celsius)

	result,_ := temperature.Convert(c, myUnit)
	fmt.Println(c)
	fmt.Println(result)
}

```
