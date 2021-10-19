package main


import (
    "crypto/md5"
    "fmt"
    "reflect"
    "encoding/hex"
)

func main() {
    i := 609043
    var s string
    
    lead := []byte{0,0,0,0,0}
    for {
        s = fmt.Sprintf("%06d", i)

        data := []byte("abcdef"+s)
        hash := hex.EncodeToString(md5.Sum(data))
        fmt.Println(hash, "abcdef"+s)
        //if reflect.DeepEqual(hash[:5],lead) {
         //   fmt.Println(i)
         //   break
        //}
        break
        i++
    }
}
