package main

import (
    "fmt"
    "os"
)


func main() {
    dat, err := os.ReadFile("data")
    check(err)
    count := 0
    firstNeg := 0
    for i, d := range dat {
        if d == '(' {
            count++
        } else if d == ')' {
            count--
        }

        if count == -1 && firstNeg == 0{
            firstNeg = i + 1
        }
    }

    fmt.Println(count, firstNeg)

}

func check(e error) {
    if e != nil {
        panic(e)
    }
}
