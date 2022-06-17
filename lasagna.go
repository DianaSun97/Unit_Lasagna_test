package lasagna

// TODO: define the 'OvenTime' constant

//Define the expected oven time in minutes
//
//Define the OvenTime constant with how many minutes the lasagna should be in the oven.
//According to the cooking book, the expected oven time in minutes is 40:

const OvenTime = 40
const layerTime = 2

//Calculate the remaining oven time in minutes

//Define the RemainingOvenTime() function that takes the actual minutes the lasagna has been in
//the oven as a parameter and returns how many minutes the lasagna still has to remain in the oven,
//based on the expected oven time in minutes from the previous task.

//RemainingOvenTime returns the remaining minutes based on the `actual` minutes already in the oven.
func RemainingOvenTime(tttime int) int {
	return OvenTime - tttime
}

//Calculate the preparation time in minutes

//Define the PreparationTime function that takes the number of layers you added to the
//lasagna as a parameter and returns how many minutes you spent preparing the lasagna,
//assuming each layer takes you 2 minutes to prepare.

// PreparationTime calculates the time needed to prepare the lasagna based on the amount of layers.
func PreparationTime(numLayers int) int {
	return layerTime * numLayers
}

//Calculate the elapsed working time in minutes
//
//Define the ElapsedTime function that takes two parameters: the first parameter is the number
//of layers you added to the lasagna, and the second parameter is the number of minutes the lasagna
//has been in the oven. The function should return how many minutes in total
//you've worked on cooking the lasagna, which is the sum of the preparation time in minutes,
//and the time in minutes the lasagna has spent in the oven at the moment.

// ElapsedTime calculates the total time needed to create and bake a lasagna.
func ElapsedTime(numLayers, tttime int) int {
	return PreparationTime(numLayers) + tttime
}
