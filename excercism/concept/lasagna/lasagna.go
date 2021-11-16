package lasagna

// TODO: define the 'OvenTime' constant

const OvenTime int = 40

// RemainingOvenTime returns the remaining minutes based on the `actual` minutes already in the oven.
func RemainingOvenTime(actualMinutesInOven int) int {

	return OvenTime - actualMinutesInOven

	// panic("RemainingOvenTime not implemented")
}

// PreparationTime calculates the time needed to prepare the lasagna based on the amount of layers.
func PreparationTime(numberOfLayers int) int {

	const timePerLayer int = 2

	return timePerLayer * numberOfLayers

	// panic("PreparationTime not implemented")
}

// ElapsedTime calculates the total time needed to create and bake a lasagna.
func ElapsedTime(numberOfLayers, actualMinutesInOven int) int {

	prepTimeInMins := PreparationTime(numberOfLayers)

	return prepTimeInMins + actualMinutesInOven

	// panic("ElapsedTime not implemented")
}
