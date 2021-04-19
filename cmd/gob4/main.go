package main

import (
    "bufio"
    "encoding/gob"
    "fmt"
    "os"
)


func main() {


    maps := map[string]int{"apple": 150, "banana": 300, "lemon": 300}

    fmt.Printf("%v\n", maps)

    path := "map.gob"

    // write
    out, err1 := os.Create(path)
    if err1 != nil {
        fmt.Printf("File write error: %v\n", err1)
        os.Exit(1)
    }
    w := bufio.NewWriter(out)
    enc := gob.NewEncoder(w)
    enc.Encode(maps)
    w.Flush()
    out.Close()

    // read
    b := make(map[string]int, 0)
    in, err2 := os.Open(path)
    if err2 != nil {
        fmt.Printf("File read error: %v\n", err2)
        os.Exit(1)
    }
    r := bufio.NewReader(in)
    dec := gob.NewDecoder(r)
    dec.Decode(&b)

    fmt.Printf("%v\n", b)

}
