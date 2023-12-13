package main

import "fmt"

func main() {
	//nutriData structure is defined in the nutriscore.go file
	var sugar, sat, fruit, fibre, protein int
	var variety scoreType
	var energy, sodium float64
	fmt.Println("***********************************************")
	fmt.Println("WELCOME TO KNOW YOUR GRADE YOUR FOOD INTAKE")
	fmt.Println("***********************************************")
	fmt.Println("Enter the value you took in energy in KJ")
	fmt.Scan(&energy)
	fmt.Println("Enter the sugar level in grams")
	fmt.Scan(&sugar)
	fmt.Println("Enter the level of Saturated fatty Acid level in mg")
	fmt.Scan(&sat)
	fmt.Println("Enter the value of sodium level in mg")
	fmt.Scan(&sodium)
	fmt.Println("Enter the value of fruit level in percent")
	fmt.Scan(&fruit)
	fmt.Println("Enter the level of fibre in mg")
	fmt.Scan(&fibre)
	fmt.Println("Enter the level of Protein in mg")
	fmt.Scan(&protein)
	fmt.Println("Enter the type of food either leave as food or cheese or water or beverage ")
	fmt.Scan(&variety)
	ns := getNutritionalScore(nutriData{
		energy:       energyKJ(energy),
		sugar:        sugarGms(sugar),
		satFattyAcid: satFattyAcid(sat),
		sodium:       sodiumMig(sodium),
		fruit:        fruitPercent(fruit),
		fibre:        fibreGms(fibre),
		protein:      protienGms(protein),
	}, variety)
	fmt.Printf("NUTRITIONAL SCORE:%d\n", ns.value)
	fmt.Printf("NUTRISCORE:%s\n", ns.getNutriScore())
	fmt.Println("*******************************************************")
	fmt.Println("THANK YOU AND CONGRADULATIONS FOR YOUR FITNESS JOURNEY")
	fmt.Println("*******************************************************")
}
