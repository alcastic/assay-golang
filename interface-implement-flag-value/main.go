package main

import (
	"flag"
	"fmt"
)

type Celsius float64

func (c *Celsius) String() string { // This makes comcrete named type "Celcius" be a fmt.Stringer
	return fmt.Sprintf("%f°C", *c)
}

func fahrenheitToCelsius(f float64) float64 {
	return (f - 32) / 1.8
}

func kelvinToCelsius(k float64) float64 {
	return k - 273.15
}

type celsiusValue struct {
	Celsius // celsiusFlagValue is also a fmt.Stringer
}

func (c *celsiusValue) Set(s string) error {
	var value float64
	var unit string
	fmt.Sscanf(s, "%f%s", &value, &unit)
	switch unit {
	case "°C":
		c.Celsius = Celsius(value)
		return nil
	case "°F":
		c.Celsius = Celsius(fahrenheitToCelsius(value))
		return nil
	case "°K":
		c.Celsius = Celsius(kelvinToCelsius(value))
		return nil
	}
	return fmt.Errorf("invalid input")
}

func CelsiusFlag(name string, value Celsius, usage string) *Celsius {
	flagValue := celsiusValue{value}
	flag.CommandLine.Var(&flagValue, name, usage)
	return &flagValue.Celsius
}

var temperature *Celsius = CelsiusFlag("temp", 20, "usage")

func main() {
	flag.Parse()
	fmt.Println(temperature)
}
