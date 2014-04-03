package main

import (
	"fmt"
	"math/rand"
)



func main() {

	const numDzStates = 7
	type matrix [numDzStates][numDzStates]float64  // A transition matrix: number of dz states * number disease states
	var TM matrix


//initializes the TM to all risk of progressions to be 1/numDzStates
	fmt.Println("The Transition Matrix is:")
	for i := 0; i < numDzStates; i++ {
		for j := 0; j < numDzStates; j++ {			
			TM[i][j]=float64(1.0/numDzStates)
			}
		}
	fmt.Println(TM)


//individ is a matrix where the 
//1st parameter is dz state (1=yes, 0=no), 
//2nd parameter is # cycles person goes through
	const numCycles = 2
	type individualpath [numDzStates][numCycles]int
	var person1 individualpath



//Random numbers. r is array of length number of cycles of psuedo-random numbers b/n 0 and 1 
	var r [numCycles]float64
	for o :=0; o<numCycles; o++ {
		r[o] = rand.Float64()
		}
	fmt.Println("The random numbers are:", r)	



// Individual 1 "person1"
	fmt.Println("The person's path is:")
	fmt.Println(person1)


/*
cycle 0:
if random number is between 0 and TM[0][1], then 
=B7*IF(C4<=B20,1,0) + B8*IF(C4<=B21,1,0) + B9*IF(C4<=B22,1,0) + B10*IF(C4<=B23,1,0) + B11*IF(C4<=B24,1,0) + B12*IF(C4<=B25,1,0) + B13*IF(C4<=B26,1,0) + B14*IF(C4<=B27,1,0)
*/




}
/*
To Do:
1. Change formatting on output.  How to print the matrices in more user friendly manner?
2. Change array from being set up in program to being imported from sv
*/
     


