package main
import "os"
import "github.com/bmatsuo/csvutil"

type Person struct {
    Name   string
    Height float64
    Weight float64
}

func main() {
    writer := csvutil.NewWriter(os.Stdout, nil)
    errdo := csvutil.DoFile(os.Args[1], func(r csvutil.Row) bool {
        if r.HasError() {
            panic(r.Error)
        }
        var person Person
        if _, errc := r.Format(&person); errc != nil {
            panic("Row is not a Person")
        }
        bmi := person.Weight / (person.Height * person.Height)
        writer.WriteRow(csvutil.FormatRow(person, bmi).Fields...)
        return true
    })
    if errdo != nil {
        panic(errdo)
    }
    writer.Flush()
}