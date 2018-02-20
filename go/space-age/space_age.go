// Package space implements the functiolity of transforming a given age in seconds into Earth-Years in all planets of our solar system
package space

// Age takes seconds and a Planet and return the seconds in Earth-Years
func Age(seconds float64, planet Planet) float64 {
	return seconds / planetsOrbitalPeriod()[planet]
}

func planetsOrbitalPeriod() map[Planet]float64 {
	planetsOrbitalPeriod := make(map[Planet]float64)
	planetsOrbitalPeriod["Earth"] = 31557600
	planetsOrbitalPeriod["Mercury"] = planetsOrbitalPeriod["Earth"] * 0.2408467
	planetsOrbitalPeriod["Venus"] = planetsOrbitalPeriod["Earth"] * 0.61519726
	planetsOrbitalPeriod["Mars"] = planetsOrbitalPeriod["Earth"] * 1.8808158
	planetsOrbitalPeriod["Jupiter"] = planetsOrbitalPeriod["Earth"] * 11.862615
	planetsOrbitalPeriod["Saturn"] = planetsOrbitalPeriod["Earth"] * 29.447498
	planetsOrbitalPeriod["Uranus"] = planetsOrbitalPeriod["Earth"] * 84.016846
	planetsOrbitalPeriod["Neptune"] = planetsOrbitalPeriod["Earth"] * 164.79132
	return planetsOrbitalPeriod
}

type Planet string
