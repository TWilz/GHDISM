//v12 can do 0th cycle of a dz with 7 states and place the individual in the correct initialized dz state in accordance with risks from TM's 1st column
//v13 has minor output formatting improvements.  still a long way to go in terms of making the output look pretty!
package main

import (
	"fmt"
	"math/rand"
)



func main() {
	fmt.Println("\n\n")
	const numDzStates = 7
	type matrix [numDzStates][numDzStates]float64  // A transition matrix: number of dz states * number disease states
	var TM matrix


//initializes the TM to all risk of progressions to be 1/numDzStates
	fmt.Println("The Transition Matrix is:\t")
	for i := 0; i < numDzStates; i++ {
		for j := 0; j < numDzStates; j++ {			
			TM[i][j]=float64(1.0/numDzStates)
			}
		}
	fmt.Println(TM,"\n")


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
	fmt.Println("The random numbers are:\t", r, "\n")	



// Individual 1 "person1"
	fmt.Println("The initialized progression matrix is\t")
	fmt.Println(person1, "\n")


/*
cycle 0:
if random number is between 0 and TM[0][1], then 
=B7*IF(C4<=B20,1,0) + B8*IF(C4<=B21,1,0) + B9*IF(C4<=B22,1,0) + B10*IF(C4<=B23,1,0) + B11*IF(C4<=B24,1,0) + B12*IF(C4<=B25,1,0) + B13*IF(C4<=B26,1,0) + B14*IF(C4<=B27,1,0)
*/

//testing output	
//	fmt.Println(r[0])
//	TM[0][0] = 5
//end testing output

	

//cycle 0 tests TM against random number r[0] to see what dz state the person should start in		
		if (r[0] > 0 && r[0] <= TM[0][0]) { 
			person1[0][0] = 1 
		}
		if (r[0] > TM[0][0] && r[0] <= TM[0][1]) {
			person1[1][0] = 1
		}
		if (r[0] > (TM[0][0]+TM[0][1]) && r[0] <= (TM[0][0]+TM[0][1]+TM[0][2])) {
			person1[2][0] = 1
		}	
		if (r[0] > (TM[0][0]+TM[0][1]+TM[0][2]) && r[0] <= (TM[0][0]+TM[0][1]+TM[0][2]+TM[0][3])) {
			person1[3][0] = 1
		}
		if (r[0] > (TM[0][0]+TM[0][1]+TM[0][2]+TM[0][3]) && r[0] <= (TM[0][0]+TM[0][1]+TM[0][2]+TM[0][3]+TM[0][4])) {
			person1[4][0] = 1
		}
		if (r[0] > (TM[0][0]+TM[0][1]+TM[0][2]+TM[0][3]+TM[0][4]) && r[0] <= (TM[0][0]+TM[0][1]+TM[0][2]+TM[0][3]+TM[0][4]+TM[0][5])) {
			person1[5][0] = 1
		}
		if (r[0] > (TM[0][0]+TM[0][1]+TM[0][2]+TM[0][3]+TM[0][4]+TM[0][5]) && r[0] <= (TM[0][0]+TM[0][1]+TM[0][2]+TM[0][3]+TM[0][4]+TM[0][5])+TM[0][6]) {
			person1[6][0] = 1
		}



//		 	for k:=1; k<numDzStates; k++ {
//						if (r[0] > TM[0][k] && r[0] <= TM[0][k+1]) {/* person1[0][k]=1 */} 
//			}		
		

//	else { person1[0][0] = 100 }

	fmt.Println("The person's path is:\t",person1,"\n")

}
/*
To Do:
1. Change formatting on output.  How to print the matrices in more user friendly manner?
2. Change array from being set up in program to being imported from csv
3. Make 1st col and row or TM the names of dz states
*/
     


