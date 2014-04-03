//v12 can do 0th cycle of a dz with 7 states and place the individual in the correct initialized dz state in accordance with risks from TM's 1st column
//v13 has minor output formatting improvements.  still a long way to go in terms of making the output look pretty!
//v16 output in readable format
//v17 functional sum of TM's
//v18 bugs in attempt to make it functional to any number of cycles. 
//v19 works for the first two states of cycle 1
//v20 not scalable for all the dz states.
//v21 cycle 1, row 1 works
//v23 cycle 1, row 1 works with for loop rather than if statements
//v26 cycle 1, row 1 AND row 2 work with for loops
//v27 follows 1 person through any number of cycles
//v28 randomize to assigned seed - see line 55
package main

import (
	"fmt"
	"math/rand"
)



func main() {
	fmt.Println("\n\n")
	const numDzStates = 7
	type matrix [numDzStates][numDzStates]float32  // A transition matrix: number of dz states * number disease states,
	//where [i][j] is i-th row and j-th column
	var TM matrix

//

//initializes the TM to all risk of progressions to be 1/numDzStates
	fmt.Println("The Transition Matrix is:\t")
	for i := 0; i < numDzStates; i++ {
		for j := 0; j < numDzStates; j++ {			
			TM[i][j]=float32(1.0/numDzStates)
			}
		}


//	fmt.Println(TM,"\n")
//Place where you could assign TM to be whatever desired


//individ is a matrix where the 
//1st parameter is dz state (1=yes, 0=no), 
//2nd parameter is # cycles person goes through
	const numCycles = 5
	type individualpath [numDzStates][numCycles]int
	var person1 individualpath



//Random numbers. r is array of length number of cycles of psuedo-random numbers b/n 0 and 1 
// whatever number goes inside of the Seed function will determine random #'s. can use time to randomize
	rand.Seed(1000)
	var r [numCycles]float32
	for o :=0; o<numCycles; o++ {
		r[o] = rand.Float32()
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

	

//Initialization - cycle 0 tests TM against random number r[0] to see what dz state the person should start in		
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





//cumulative TM:
	var sumTM matrix
	for r:=0; r<numDzStates; r++ {
		sumTM[r][0]=TM[r][0]		
		}
	for r:=0; r<numDzStates; r++ {
		for c:=1; c<numDzStates; c++ {
			sumTM[r][c]=sumTM[r][c-1]+TM[r][c]
			}
		}
	fmt.Println("The summative values of the TM are:")
	fmt.Println("\n\n")	
	for i := 0; i< numDzStates; i++ {
		for k :=0; k<numDzStates; k++ {
			fmt.Print(sumTM[i][k],"\t")
			}
		fmt.Println("\n")
		}
//
	



//determines person's path through dz states
	fmt.Println("Number of Dz States:",numDzStates)//db
	fmt.Println("Determining the individual's path:") //debugging

	cycle :=1
	fmt.Println("Cycle:",cycle)
	fmt.Println(person1[4][0])
	
//	r[cycle]=0.1//db
//	person1[4][0]=0//db
//	person1[0][0]=1//db

for cycle:=1; cycle<numCycles; cycle++ {	
//Determines if person belongs in TM[0][1] i.e. cycle 1, dz state 0. WORKS!!
	for i:=0; i<numDzStates; i++ {
	if (person1[i][cycle-1]==1 && r[cycle]<=sumTM[i][0]) {
			person1[0][cycle] = 1
			}
	}
//Determine if person belongs in TM[x][1] i.e. cycle 1, dz state 1 through numDzStates. WORKS!!
	for o:=1; o<numDzStates; o++ {
		for j:=0; j<numDzStates; j++ {
			if (person1[j][cycle-1]==1 && r[cycle]>sumTM[j][o-1] && r[cycle]<=sumTM[j][o]) {
				person1[o][cycle] = 1
				}
		}
	}
}

	
//	if (person1[1][cycle-1]==1 && r[cycle]>sumTM[1][0] && r[cycle]<=sumTM[1][1]) {
//		person1[1][cycle] = 1
//		}
//	if (person1[2][cycle-1]==1  && r[cycle]>sumTM[2][0] && r[cycle]<=sumTM[2][1]) {
//		person1[1][cycle] = 1
//		}
	
//	for i:=1; i<numDzStates; i++ {
//		if (person1[i][cycle-1]==1 && r[cycle]>sumTM[i-1][0] && r[cycle]<=sumTM[i][0]) {
//			person1[0][cycle] = 1
//			}
//		}	


//Determines if person belongs in TM[1][1] i.e. cycle 1, dz state 1. WORKS!
//	for i:=0; i<numDzStates; i++ {	
//		if (person1[i][cycle-1]==1 && r[cycle]>sumTM[i][0] && r[cycle]<=sumTM[i][1]) {
//			person1[1][cycle] = 1
//			}
//		}
		
		
//		if (person1[i][cycle-1]==1 && r[cycle]>sumTM[i][0] && r[cycle]<=sumTM[i][1]) {
//			person1[1][cycle] = 1
//			}
//		if (person1[i][cycle-1]==1 && r[cycle]>sumTM[i][0] && r[cycle]<=sumTM[i][1]) {
//			person1[1][cycle] =1
//			}
//		}

//		if (person1[1][cycle-1]==1 && r[cycle]>sumTM
	
	
/*		
	for run:=1;run<numDzStates;run++ {
		if (person1[run][cycle-1]==1 && r[cycle]>sumTM[run-1][0] && r[cycle]<=sumTM[run][0]) {
			person1[0][cycle]=1
			}
		}
		


	if (person1[0][cycle-1]==1 && r[cycle]<=sumTM[0][1]) {
		person1[1][cycle] = 1
		fmt.Print("Here is bug!")//db
		}
	for run:=1;run<numDzStates;run++ {
		if (person1[run][cycle-1]==1 && r[cycle]>sumTM[run-1][1] && r[cycle]<=sumTM[run][1]) {
			person1[1][cycle]=1
			fmt.Println("Here is bug! @ run=",run)//db
			}
		}
*/
//THE sum TM's should be lateral sum's not vertical. 			
		
/*	
	if (person1[0][c-1]==1 && r[c]<=sumTM[0][0]) {
		person1[0][c] = 1
		}

	if (person1[1][c-1]==1 && r[c]>sumTM[0][0] && r[c]<=sumTM[1][0]) {
		person1[0][c] = 1
		}
	if (person1[2][c-1]==1 && r[c]>sumTM[1][0] && r[c]<=sumTM[2][0]) {
		person1[0][c] = 1
		}
	if (person1[3][c-1]==1 && r[c]>sumTM[2][0] && r[c]<=sumTM[3][0]) {
		person1[0][c] = 1
		}
	if (person1[4][c-1]==1 && r[c]>sumTM[3][0] && r[c]<=sumTM[4][0]) {
		person1[0][c] = 1
		}
	if (person1[5][c-1]==1 && r[c]>sumTM[4][0] && r[c]<=sumTM[5][0]) {
		person1[0][c] = 1
		}
	if (person1[6][c-1]==1 && r[c]>sumTM[5][0] && r[c]<=sumTM[6][0]) {
		person1[0][c] = 1
		}
*/



//Prints person's path in readable format
	fmt.Println("The person's path is:")
	fmt.Println("Cycle")
	for i := 0; i< numCycles; i++ {
		fmt.Print(i,"\t")
		}
	fmt.Println("\n----------------------------------------")	
	for i := 0; i< numDzStates; i++ {
		for k :=0; k<numCycles; k++ {
			fmt.Print(person1[i][k],"\t")
			}
		fmt.Println("\n")
		}
//


}
/*
To Do:
1. Change formatting on output.  How to print the matrices in more user friendly manner?
2. Change array from being set up in program to being imported from csv
3. Make 1st col and row or TM the names of dz states
4. Make initialization scalable to any number of dz states. 
*/
     


