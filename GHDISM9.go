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
 //GHDISM4 has state series for graphing
 //GHDISM5 graphing - failed, doesn't work. 
 //GHDISM6 put into functions. REALIZATION: the person is initialized randomly based upon 0th row of TM. therefore the 0th row shold
 //represent the prevalence of disease states. and the 1st column should represent the risk of going from dz state to dz state!!!
 //GHDISM7 the remaining parts of code put into functions
 //GHDISM8 adding a utility to each disease state. map a utility to each numDzStates[each dz]. then if numcycles[q]. not yet implemented
 //GHDISM9 switch intervention to true intervention and not just another disease. 
package main

import (
	"fmt"
	"math/rand"
//	"code.google.com/p/plotinum/plot"
//  "code.google.com/p/plotinum/plotter"
//    "code.google.com/p/plotinum/plotutil"
	)

//Declaration of global variables
	const numDiseases = 20
	const maxNumDzs = 20
// const numInterventions = 1
	const maxNumDzStates = 20
	const numDzs = 20
	const maxNumCycles = 20
	type individualpath [numDzs][maxNumDzStates][maxNumCycles]int
	var person1 individualpath
	type stateSeries [numDzs][maxNumCycles]int
	var state1 stateSeries
	var numDzStates [numDzs]int
	var numCycles [numDzs]int
	type matrix [numDzs][maxNumDzStates][maxNumDzStates]float32  
	// A transition matrix which is actually an 3D array
	//where [i][j] is i-th row and j-th column
	var TM matrix
	type interactionData [numDzs][numDzs][maxNumDzStates][maxNumDzStates]float32
	var AA interactionData
	
	type interventionData [numInterventions][maxNumDzs][maxNumDzState][maxNumDzStates]float32
	var interventions interventionsData
	//an intervention can affect the TM values of every dz. the interventions won't be able to affect other interventions.
	
	var q int //q is used to denote the disease number
	//local function variables take the name of the global variable but with an f appended on e.g. q becomes qf when used locally
	var TMI matrix //TMI = transitions matrix with interactions =  TM x corresponding AA value. TMI and not TM should be used when computing a person's path
	var sumTMI matrix //matrix of TMI where each row is horizontally summed i.e. row 1 = row 0 + row 1, etc. this is used for backend calculations and will likely not need to be outputted.
	var verticalSumTMI matrix
	type randomMatrix [numDzs][maxNumCycles]float32
	var random randomMatrix
	type utilityMatrix [numDzs][maxNumDzStates]float32
	var u utilityMatrix
	
func main() {
	printSizeOfModelStructure()
	defineNumDzStatesAndNumCyclesForEachDz()
	initializeValuesOfTM()
	printTM(0)
//	printAllTMs()
	assignValuesToAllAssociationArrays()
	calculateTMI()
	printTMI(0)
//	printAllTMI()
	calculateSumTMI()
	randomize()
	initializePeopleToStartStates()
	calculatePath()
//All calculations are done by here. Now it's just output. 
	printPath(0)
//	printAllPaths()
	printStateSeries(0)
//	printAllStateSeries()


//Utility parts
//define utility values
//for q:=0; q<


	fmt.Println("\n")
}//end main





//THE FUNCTIONS:	
func printSizeOfModelStructure() {
	fmt.Println("\n\n")
	fmt.Println("Number of Diseases:",numDiseases,"\n\n")
	fmt.Println("Number of Interventions:",numInterventions,"\n\n")
	fmt.Println("Max number of Disease States:",maxNumDzStates,"\n\n")
	fmt.Println("Max number of Cycles:",maxNumCycles,"\n\n")//db
}

