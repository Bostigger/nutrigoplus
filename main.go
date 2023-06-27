package main

import (
	"fmt"
	"github.com/bostigger/nutrigoplus/controllers"
	"github.com/bostigger/nutrigoplus/helpers"
	"github.com/bostigger/nutrigoplus/models"
)

func main() {

	fmt.Println("NutriGoPLus : Nutrition Score Calculator")
	ns := controllers.NutritionalScoreCalc(controllers.NutritionData{
		NutritionalData: models.NutritionalData{
			Energy:                  helpers.EnergyFromKcal(0),
			SugarGram:               helpers.SugarGram(10),
			SaturatedFattyAcidsGram: helpers.SaturatedFattyAcid(0),
			SodiumMilliGram:         helpers.SodiumMilligram(550),
			FiberGram:               helpers.FibreGram(4.5),
			ProteinGram:             helpers.ProteinGram(30),
		},
	}, models.Food)

	fmt.Printf("Your Nutritional Score is %d\n", ns.Value)
	fmt.Printf("Your Nutritional Grade is %s\n", ns.GetNutriScore())
}
