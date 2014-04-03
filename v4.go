package main

import (
    "encoding/csv"
    "os"
    "fmt"
)


func main() {
    filename, err := os.Open("TMv1.csv")
    	if err != nil { 
		panic(err) 
	}
    	defer filename.Close()
	r := csv.NewReader(filename)
    fmt.Println(r)
}