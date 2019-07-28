// Package raindrops is a modified FizzBuzz
package raindrops

import "strconv"

// Convert takes a number and returns a string depending on the number's factors
func Convert(number int) string {

	output := ""

	if number%3 == 0 {
		output += "Pling"
	}
	if number%5 == 0 {
		output += "Plang"
	}
	if number%7 == 0 {
		output += "Plong"
	}

	if output != "" {
		return output
	}

	return strconv.Itoa(number)
}
