package main

import (
	"fmt"
	"math/rand"
)



func main() {

	const numDzStates = 5
	type matrix [numDzStates][numDzStates]float64  // A transition matrix: number of dz states * number disease states
	var TM matrix

//initializes all of the TM to 1's
	for i := 1; i < numDzStates; i++ {
		for j :=1; j < numDzStates; j++ {			
			TM[i][j]=1
			}
		}


//individ is a matrix where the 
//1st parameter is dz state (1=yes, 0=no), 
//2nd parameter is # cycles person goes through
	const numCycles = 2
	type individ [numDzStates][numCycles]int
//	var person1 individ



//Random numbers. r is array of length number of cycles of psuedo-random numbers b/n 0 and 1 
	var r [numCycles]float64
	for o :=0; o<numCycles; o++ {
		r[o] = rand.Float64()
		}
	fmt.Println("The random numbers are:", r)	




	
//	func Float64() float64
	
//	fmt.Println(Float64())

//	fmt.Println(TM[1][1])


    
}