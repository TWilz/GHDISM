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
 //GHDISM 8,9 broken. doesn't work
 //GHDISM 10 reverted back to v7. set up where everyone starts in dz state 0. 
 //GHDISM11 change intervention to be different from a dz. works!
 //GHDISM12 adding in any number of people. Works for any number of people. Tried with 20 dzs, 10 interventions, 100 people, 200 cycles and it worked. 
 //GHDISM13 adding in utilities. works except for rand(10) w/ 2 dzs, 1 intervention, 20 max cyces, 2 people. 10th cycle of person1,dz0 is discordant in stateseries and path of 1s and 0s
 //GHDISM14 add in input from user
 
package main

import (
	"fmt"
	"math/rand"
//	"code.google.com/p/plotinum/plot"
//  "code.google.com/p/plotinum/plotter"
//    "code.google.com/p/plotinum/plotutil"
	)

//Declaration of global variables
	const numDiseases = 2
	const numInterventions = 1
	const numDzs = numDiseases 
	const maxNumDzStates = 10
	const maxNumCycles = 20
	const numberOfPeople = 2 //1 being 1st and 0th person
	type individualpath [numberOfPeople][numDzs][maxNumDzStates][maxNumCycles]int
	var person individualpath
	type stateSeries [numberOfPeople][numDzs][maxNumCycles]int
	var state stateSeries
	var numDzStates [numDzs]int
	var numCycles [numDzs]int
	type matrix [numDzs][maxNumDzStates][maxNumDzStates]float32  
	// A transition matrix which is actually an 3D array
	//where [i][j] is i-th row and j-th column
	var TM matrix
	type interactionData [numDzs][numDzs][maxNumDzStates][maxNumDzStates]float32
	var AA interactionData
	type interventionData [numInterventions][numDzs][maxNumDzStates][maxNumDzStates]float32
	var intervention interventionData
	//an intervention can affect the TM values of every dz. the interventions won't be able to affect other interventions.
	
	var q int //q is used to denote the disease number
	//local function variables take the name of the global variable but with an f appended on e.g. q becomes qf when used locally
	var TMI matrix //TMI = transitions matrix with interactions =  TM x corresponding AA value. TMI and not TM should be used when computing a person's path
	var sumTMI matrix //matrix of TMI where each row is horizontally summed i.e. row 1 = row 0 + row 1, etc. this is used for backend calculations and will likely not need to be outputted.
	var verticalSumTMI matrix
	type randomMatrix [numberOfPeople][numDzs][maxNumCycles]float32
	var random randomMatrix
	
	type utilityMatrix [numDzs][maxNumDzStates]float32
	var u utilityMatrix
	type qalyMatrix[numberOfPeople][maxNumCycles]float32
	var qaly qalyMatrix
	
