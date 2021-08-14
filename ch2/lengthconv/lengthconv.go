// Package length performs Feet and Meters conversions.

package lengthconv

import "fmt"

type Feet float64;
type Meters float64;




func (m Meters) String() string {
	return fmt.Sprintf("%g Meters",m)

}

func (f Feet) String() string {
	return fmt.Sprintf("%g Feet",f)
}

