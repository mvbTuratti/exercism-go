package lasagna

// PreparationTime returns the time required to prepare based on the amount of layers.
func PreparationTime(recepies []string, preparationTime int) int {
	if preparationTime == 0 {
		preparationTime = 2
	}
	return len(recepies) * preparationTime
}

// Quantities function receives an array of strings, returns a tuple containing the quantities
// of noodles and the amount of sauce required.
func Quantities(items []string) (int, float64) {
	var noodle int
	var sauce float64
	for _, item := range items {
		switch item {
		case "noodles":
			noodle += 50
		case "sauce":
			sauce += 0.2
		}
	}
	return noodle, sauce
}

// AddSecretIngredient updates last element of my list with my friends value.
func AddSecretIngredient(friendList, myList []string) {
	myLastElement := len(myList) - 1
	friendsLastElement := len(friendList) - 1
	if myLastElement >= 0 && friendsLastElement >= 0 {
		myList[myLastElement] = friendList[friendsLastElement]
	}
}

// ScaleRecipe function.
func ScaleRecipe(quantities []float64, number int) []float64 {
	modifiedQuantities := []float64{}
	for _, quantity := range quantities {
		modifiedQuantities = append(modifiedQuantities, quantity/2*float64(number))
	}
	return modifiedQuantities
}
