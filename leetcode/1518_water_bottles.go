package main

func WaterBottles(numBottles, numExchange int) int {
	numDrank := numBottles
	for numBottles >= numExchange {
		fullBottles, remainder := numBottles/numExchange, numBottles%numExchange
		numDrank += fullBottles
		numBottles = fullBottles + remainder
	}

	return numDrank
}
//
// func waterBottlesDrankRecurse(fullBottles, emptyBottles, numExchange, numDrank int) int {
// 	numDrank += fullBottles
// 	if fullBottles+emptyBottles >= numExchange {
// 		numBottles, remainder := fullBottles/numExchange, fullBottles%numExchange
// 		return waterBottlesDrankRecurse(numBottles, remainder+emptyBottles, numExchange, numDrank)
// 	}
// 	return numDrank
// }
