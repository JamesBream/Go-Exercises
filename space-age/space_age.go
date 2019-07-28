// Package space calculates ages on different planets
package space

import "log"

// Planet type with name as string
type Planet string

// Age takes an age in seconds and a planet to calculate an age based on that planet
func Age(seconds float64, planet Planet) float64 {

	years := seconds / 31557600.0

	// Switching and returning here without extra assignment appears to be the most efficient
	// solution - map lookup is ~300ms slower to bench - interesting.
	switch planet {
	case "Earth":
		return years
	case "Mercury":
		return years / 0.2408467

	case "Venus":
		return years / 0.61519726

	case "Mars":
		return years / 1.8808158

	case "Jupiter":
		return years / 11.862615

	case "Saturn":
		return years / 29.447498

	case "Uranus":
		return years / 84.016846

	case "Neptune":
		return years / 164.79132

	default:
		log.Printf("Unsupported planet %s, returning age in Earth years", planet)
		return years
	}
}
