package controllers

import "github.com/bostigger/nutrigoplus/models"

type NutritionData struct {
	models.NutritionalData
}
type NutritionScore struct {
	models.NutritionalScore
}

var energyLevels = []float64{3350, 3015, 2680, 2345, 2010, 1675, 1340, 10005, 670, 335}
var sodiumLevels = []float64{900, 810, 720, 630, 540, 450, 360, 270, 180, 90}
var SaturattedFattyAcidsLevels = []float64{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
var sugarLevels = []float64{45, 60, 36, 21, 27, 22.5, 18, 13, 5, 9, 4.5}
var fibreLevels = []float64{4.7, 3.7, 2.8, 1.9, 0.9}
var proteinLevels = []float64{8, 6.4, 4.8, 3.2, 1.6}

var energyBeverageLevels = []float64{270, 240, 210, 180, 150, 120, 90, 60, 30, 0}
var sugarBeverageLevels = []float64{13.5, 12, 10.5, 9, 7.5, 6, 4.5, 3, 1.5, 0}

var scoreLetter = []string{"A", "B", "C", "D", "E", "F"}

func getPointsFromRange(p float64, charts []float64) int {
	chartsLen := len(charts)
	scorePoint := 0
	for i, value := range charts {
		if p > value {
			return chartsLen - i
		}
	}
	return scorePoint
}

func (n *NutritionData) getProteinPoints(scoreType models.ScoreType) int {
	return getPointsFromRange(n.ProteinGram, proteinLevels)
}
func (n *NutritionData) getFibrePoints(scoreType models.ScoreType) int {

	return getPointsFromRange(n.FiberGram, fibreLevels)
}

func (n *NutritionData) getFruitPoints(scoreType models.ScoreType) int {
	if scoreType == models.Beverage {
		if n.Fruits > 80 {
			return 10
		} else if n.Fruits > 60 {
			return 4
		} else if n.Fruits > 40 {
			return 2
		}
		return 0
	}
	if n.Fruits > 80 {
		return 5
	} else if n.Fruits > 60 {
		return 2
	} else if n.Fruits > 40 {
		return 1
	}
	return 0
}
func (n *NutritionData) getEnergyPoints(scoreType models.ScoreType) int {
	if scoreType == models.Beverage {
		return getPointsFromRange(n.Energy, energyBeverageLevels)
	}
	return getPointsFromRange(n.Energy, energyLevels)
}
func (n *NutritionData) getSugarPoints(scoreType models.ScoreType) int {
	if scoreType == models.Beverage {
		return getPointsFromRange(n.SugarGram, sugarBeverageLevels)
	}
	return getPointsFromRange(n.SugarGram, sugarLevels)
}
func (n *NutritionData) getSodiumPoints(scoreType models.ScoreType) int {
	return getPointsFromRange(n.SodiumMilliGram, sodiumLevels)
}
func (n *NutritionData) getSaturatedFatsPoints(scoreType models.ScoreType) int {
	return getPointsFromRange(n.SaturatedFattyAcidsGram, SaturattedFattyAcidsLevels)
}

func NutritionalScoreCalc(nd NutritionData, st models.ScoreType) *NutritionScore {

	value := 0
	positive := 0
	negative := 0

	if st != models.Water {
		positive = nd.getProteinPoints(st) + nd.getFibrePoints(st) + nd.getFruitPoints(st)
		negative = nd.getSaturatedFatsPoints(st) + nd.getSodiumPoints(st) + nd.getSugarPoints(st) + nd.getEnergyPoints(st)

		if st == models.Cheese {
			value = negative - positive
		} else {
			if negative >= 11 && nd.getFruitPoints(st) < 5 {
				value = negative - positive - nd.getFruitPoints(st)
			} else {
				value = negative - positive
			}
		}
	}

	return &NutritionScore{
		models.NutritionalScore{
			Value:     value,
			Positive:  positive,
			Negative:  negative,
			ScoreType: st,
		},
	}
}

func (ns *NutritionScore) GetNutriScore() string {
	if ns.ScoreType == models.Food {
		return scoreLetter[getPointsFromRange(float64(ns.Value), []float64{18, 10, 2, -1})]
	}
	if ns.ScoreType == models.Water {
		return scoreLetter[0]
	}
	return scoreLetter[getPointsFromRange(float64(ns.Value), []float64{9, 5, 1, -2})]
}
