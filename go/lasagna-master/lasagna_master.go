package lasagna

func PreparationTime(layers []string, mins int) int {
	if mins == 0 {
		return len(layers) * 2
	}

	return len(layers) * mins
}

func Quantities(layers []string) (noodles int, sauce float64) {
	for i := 0; i < len(layers); i++ {
		if layers[i] == "noodles" {
			noodles += 50
		}

		if layers[i] == "sauce" {
			sauce += 0.2
		}
	}

	return noodles, sauce
}

func AddSecretIngredient(friendsList, myList []string) {
	(myList)[len(myList)-1] = friendsList[len(friendsList)-1]
}

func ScaleRecipe(quantities []float64, portions int) []float64 {
	newQuantities := make([]float64, len(quantities))

	for i := 0; i < len(quantities); i++ {
		newQuantities[i] = quantities[i] * float64(portions) / 2
	}

	return newQuantities
}
