package weather

import "strconv"

type Fahrenheit int
type Kelvin float64

func (f Fahrenheit) Int() int {
	return int(f)
}

func (f Fahrenheit) String() string {
	return strconv.Itoa(f.Int())
}

func (k Kelvin) ConvtoF() Fahrenheit {
	return Fahrenheit((k-273.24)*9/5 + 32) //nolint:gomnd
}
