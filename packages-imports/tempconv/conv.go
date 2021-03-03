package tempconv

// CToF Converts a Celsius Temp to Fahrenheit
func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)

}

// FToC Converts a Fahrenheit Temp to Celsius
func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

