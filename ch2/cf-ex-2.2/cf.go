// Cf converts its numeric argument to Celsius and Fahrenheit.
package main

import (
	"flag"
	"fmt"

	"mygopl.io/ch2/lengthconv"
	"mygopl.io/ch2/tempconv"
)


var t = flag.Float64("t", 0, "performs Celsius and Fahrenheit conversions")
var tt = flag.String("tt", "", "temprater type C = Celsius or F=Fahrenheit")
var l = flag.Float64("l", 0, "performs Feet and Meters conversions")
var lt = flag.String("lt", "", "length type F=Feet and M=Meters")

func tempConversion(t float64, tt string) {
	var c tempconv.Celsius;
	var f tempconv.Fahrenheit;


	if(tt == "C") {
        c = tempconv.Celsius(t)
		fmt.Printf("%s = %s,\n", c, tempconv.CToF(c))

	} else {
		f = tempconv.Fahrenheit(t)
		fmt.Printf("%s = %s,\n", f, tempconv.FToC(f))
	}
}


func lengthConversion(l float64, lt string) {
	var m lengthconv.Meters;
	var f lengthconv.Feet;


	if(lt == "F") {
        f = lengthconv.Feet(l)
		fmt.Printf("%s = %.3g,\n", f, lengthconv.FeetToMeters(f))

	} else {
		m = lengthconv.Meters(l)
		fmt.Printf("%s = %s,\n", m, lengthconv.MetersToFeet(m))
	}
}

func main() {
	flag.Parse()
	if(*tt == "C" || *tt == "F") {
		tempConversion(*t, *tt)
	}

	if(*lt == "M" || *lt == "F") {
		lengthConversion(*l, *lt)
	}
}
