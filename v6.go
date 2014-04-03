package main

import (
	"fmt"
)



func main() {

	
	const numberDzStates = 5
	var TM [numberDzStates]float64
	
	
	
	for i := 1; i < len(TM); i++ {
		TM[i]=1
		}
	fmt.Println(TM)

}