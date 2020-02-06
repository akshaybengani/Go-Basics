package main

import "fmt"

func main() {
	var batman string = "I am batman"
	fmt.Println(batman)

	var superman string = "I am superman"
	fmt.Println(superman)

	thor := "I am thor"
	fmt.Println(thor)

	thorRating := 3.
	fmt.Printf("%v, %T", thorRating, thorRating)

	var Ironman, CatAmerica string = "I am Ironman", "I am capt America"
	fmt.Println("\n", Ironman, CatAmerica)

	var defaultValue bool
	fmt.Println(defaultValue)

	var (
		spiderman = "I am Spiderman"
		age       = 18
		powers    = "Web Slings, spidy Sense"
		antman    = "I am Antman"
	)
	fmt.Println(spiderman, age, powers, antman)

}
