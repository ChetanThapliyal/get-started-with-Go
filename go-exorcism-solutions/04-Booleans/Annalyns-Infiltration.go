
// To print and test the code, uncomment package and func main, and comment out package annalyn.

//package main
package annalyn

// CanFastAttack can be executed only when the knight is sleeping.
func CanFastAttack(knightIsAwake bool) bool {
    return !knightIsAwake
}

// CanSpy can be executed if at least one of the characters is awake.
func CanSpy(knightIsAwake, archerIsAwake, prisonerIsAwake bool) bool {
	return knightIsAwake || archerIsAwake || prisonerIsAwake
}

// CanSignalPrisoner can be executed if the prisoner is awake and the archer is sleeping.
func CanSignalPrisoner(archerIsAwake, prisonerIsAwake bool) bool {
	return !archerIsAwake && prisonerIsAwake
}

// CanFreePrisoner can be executed if the prisoner is awake and the other 2 characters are asleep
// or if Annalyn's pet dog is with her and the archer is sleeping.
func CanFreePrisoner(knightIsAwake, archerIsAwake, prisonerIsAwake, petDogIsPresent bool) bool {
	return (!archerIsAwake && petDogIsPresent) || (!petDogIsPresent && prisonerIsAwake && !archerIsAwake && !knightIsAwake)
}

/* func main () {
	fmt.Println(CanFastAttack(true))
	fmt.Println(CanSpy(true,false,true))
	fmt.Println(CanSignalPrisoner(true,false))
	fmt.Println(CanFreePrisoner(false,false,true,true))
	} 
*/
