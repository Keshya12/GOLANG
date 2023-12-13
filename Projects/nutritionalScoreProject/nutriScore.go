package main

//import "fmt"

type scoreType int

// simailar to define enum (itematic way of writing golang)
const (
	food scoreType = iota //The scoretype is the type we have designed in the before line.
	beverage
	water
	cheese
) //Because the nutritional value for the type of categories will be const for calculation.

type nutritionalScore struct {
	value     int
	positive  int
	negative  int
	scoreType scoreType //The type we have designed in the begining of this file.
}

//The values were based on the nutritional chart refered
var scoreToLetter = []string{"", "THE BEST", "BETTER", "NOT TO TAKE FREQUENTLY", "BETTER TO AVOID", "WORST"}
var energyLevel = []float64{3350, 3015, 2680, 2345, 2010, 1675, 1005, 670, 335}
var sugarLevel = []float64{45, 60, 36, 31, 27, 22.5, 18, 13.5, 9, 4.5}
var satFattyAcidLevel = []float64{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
var sodiumLevel = []float64{900, 810, 720, 630, 540, 450, 360, 270, 180, 90}
var fibreLevel = []float64{4.7, 3.7, 2.8, 1.9, 0.9}
var proteinLevel = []float64{8, 6.4, 4.8, 3.2, 1.6}

var beverageEnergyLevel = []float64{270, 240, 110, 180, 150, 120, 90, 60, 30, 0}
var beverageSugarLevel = []float64{13.5, 12, 10, 9, 7.5, 6, 5, 3, 1.5, 0}

// Defining our own types below
type energyKJ float64
type sugarGms float64
type satFattyAcid float64
type sodiumMig float64
type fruitPercent float64
type fibreGms float64
type protienGms float64

// Structure to get the values from the user to alculate score
type nutriData struct {
	energy       energyKJ
	sugar        sugarGms
	satFattyAcid satFattyAcid
	sodium       sodiumMig
	fruit        fruitPercent
	fibre        fibreGms
	protein      protienGms
	//isWater      bool
}

// struct methods
func (e energyKJ) getPoints(st scoreType) int {
	if st == beverage {
		return getPointsFromRange(float64(e), beverageEnergyLevel)
	}
	return getPointsFromRange(float64(e), energyLevel)
}
func (s sugarGms) getPoints(st scoreType) int {
	if st == beverage {
		return getPointsFromRange(float64(s), beverageSugarLevel)
	}
	return getPointsFromRange(float64(s), sugarLevel)

}
func (sat satFattyAcid) getPoints(st scoreType) int {
	return getPointsFromRange(float64(sat), satFattyAcidLevel)
}
func (sod sodiumMig) getPoints(st scoreType) int {
	return getPointsFromRange(float64(sod), sodiumLevel)
}
func (f fibreGms) getPoints(st scoreType) int {
	return getPointsFromRange(float64(f), fibreLevel)
}
func (fr fruitPercent) getPoints(st scoreType) int {
	if st == beverage {
		if fr > 80 {
			return 10
		} else if fr > 60 {
			return 4
		} else if fr > 40 {
			return 2
		}
		return 0
	}
	if fr > 80 {
		return 5
	} else if fr > 60 {
		return 2
	} else if fr > 40 {
		return 1
	}
	return 0
}
func (p protienGms) getPoints(st scoreType) int {
	return getPointsFromRange(float64(p), proteinLevel)
}

/*func  energyFromKJ(kcal float64) energyKJ {
	return energyKJ(kcal * 4.184)
}
func  sodiumMg(saltMg float64) sodiumMig {
	return sodiumMig(saltMg / 2.5)
}*/
func getNutritionalScore(n nutriData, st scoreType) nutritionalScore {
	value := 0
	positive := 0
	negative := 0
	if st != water {
		fruitPoints := n.fruit.getPoints(st)
		fibrePoints := n.fibre.getPoints(st)

		negative = n.energy.getPoints(st) + n.sugar.getPoints(st) + n.satFattyAcid.getPoints(st) + n.sodium.getPoints(st)
		positive = fibrePoints + n.protein.getPoints(st) + fruitPoints

		if st == cheese {
			value = negative - positive
		} else {
			if negative >= 11 && fruitPoints < 5 {
				value = negative - positive - fruitPoints
			} else {
				value = negative - positive
			}
		}
	}
	return nutritionalScore{
		value:     value,
		positive:  positive,
		negative:  negative,
		scoreType: st, //"st" the variable stores the type of food sent through "ns" variable.
	}
}
func (ns nutritionalScore) getNutriScore() string {
	if ns.scoreType == food {
		return scoreToLetter[getPointsFromRange(float64(ns.value), []float64{18, 10, 2, -1})] //here in float64 we have passed final value of nutritionalScore
		//in the result the scoreletter value after processing the function the value is between 0 to 4 the result is mapped to the appropriate index either A,B,C,D,E
	}
	var isWater bool
	if isWater {
		return scoreToLetter[0]
	}
	return scoreToLetter[getPointsFromRange(float64(ns.value), []float64{9, 5, 1, -2})]
}
func getPointsFromRange(v float64, steps []float64) int {
	lenSteps := len(steps)
	for i, l := range steps {
		if v > l {
			return lenSteps - i
		}
	}
	return 0
}
