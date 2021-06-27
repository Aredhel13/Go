package main

import (
    "fmt"
)

func main ()  {

var (
    pi rune = 960
    alpha rune = 940
    omega rune = 969
    bang byte = 33
    )
message := "uv vagreangvbany fcnpr fgngvba"

for i := 0; i < len(message); i++ {                 
    c := message[i]
    if c >= 'a' && c <= 'z' {           
        c = c + 13
        if c > 'z' {
            c = c - 26
        }
    }
    fmt.Printf("%c", c)
}
fmt.Printf("\n%c%c%c%c\n", pi, alpha, omega, bang) // Выводит: πάω!0
}