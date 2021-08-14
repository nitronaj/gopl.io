package lengthconv

// FeetToMeter converts a Feet to Meter
func FeetToMeters(f Feet) Meters{
	return Meters(f*0.3048)
}


// MeterToFeet converts a Meters to Feet
func MetersToFeet(m Meters) Feet{
	return Feet(m/0.3048)
}

