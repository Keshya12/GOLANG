package main

import "fmt"

func main() {
	//nutriData structure is defined in the nutriscore.go file
	ns := getNutritionalScore(nutriData{
		energy:       energyKJ(0),
		sugar:        sugarGms(10),
		satFattyAcid: satFattyAcid(2),
		sodium:       sodiumMg(500),
		fruit:        fruitPercent(60),
		fibre:        fibreGms(4),
		protein:      protienGms(2),
	}, food)
	fmt.Printf("NUTRITIONAL SCORE:%d\n", ns.value)
	fmt.Printf("NUTRISCORE:%d\n", ns.getNutriScore())
}
