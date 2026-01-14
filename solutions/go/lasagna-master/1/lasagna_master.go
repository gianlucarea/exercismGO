package lasagna

 func PreparationTime(layers []string, avgTimeForLayer int) int {
    if avgTimeForLayer == 0 {
        avgTimeForLayer = 2
    }
    return len(layers) * avgTimeForLayer
}

func Quantities(layers []string) (noodles int, sauce float64) {
    for _ , layer := range layers {
        if layer == "noodles" {
            noodles += 50
        }
        if layer == "sauce" {
            sauce += 0.2
        }
    }
    return noodles, sauce
}

 func AddSecretIngredient(friendsList , myList []string) {
    myList[len(myList) - 1] = friendsList[len(friendsList) - 1]
}

 func ScaleRecipe(quantities []float64, portions int) []float64 {
	multiplier := float64(portions) / 2.0
    result := make([]float64,len(quantities))
    for i, v := range quantities {
        result[i] = v * multiplier
    }
    return result
} 