func defineNumDzStatesAndNumCyclesForEachDz() {
	for q:=0; q<numDzs; q++ {
		numDzStates[q] = 5 //default
		//this is where we'd alter the number of disease states for each, e.g.
		numDzStates[1] = 8
		numCycles[q] = 10 //default
		//this is where we'd alter the number of disease cycles for each e.g.
		//numCycles[1] = 11
		fmt.Println("For disease number",q,"there are ",numDzStates[q],"number of disease states")
		fmt.Println("For disease number",q,"there are ",numCycles[q],"number of cycles")
		}
	fmt.Println("\n\n")
}


func initializeValuesOfTM() {
	//initializes the TM to all risk of progressions to be 1/numDzStates
	for q:=0; q<numDzs; q++ {
		for i := 0; i < numDzStates[q]; i++ {
			for j := 0; j < numDzStates[q]; j++ {						
				TM[q][i][j]=float32(1.0/float32(numDzStates[q]))
				}
			}
		}
//Here is where any values of TM for any disease q should be modified. The default setting is to put each to be 1/numDzStates.
}

func printAllTMs() {
	for q:=0; q<numDzs; q++ { 
		//Prints the TMs for all diseases
		fmt.Println("For disease ",q," the Transition Matrix is:")
		fmt.Println("--------------------------------------------------------------------------")		
		for i := 0; i< numDzStates[q]; i++ {
			for k :=0; k<numDzStates[q]; k++ {
				fmt.Print(TM[q][i][k],"\t")
				}
			fmt.Println("\n")
			}	
	}
}

func printTM(qf int) {
	//Prints the TMs for one disease qf
		fmt.Println("For disease ",qf," the Transition Matrix is:")
		fmt.Println("--------------------------------------------------------------------------")		
		for i := 0; i< numDzStates[qf]; i++ {
			for k :=0; k<numDzStates[qf]; k++ {
				fmt.Print(TM[qf][i][k],"\t")
				}
			fmt.Println("\n")
			}	
}


func assignValuesToAllAssociationArrays() {
	//AA = association array
	//AA[dz number ex TB where this data refers to TB's effects on other dz's][number of diseases - 
	//for each other disease that will be affected by TB - if this number is equal to the number in the
	//previous parameter, then all values have to be 1 because otherwise a disease would be affecting itself]
	//[number of disease states[previous parameter]][number of disease states[previous parameter]]
	for q:=0; q<numDzs; q++ {
		//for loops assigning all of AA to 1.0
		for affectedDzs :=0; affectedDzs < numDzs; affectedDzs++ {
			for fromDzState := 0; fromDzState < numDzStates[q]; fromDzState++ {
				for toDzState :=0; toDzState < numDzStates[q]; toDzState++ {
					AA[q][affectedDzs][fromDzState][toDzState] = 1.0
//					println("Association Array:",q,affectedDzs,fromDzState,toDzState,AA[q][affectedDzs][fromDzState][toDzState])
					}
				}
			}
		}
//Here is where any values of the Association Arrays for that are NOT 1 should be entered:

}

func calculateTMI() {
	//given TM and AA, this function calculates the TMI the Transition Matrix with Interactions. It does this for ALL diseases.
	for q:=0; q<numDzs; q++ {
		for affectedDzs :=0; affectedDzs < numDzs; affectedDzs++ {
			for fromDzState := 0; fromDzState < numDzStates[q]; fromDzState++ {
				for toDzState :=0; toDzState < numDzStates[q]; toDzState++ {
					TMI[q][fromDzState][toDzState]= AA[q][affectedDzs][fromDzState][toDzState] * TM[q][fromDzState][toDzState]
					//println(q,affectedDzs,fromDzState,toDzState,AA[q][affectedDzs][fromDzState][toDzState])
					}
				}
			}
		}
}

func printAllTMI() {
	for q:=0; q<numDzs; q++ {
			//Prints the TMI
		fmt.Println("For disease ",q," the Transition Matrix WITH INTERACTIONS factored in is:")
		fmt.Println("--------------------------------------------------------------------------")		
		for i := 0; i< numDzStates[q]; i++ {
			for k :=0; k<numDzStates[q]; k++ {
				fmt.Print(TMI[q][i][k],"\t")
				}
			fmt.Println("\n")
			}
		}
}

