package helpers

func EnergyFromKcal(kcal float64) float64 {
	return kcal * 4.184
}

func SugarGram(saltMg float64) float64 {
	return saltMg / 2.5
}

func SaturatedFattyAcid(sfa float64) float64 {
	return sfa
}

func SodiumMilligram(sd float64) float64 {
	return sd
}

func FibreGram(fbg float64) float64 {
	return fbg
}

func FruitsPercent(fp float64) float64 {
	return fp
}

func ProteinGram(pg float64) float64 {
	return pg
}
