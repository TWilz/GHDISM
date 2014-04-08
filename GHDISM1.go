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
//v30 improved output formatting
//v32 scalable to any number of dz states AND any number of cycle
//GHDISM1 works for any number of diseases. note, the diseases do NOT interact yet
//GHDISM2 attempting to make diseases interact
 
package main

import (
	"fmt"
	"math/rand"
)



func main() {
	fmt.Println("\n\n")
	
	const numDzs = 5
	fmt.Println("Number of Diseases:",numDzs,"\n\n")
	
	const maxNumDzStates = 20
	fmt.Println("Max number of Disease States:",maxNumDzStates,"\n\n")
	
	const maxNumCycles = 20
	fmt.Println("Max number of Cycles:",maxNumCycles,"\n\n")//db
	
	type individualpath [numDzs][maxNumDzStates][maxNumCycles]int
	var person1 individualpath

	var numDzStates [numDzs]int
	var numCycles [numDzs]int
	//each disease has a different number of disease states and # of cycles
	
	for q:=0; q<numDzs; q++ {
		numDzStates[q] = 5 //default
		//this is where we'd alter the number of disease states for each, e.g.
		numDzStates[1] = 8
		numCycles[q] = 10 //default
		//this is where we'd alter the number of disease cycles for each e.g.
		numCycles[1] = 11
		fmt.Println("For disease number",q,"there are ",numDzStates[q],"number of disease states")
		fmt.Println("For disease number",q,"there are ",numCycles[q],"number of cycles")
		}
	fmt.Println("\n\n")
	
	type matrix [numDzs][maxNumDzStates][maxNumDzStates]float32  
	// A transition matrix which is actually an 3D array
	//where [i][j] is i-th row and j-th column
	var TM matrix
	
	

/*
last make the person1 to an array so that way you can have x number of individuals
for now to make more than 1 dz, make the indivual path have an additional parameter corresponding 
to the dz and also have the TM have an additional parameter for the dz state.  
Then have the interactions be a function which operates on those variables (e.. the Interactions array)
The problem with this approach is that we have a n x m x o array and that since all diseases dont
have the same number of disease states nor the same number of desired cycles, some of these data 
will be 0.  i'm not sure if that will screw up things or not.  i think no as long as when doing the output 
and calculations, we use the vectors numdzstates[q], numcycles[q] rather than the length of the array

*/
for q:=0; q<numDzs; q++ {


//initializes the TM to all risk of progressions to be 1/numDzStates
	for i := 0; i < numDzStates[q]; i++ {
		for j := 0; j < numDzStates[q]; j++ {						
			TM[q][i][j]=float32(1.0/float32(numDzStates[q]))
//			fmt.Println("DEBUGGING:",float32(1.0/float32(numDzStates[q])))
			}
		}
//

//Here is where you would alter the TM values if desired.

//Prints the TM
	fmt.Println("For disease ",q," the Transition Matrix is:")
	fmt.Println("--------------------------------------------------------------------------")		
	for i := 0; i< numDzStates[q]; i++ {
		for k :=0; k<numDzStates[q]; k++ {
			fmt.Print(TM[q][i][k],"\t")
			}
		fmt.Println("\n")
		}
//


	
//horizontal cumulative TM:
	var sumTM matrix
	for r:=0; r<numDzStates[q]; r++ {
		sumTM[q][r][0]=TM[q][r][0]		
		}
	for r:=0; r<numDzStates[q]; r++ {
		for c:=1; c<numDzStates[q]; c++ {
			sumTM[q][r][c]=sumTM[q][r][c-1]+TM[q][r][c]
			}
		}
	fmt.Println("The horizontal summative values of the TM are:")
	fmt.Println("--------------------------------------------------------------------------")		
	for i := 0; i< numDzStates[q]; i++ {
		for k :=0; k<numDzStates[q]; k++ {
			fmt.Print(sumTM[q][i][k],"\t")
			}
		fmt.Println("\n")
		}
//


//Random numbers. r is array of length number of cycles of psuedo-random numbers b/n 0 and 1 
// whatever number goes inside of the Seed function will determine random #'s. can use time to randomize
//this is inside the for loop for each disease, so each disease q has an array of random numbers
//each corresponding to one of the cycles.  each disease has a different randome number for each cycle.
	rand.Seed(10)
	type randomMatrix [numDzs][maxNumCycles]float32
	
	var random randomMatrix
	
	for o :=0; o<numCycles[q]; o++ {
		random[q][o] = rand.Float32()
		}
	fmt.Println("The random numbers are:")
	for o :=0; o<numCycles[q]; o++ {
		fmt.Print(random[q][o],"\t")	
		}



//Initialization - cycle 0 tests TM against random number r[0] to see what dz state the person should start in		
//makes vertical cumulative TM for initializing person in cycle 0
var verticalSumTM matrix
verticalSumTM[q][0][0] = TM[q][0][0]
for r:=1; r<numDzStates[q]; r++ {
	verticalSumTM[q][r][0]=verticalSumTM[q][r-1][0]+TM[q][r][0]
	}
//checks what dz state person should start in
if (random[q][0] <= verticalSumTM[q][0][0]) {
	person1[q][0][0]=1
	}
for r:=1; r<numDzStates[q]; r++ {
	if ((random[q][0] > verticalSumTM[q][r-1][0]) && (random[q][0] <= verticalSumTM[q][r][0])) {
		person1[q][r][0]=1
	}
}



//Determines person's path through dz states
fmt.Println("\n\nDetermining the individual's path...") //debugging
for cycle:=1; cycle<numCycles[q]; cycle++ {	
//Determines if person belongs in TM[0][1] i.e. cycle 1, dz state 0. WORKS!!
	for i:=0; i<numDzStates[q]; i++ {
	if (person1[q][i][cycle-1]==1 && random[q][cycle]<=sumTM[q][i][0]) {
			person1[q][0][cycle] = 1
			}
	}
//Determine if person belongs in TM[x][1] i.e. cycle 1, dz state 1 through numDzStates. WORKS!!
	for o:=1; o<numDzStates[q]; o++ {
		for j:=0; j<numDzStates[q]; j++ {
			if (person1[q][j][cycle-1]==1 && random[q][cycle]>sumTM[q][j][o-1] && random[q][cycle]<=sumTM[q][j][o]) {
				person1[q][o][cycle] = 1
				}
		}
	}
}
//


//Prints person's path in readable format
	fmt.Println("For disease",q,", the individual's path is:")
	fmt.Println("\t\tCycle")
	fmt.Print("\t")
	for i := 0; i< numCycles[q]; i++ {
		fmt.Print("\t",i)
		}
	fmt.Println("\n-------------------------------------------------------------------------------------------")	
	for i := 0; i< numDzStates[q]; i++ {
		fmt.Print("DzState",i,"\t")
		for k :=0; k<numCycles[q]; k++ {
			fmt.Print(person1[q][i][k],"\t")
			}
		fmt.Println("\n")
		}
//

}//ends for loop for number of diseases, variable q


}//end main

/*
To Do:
1. Change formatting on output.  How to print the matrices in more user friendly manner?
2. Change array from being set up in program to being imported from csv
3. Make 1st col and row or TM the names of dz states
4. Make initialization scalable to any number of dz states. 


person1 represents a person going through the world with susceptibility to various diseases. it is a instance of
the variable individualpath.  so eventually person[] will be an array of individualpaths

every other variable works for a given disease.  should we just make ever variable a vector corresponding to the
number disease that it is? that would probably be the simpliest. so the variables are 
numDzStates, numCycles, TM, and variations off of TM - sumTM both horizontal and vertical. 
*/
     


