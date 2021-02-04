package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    rand.Seed(time.Now().Unix())

   
    rangeLower := 2
    rangeUpper := 50
    randomNum := rangeLower + rand.Intn(rangeUpper-rangeLower+1)
    fmt.Println(randomNum)

    
}
