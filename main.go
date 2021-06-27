package main

import (
	"fmt"
	"math/rand"
)

var (
	distance = 150 //km
    company, trip string
    i = 0
	)

func main ()  {
	var pi rune = 960
var alpha rune = 940
var omega rune = 969
var bang byte = 33
	
	fmt.Println("Company        Hours Trip_type  Price")
    fmt.Println("======================================")

    for ; i < 5; i++ {
        switch rand.Intn(3) {
        case 0:
            company = "YandexGo"
        case 1:
            company = "Ситимобил"
        case 2:
            company = "Uber"
        }
 	
 	speed := 60 + rand.Intn(60)                 // 60<n<80 км/ч
        hours := distance / speed 				
        price := 1000.0 + speed                        
 
        if rand.Intn(2) == 1 {
            trip = "Round-trip"
            price = price * 2
        } else {
            trip = "One-way"
        }
 
        fmt.Printf("%-14v %5v %-10v ₽%4v %4v\n", company, hours, trip, price)
    }
    
}