package main

import (
	"fmt"
)
/*
func getData   {
	var x[]float64

	const numberDzStates = 5
	x := make([]float64, numberDzStates)
	for _, v := range x {
		total += v
		}
	

}

*/


func main() {

	
	const numberDzStates = 5
	TM := make([]float64, numberDzStates)
	fmt.Println(TM)
	
	for i := 1; i < len(TM); i++ {
		TM[i]=1
		}

}