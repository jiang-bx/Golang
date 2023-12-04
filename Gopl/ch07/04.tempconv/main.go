package main

import (
	"flag"
	"fmt"
)

type Celsius float64

type Fahrenheit float64

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9.9/5.0 + 32.0)
}

func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32.0) * 5.0 / 9.9)
}

func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}

// celsiusFlag 实现了 Set , 以及默认的 String 
type celsiusFlag struct {
	Celsius
}

func (f *celsiusFlag) Set(s string) error {
	var uint string
	var value float64

	fmt.Sscanf(s, "%f%s", &value, &uint)

	switch uint {
	case "C", "°C":
		f.Celsius = Celsius(value)
		return nil
	case "F", "°F":
		f.Celsius = FToC(Fahrenheit(value))
		return nil
	}

	return fmt.Errorf("invalid temperature %q", s)
}

func CelsiusFlag(name string, value Celsius, usage string) *Celsius {
	f := celsiusFlag{value}
	// f.Set()
	// f.String() 

	// Var(value Value, name string, usage string)
	
	// type Value interface {
	// 	String() string
	// 	Set(string) error
	// }

	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}

var temp = CelsiusFlag("temp", 20.0, "the temperature")

func main() {
	flag.Parse()
	fmt.Println(*temp)
}

// go build main.go

// ./main -temp -18C
// ./main -temp 212F