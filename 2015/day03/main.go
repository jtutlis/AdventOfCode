package main

import (
    "fmt"
    "os"
)

type Point struct {
    X,Y int
}

func main() {
    x,y:=0,0
    q,w:=0,0
    m := make(map[Point]struct{})


    data,_ := os.ReadFile("data")
    var robo bool
    for _,d := range data {
        if robo {
             switch d {
                case '^':
                    q++
                case 'v':
                    q--
                case '>':
                    w++
                case '<':
                    w--
            }
        } else {
            switch d {
                case '^':
                    y++
                case 'v':
                    y--
                case '>':
                    x++
                case '<':
                    x--
            }
        }
        m[Point{x,y}] = struct{}{}
        m[Point{w,q}] = struct{}{}
        robo = !robo
    }
    fmt.Println(len(m))



}
