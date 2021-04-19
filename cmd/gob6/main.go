package main

import (
    "bufio"
    "encoding/gob"
    "fmt"
    "os"
)

type MyFace interface {
    A()
}

type Cat struct{}
type Dog struct{}

func (c *Cat) A() {
    fmt.Println("Meow")
}

func (d *Dog) A() {
    fmt.Println("Woof")
}


func main() {

    gob.Register(&Cat{})
    gob.Register(&Dog{})




    path := "interface.gob"

    // write
    out, err1 := os.Create(path)
    if err1 != nil {
        fmt.Printf("File write error: %v\n", err1)
        os.Exit(1)
    }
    w := bufio.NewWriter(out)
    enc := gob.NewEncoder(w)
    //enc.Encode(test)

    var inter MyFace
    inter = new(Cat)

    // Note: pointer to the interface
    err := enc.Encode(&inter)
    if err != nil {
        panic(err)
    }

    inter = new(Dog)
    err = enc.Encode(&inter)
    if err != nil {
        panic(err)
    }


    w.Flush()
    out.Close()

    // read
    var get MyFace
    in, err2 := os.Open(path)
    if err2 != nil {
        fmt.Printf("File read error: %v\n", err2)
        os.Exit(1)
    }
    r := bufio.NewReader(in)
    dec := gob.NewDecoder(r)
    dec.Decode(&get)

    get.A()

    dec.Decode(&get)

    get.A()

}
