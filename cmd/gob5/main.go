package main

import (
    "bufio"
    "encoding/gob"
    "fmt"
    "os"
)

//type Post struct {
//    Id    int
//    Name  string
//    Text  string
//    Posts []Post
//}


type Test struct {
  Value interface{}
}

func main() {

//    var posts []Post
//
//    posts = append(posts, Post{Id: 0, Name: "a", Text: "b"})
//    posts[0].Posts = append(posts[0].Posts, Post{Id: 1, Name: "c", Text: "d"})
//
//    posts = append(posts, Post{Id: 2, Name: "e", Text: "f"})
//    posts[0].Posts = append(posts[0].Posts, Post{Id: 3, Name: "h", Text: "d"})
//    fmt.Printf("%v\n", posts)
//
//    path := "post.gob"

    gob.Register(Value interface{})

    var test Test
    test.Value = "test"
    fmt.Printf("%v\n", test)
    path := "interface.gob"

    // write
    out, err1 := os.Create(path)
    if err1 != nil {
        fmt.Printf("File write error: %v\n", err1)
        os.Exit(1)
    }
    w := bufio.NewWriter(out)
    enc := gob.NewEncoder(w)
    enc.Encode(test)
    w.Flush()
    out.Close()

    // read
    b := make([]Test, 10)
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
