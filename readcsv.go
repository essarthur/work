package main

import (
    "encoding/csv"
    "fmt"
    "os"
)

func main() {

    file, err := os.Open("Account_balances.csv")
    if err != nil {
        fmt.Println("Error", err)
        return
    }
    defer file.Close()

    reader := csv.NewReader(file)
    record, err := reader.ReadAll()
    if err != nil {
        fmt.Println("Error", err)
    }

    for value:= range record{ // for i:=0; i<len(record)
        fmt.Println("", record[value])
    }
}
