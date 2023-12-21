package population

func NbYear(p0 int, percent float64, aug int, p int) (years int) {
	population := p0

	for population < p {
		populationPercent := float64(population) * (percent / 100)
		population = population + int(populationPercent) + aug

		years++
	}

	return years
}