func printTMI(qf int) {
//Prints the TMI
	fmt.Println("For disease ",qf," the Transition Matrix WITH INTERACTIONS factored in is:")
	fmt.Println("--------------------------------------------------------------------------")		
	for i := 0; i< numDzStates[qf]; i++ {
		for k :=0; k<numDzStates[qf]; k++ {
			fmt.Print(TMI[qf][i][k],"\t")
			}
		fmt.Println("\n")
		}
}

func calculateSumTMI() {
	//calculates the horizontal and vertical TMI
	for q:=0; q<numDzs; q++ {
		//calculates the horizontal cumulative TMI
		for r:=0; r<numDzStates[q]; r++ {
			sumTMI[q][r][0]=TMI[q][r][0]		
			}
		for r:=0; r<numDzStates[q]; r++ {
			for c:=1; c<numDzStates[q]; c++ {
				sumTMI[q][r][c]=sumTMI[q][r][c-1]+TMI[q][r][c]
				}
			}
		//calculates the vertical cumulative TMI
		verticalSumTMI[q][0][0] = TMI[q][0][0]
		for r:=1; r<numDzStates[q]; r++ {
			verticalSumTMI[q][r][0]=verticalSumTMI[q][r-1][0]+TMI[q][r][0]
			}			
//		fmt.Println("The horizontal summative values of the TM WITH INTERACTIONS are:")
//		fmt.Println("--------------------------------------------------------------------------")		
//		for i := 0; i< numDzStates[q]; i++ {
//			for k :=0; k<numDzStates[q]; k++ {
//				fmt.Print(sumTMI[q][i][k],"\t")
//				}
//			fmt.Println("\n")
//			}
		}
}

func randomize() {
	//Random numbers. r is array of length number of cycles of psuedo-random numbers b/n 0 and 1 
	// whatever number goes inside of the Seed function will determine random #'s. can use time to randomize
	//this is inside the for loop for each disease, so each disease q has an array of random numbers
	//each corresponding to one of the cycles.  each disease has a different randome number for each cycle.
	rand.Seed(12)  
	fmt.Println("\nThe random numbers used in this model are:")
	for q:=0; q<numDzs; q++ {	
		for o :=0; o<numCycles[q]; o++ {
			//NEED to seed random number gen here. Otherwise each dz has the same set of random numbers. 
			random[q][o] = rand.Float32()
			}
		fmt.Println("\nDisease ",q)
		for o :=0; o<numCycles[q]; o++ {
			fmt.Print(random[q][o],"\t")	
			}
		}
}


func initializePeopleToStartStates() {
	for q:=0; q<numDzs; q++ {
	//puts the person into an appropriate start position, default is to place randomly proportional to 0th column of TM probabilities.
	//this initialization 
	if (random[q][0] <= verticalSumTMI[q][0][0]) {
		person1[q][0][0]=1
		state1[q][0]=1
		}
	for r:=1; r<numDzStates[q]; r++ {
		if ((random[q][0] > verticalSumTMI[q][r-1][0]) && (random[q][0] <= verticalSumTMI[q][r][0])) {
			person1[q][r][0]=1
			state1[q][0]=r
			}
		}
	}
//If you want to deliberately place a person in any disease state to start, here is where you would do so. 
}


