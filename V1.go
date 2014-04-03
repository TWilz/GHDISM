package main


import "fmt"
import "encoding/csv"
import "io"
import "os"


/*
	imports excel spreadsheet, 
	first sheet is TM
	first version is just for hiv w/ predefined states
	
*/
 


func main() {

	csvFile, err := os.Open("TMv1.csv")
	defer csvFile.Close()
	
	var input float64
	fmt.Scanf("%f", &input)
	fmt.Println(input)
}