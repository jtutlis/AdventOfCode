package main

import (
    "fmt"
    "bufio"
    "os"
    "strings"
    "strconv"
)


func main() {
    file, _ := os.Open("data")
    defer file.Close()

    scanner := bufio.NewScanner(file)
    totalArea,totalRibbon := 0,0
    for scanner.Scan() {
        totalArea += getSurfaceArea(scanner.Text())
        totalRibbon += getRibbon(scanner.Text())
    }
    fmt.Println(totalArea, totalRibbon)
}

func getSurfaceArea(s string) int {
    v := strings.Split(s, "x")
    l,_ := strconv.Atoi(v[0])
    w,_ := strconv.Atoi(v[1])
    h,_ := strconv.Atoi(v[2])


    return 2*l*w + 2*w*h + 2*h*l + MinOf(l*w,w*h,h*l)
}

func getRibbon(s string) int {
    v := strings.Split(s, "x")
    l,_ := strconv.Atoi(v[0])
    w,_ := strconv.Atoi(v[1])
    h,_ := strconv.Atoi(v[2])
    // 20 iq code
    return l*w*h + MinOf(l*2+w*2, w*2+h*2, l*2+h*2)

    

}
func MinOf(vars ...int) int {
    min := vars[0]

    for _, i := range vars {
        if min > i {
            min = i
        }
    }

    return min
}