func calculatePath() {
	//Determines person's path through dz states for every disease
	//state[qth dz][c-th cycle], will need to make this an extra dimension for more than 1 person
	fmt.Println("\n\nDetermining the individual's path...") //debugging
	for q:=0; q<numDzs; q++ {
		for cycle:=1; cycle<numCycles[q]; cycle++ {	
			//Determines if person belongs in TM[0][1] i.e. cycle 1, dz state 0. WORKS!!
			for i:=0; i<numDzStates[q]; i++ {
				if (person1[q][i][cycle-1]==1 && random[q][cycle]<=sumTMI[q][i][0]) {
						person1[q][0][cycle] = 1
						state1[q][cycle]=i
						}
				}
		//Determine if person belongs in TM[x][1] i.e. cycle 1, dz state 1 through numDzStates. WORKS!!
		for o:=1; o<numDzStates[q]; o++ {
			for j:=0; j<numDzStates[q]; j++ {
				if (person1[q][j][cycle-1]==1 && random[q][cycle]>sumTMI[q][j][o-1] && random[q][cycle]<=sumTMI[q][j][o]) {
					person1[q][o][cycle] = 1
					state1[q][cycle]=o
					}
				}
			}
		}
	}
}

func printPath(qf int) {
//Prints person's path in for disease qf in readable format
	fmt.Println("For disease",qf,", the individual's path is:")
	fmt.Println("\t\tCycle")
	fmt.Print("\t")
	for i := 0; i< numCycles[qf]; i++ {
		fmt.Print("\t",i)
		}
	fmt.Println("\n-------------------------------------------------------------------------------------------")	
	for i := 0; i< numDzStates[qf]; i++ {
		fmt.Print("DzState",i,"\t")
		for k :=0; k<numCycles[qf]; k++ {
			fmt.Print(person1[qf][i][k],"\t")
			}
		fmt.Println("\n")
		}
}

func printAllPaths() {
for q:=0; q<numDzs; q++ {
//Prints person's path in ALL diseases in readable format
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
	}
}

func printStateSeries(qf int) {
	fmt.Println("\nState Series:")
	fmt.Println("\nDisease ",qf)
	fmt.Println("Cycle")
	for c:=0; c<numCycles[qf]; c++ {
		fmt.Print(c,"\t")
		}
	fmt.Println("\n")	
	for c:=0; c<numCycles[qf]; c++ {
		fmt.Print(state1[qf][c],"\t")
		}
}

func printAllStateSeries() {
//Prints ALL state series
//fmt.Println("\n\n\n State Series is:",state1)
	fmt.Println("\nState Series:")
for q:=0; q<numDzs; q++ {
	fmt.Println("\n\nDisease ",q)
	fmt.Println("Cycle")
	for c:=0; c<numCycles[q]; c++ {
		fmt.Print(c,"\t")
		}
	fmt.Println("\n")	
	for c:=0; c<numCycles[q]; c++ {
		fmt.Print(state1[q][c],"\t")
		}
	}
}


/*
To Do:
1. random num generator and put into for loop
2. Change array from being set up in program to being imported from csv
3. Mapping names to dz states
4. input for the variables
5. utility associated with each disease state. for each disease (q of them), there will have to be a u value for each disease state.  each person should have that utility
//so maybe make the person1 has a variable that corresponds to utility. no that's inefficient.  if person[q].  no we could have it mirror the state series.  
// or an array u which the value is the utility [q = diseases][numDzStates]. then if utility[q][[stateSeries[q][numCycles[q]]]] will be the utility for that person.
//do we save that utility as a parameter of the person?  yes, probably so? if we can make one dimension of that array a float with the others being integers? 
//if not, then make a separate utility variable
//I think this is best because it allows person to just be the path and the data to be kept in stateSeries and utility matrices.  
6. graphing
7. more than 1 person. just treat as array of people.
8. interventions should have an AA but not a fully fledged disease.  there should be a function which says 
// apply intervention!!
9. function to average the utilities at the end

person1 represents a person going through the world with susceptibility to various diseases. it is a instance of
the variable individualpath.  so eventually person[] will be an array of individualpaths

every other variable works for a given disease.  should we just make ever variable a vector corresponding to the
number disease that it is? that would probably be the simpliest. so the variables are 
numDzStates, numCycles, TM, and variations off of TM - sumTM both horizontal and vertical. 
*/
     