func printSizeOfModelStructure() {
	fmt.Println("\n\n")
	fmt.Println("There are ",numberOfPeople," in this model.")
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

func calculateTMIforDzInteractions() {
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
	rand.Seed(100)  //NEED to seed random number gen here. Otherwise each dz has the same set of random numbers. 
	for per:=0;per<numberOfPeople;per++{
		fmt.Println("\nThe random numbers for person",per," are")
		for q:=0; q<numDzs; q++ {	
			for o :=0; o<numCycles[q]; o++ {
				random[per][q][o] = rand.Float32()
				}
			
			for o :=0; o<numCycles[q]; o++ {
				fmt.Print(random[per][q][o],"\t")	
				}
			}
		}

}


func initializePeopleToStartStates() {
	//puts the person into an appropriate start position, default is to place randomly proportional to 0th column of TM probabilities. 
	//this initialization doesn't make sense with how we've set up our TM.  
	//Let's use a separate sheet in the csv's to store data to indicate prevalence. 
	//For now, let's make everyone start off healthy	
for per:=0;per<numberOfPeople;per++{
	for q:=0; q<numDzs; q++ {
		//initializes everyone to start in dz state 0
		person[per][q][0][0]=1
		state[per][q][0]=1
		/* This next section of code somewhat randomly puts a person into the disease states.  It is based off of the
			the 0th column of the transition matrix which doesn't make intuitive sense for the TM's design
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
		*/
		}
//If you want to deliberately place a person in any disease state to start, here is where you would do so. 
//likely import from a csv page depending on the prevalence of the disease states. 
}
}


func calculatePath() {
	//Determines person's path through dz states for every disease
	//state[qth dz][c-th cycle], will need to make this an extra dimension for more than 1 person
	fmt.Println("\n\nDetermining the paths...") //debugging
for per:=0;per<numberOfPeople;per++{
	for q:=0; q<numDzs; q++ {
		for cycle:=1; cycle<numCycles[q]; cycle++ {	
			//Determines if person belongs in TM[0][1] i.e. cycle 1, dz state 0. WORKS!!
			for i:=0; i<numDzStates[q]; i++ {
				if (person[per][q][i][cycle-1]==1 && random[per][q][cycle]<=sumTMI[q][i][0]) {
						person[per][q][0][cycle] = 1
						state[per][q][cycle]=i
						}
				}
			//Determine if person belongs in TM[x][1] i.e. cycle 1, dz state 1 through numDzStates. WORKS!!
			for o:=1; o<numDzStates[q]; o++ {
				for j:=0; j<numDzStates[q]; j++ {
					if (person[per][q][j][cycle-1]==1 && random[per][q][cycle]>sumTMI[q][j][o-1] && random[per][q][cycle]<=sumTMI[q][j][o]) {
						person[per][q][o][cycle] = 1
						state[per][q][cycle]=o
						}
					}
				}
			}
		}
}
}

func printPath(qf int, per int) {
//Prints person's path in for disease qf in readable format
	fmt.Println("For person ",per," disease ",qf,", the individual's path is:")
	fmt.Println("\t\tCycle")
	fmt.Print("\t")
	for i := 0; i< numCycles[qf]; i++ {
		fmt.Print("\t",i)
		}
	fmt.Println("\n-------------------------------------------------------------------------------------------")	
	for i := 0; i< numDzStates[qf]; i++ {
		fmt.Print("DzState",i,"\t")
		for k :=0; k<numCycles[qf]; k++ {
			fmt.Print(person[per][qf][i][k],"\t")
			}
		fmt.Println("\n")
		}
}

func printAllPaths() {
for per:=0;per<numberOfPeople;per++{
for q:=0; q<numDzs; q++ {
//Prints person's path in ALL diseases in readable format
	fmt.Println("For person ",per," disease ",q,", the individual's path is:")
	fmt.Println("\t\tCycle")
	fmt.Print("\t")
	for i := 0; i< numCycles[q]; i++ {
		fmt.Print("\t",i)
		}
	fmt.Println("\n-------------------------------------------------------------------------------------------")	
	for i := 0; i< numDzStates[q]; i++ {
		fmt.Print("DzState",i,"\t")
		for k :=0; k<numCycles[q]; k++ {
			fmt.Print(person[per][q][i][k],"\t")
			}
		fmt.Println("\n")
		}
	}
}
}

func printStateSeries(qf int, per int) {
	fmt.Println("\n\nPerson",per,"Disease ",qf)
	fmt.Println("Cycle")
	for c:=0; c<numCycles[qf]; c++ {
		fmt.Print(c,"\t")
		}
	fmt.Println("\n")	
	for c:=0; c<numCycles[qf]; c++ {
		fmt.Print(state[per][qf][c],"\t")
		}
}

func printAllStateSeries() {
//Prints ALL state series for ALL people
//fmt.Println("\n\n\n State Series is:",state1)
for per:=0;per<numberOfPeople;per++{
for q:=0; q<numDzs; q++ {
	fmt.Println("\n\nPerson ",per," Disease ",q)
	fmt.Println("Cycle")
	for c:=0; c<numCycles[q]; c++ {
		fmt.Print(c,"\t")
		}
	fmt.Println("\n")	
	for c:=0; c<numCycles[q]; c++ {
		fmt.Print(state[per][q][c],"\t")
		}
	}
}
}

func assignValuesToAllInterventionArrays() {
//assigns all intervention coefficients to 1. This is where we'd input data from intervention csv's into intervention array
	for i:=0;i<numInterventions;i++{
		for d:=0;d<numDzs;d++{
			for fromDzState:=0;fromDzState<numDzStates[d];fromDzState++ {
				for toDzState:=0;toDzState<numDzStates[d];toDzState++{
					intervention[i][d][fromDzState][toDzState]=1.0
					}
				}
			}
		
		}
//Here is where any changes to the intervention coefficients could be made
//intervention[0][0][0][1]=100.0 //a test to make sure the correct intervention is tied to the correct value in the TMI
}

func calculateTMIforAllInterventions() {
	//Given intervention arrays, it calculates the interventions' effects on the TMI.  Note that interactions b/n dz's must be run first
	for i:=0;i<numInterventions;i++{
		for d:=0;d<numDzs;d++{
			for fromDzState:=0;fromDzState<numDzStates[d];fromDzState++ {
				for toDzState:=0;toDzState<numDzStates[d];toDzState++{
					TMI[d][fromDzState][toDzState]=TMI[d][fromDzState][toDzState]*intervention[i][d][fromDzState][toDzState]
					}
				}
			}
		
		}	
}

func assignUtilityMatrix() {
//	assigns utilities to all values of the utility matrix. 
//	this is where we would import utility data from csv 
//the default here is to make healthy utility 1 and each dz state beneath that to be a fixed amount (proportional to number of dz states for that dz) lower 
//	type utilityMatrix [numDzs][maxNumDzStates]float32
	for dz:=0;dz<numDzs;dz++{
		u[dz][0]=1.0
		for dzState:=1;dzState<numDzStates[dz];dzState++{
			u[dz][dzState]=float32(u[dz][dzState-1] - (1.0 / float32((int(numDzStates[dz]) - 1))))
			}
		}
}

func printAllUtilityMatrix() {
	fmt.Println("\nThe Utility Matrices are:")
	for dz:=0;dz<numDzs;dz++{
		fmt.Println("\nDisease ",dz,"\n")
		for dzState:=1;dzState<numDzStates[dz];dzState++{
			fmt.Println("State:",dzState,", Utility:",u[dz][dzState],"\n")
			}
		}
}


func printUtilities() {
//type stateSeries [numberOfPeople][numDzs][maxNumCycles]int
	fmt.Println("\n\nAfter modeling the diseases and interventions, the utilities for each cycle of the people are:")
	for per:=0;per<numberOfPeople;per++ {
		fmt.Println("\nFor person",per,":")
			for dz:=0;dz<numDzs;dz++{
//				fmt.Println("per:",per)//db
//				fmt.Println("dz:",dz)//db
//				fmt.Println("numCycles[dz]:",numCycles[dz])//db
//				fmt.Println("state[per][dz][numCycles[dz]:",state[per][dz][numCycles[dz]])//db
//				fmt.Println("state[0][0][9]:",state[0][0][9])
				for cycle:=0;cycle<numCycles[dz];cycle++{
					fmt.Println("Disease ",dz," cycle ",cycle," the utility:",u[dz][state[per][dz][cycle]])
				}
				}
			}

}


func calculateQALYs() {
//
//type qaly[numberOfPeople][maxNumCycles]
//type utilityMatrix [numDzs][maxNumDzStates]float32
	for per:=0;per<numberOfPeople;per++ {
  		for dz:=0;dz<numDzs;dz++ {
  			for cycle:=1;cycle<numCycles[dz];cycle++ {
  				qaly[per][cycle]= qaly[per][cycle-1] + (u[dz][state[per][dz][cycle]] * 1) 
				//all initial qaly's i.e. qaly[0] are 0.
				//times 1 b/c assume 1 cycle= 1 year. 
				//this could be changed *(1/12) if each cycle was a month rather than a year.
				}
			}
		}
	fmt.Println("QALYs:",qaly)
}

func main() {
//	var dzOfInterest int
//	var interventionOfInterest int
//var personOfInterest int
//	dzOfInterest = 0
//	interventionOfInterest = 0		
// personOfInterest =0
	printSizeOfModelStructure()

//Initializes model	
	defineNumDzStatesAndNumCyclesForEachDz()
	initializeValuesOfTM()
	assignValuesToAllAssociationArrays()
	assignValuesToAllInterventionArrays()
	assignUtilityMatrix()
//This would be the place to alter any values


//Computes back end numbers for model
	calculateTMIforDzInteractions()
	calculateTMIforAllInterventions() //i.e. this applies interventions. Note: calculatingTMI for Dz Interactions must be done before TMI calc for interventions.
//	calculateTMIforIntervention(interventionOfInterest) //this applies a single intervention and calculates the TMI
	calculateSumTMI()
	randomize()

//Calculates the path of the individuals
	initializePeopleToStartStates()
	calculatePath()

	
//Request user input on what interested in

//All calculations are done by here. Now it's just output. 
//	printTM(dzOfInterest)
	printAllTMs()
//	printTMI(dzOfInterest)
	printAllTMI()
	printAllUtilityMatrix() 
//	printPath(dzOfInterest, personOfInterest)
	printAllPaths()
//	printStateSeries(dzOfInterest, personOfInterest)
	printAllStateSeries()

	printUtilities()
	fmt.Println("\n")
	calculateQALYs()
}//end main

/*
To Do:
1. random num generator and put into for loop
2. Change array from being set up in program to being imported from csv
3. Mapping names to dz states
4. input for the variables
	what would you like to do? then give a list of options - have it in a do loop with last option being exit
5. utility associated with each disease state
6. graphing
7. more than 1 person
8. need to put in ability to apply an individual intervention arbitrary number of times. 

assign utlities, put in user interface, make model for more than 1 person
more than 1 person - need to also add dimension to stateseries and randomize for each person

add in utilities

person1 represents a person going through the world with susceptibility to various diseases. it is a instance of
the variable individualpath.  so eventually person[] will be an array of individualpaths

every other variable works for a given disease.  should we just make ever variable a vector corresponding to the
number disease that it is? that would probably be the simpliest. so the variables are 
numDzStates, numCycles, TM, and variations off of TM - sumTM both horizontal and vertical. 

Questions: 
Alex - can we have more than one sheet in a csv? I think not.
	random number generator
	graphing
Jim - How to set up csv so that way it conveys info for a disease and one that conveys info for an intervention. 
	see example.  also how should we indicate initial prevalences and utilities?
Bug: with random(10), check per1,dz0,10th(#9) cycle. pathprinting and state series are discordant. 
*/
     


