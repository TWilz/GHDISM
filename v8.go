package main

import (
	"fmt"
)



func main() {


	const numDzStates = 5
	

	type matrix [numDzStates][numDzStates]float64  // A 3x3 array, really an array of arrays.


	var TM matrix

	for i := 1; i < numDzStates; i++ {
		for j :=1; j < numDzStates; j++ {			
			TM[i][j]=1
			}
		}


	fmt.Println(TM[1][1])


    
}