package models

type ScoreType int

const (
	Food ScoreType = iota
	Beverage
	Water
	Cheese
)

type NutritionalScore struct {
	Value     int
	Positive  int
	Negative  int
	ScoreType ScoreType
}

type NutritionalData struct {
	Energy                  float64
	SugarGram               float64
	SaturatedFattyAcidsGram float64
	SodiumMilliGram         float64
	FiberGram               float64
	ProteinGram             float64
	Fruits                  float64
	isWater                 bool
}
