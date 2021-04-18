package main

import (
	"fmt"
	"github.com/kezhuw/treedb"
	"os"
	"reflect"
	"time"
	//"github.com/kezhuw/treedb/cmd/treedb/cmds"
	//"github.com/peterh/liner"
	"bytes"
	"encoding/gob"
)

type Post struct {
	Id    int
	Name  string
	Text  string
	Posts []Post
}

func P(t interface{}) {
	fmt.Println(reflect.TypeOf(t))
}

func main() {

	client, err := treedb.Dial("tcp4://:3456", nil)
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(0)
	}

	db, err := client.OpenDB("gusadb", nil)
	if err != nil {
		os.Exit(0)
	}

	var posts []Post
	path := "top/hyogo/tree"

	posts = append(posts, Post{Id: 0, Name: "a", Text: "b"})
	posts[0].Posts = append(posts[0].Posts, Post{Id: 1, Name: "c", Text: "d"})

	posts = append(posts, Post{Id: 2, Name: "e", Text: "f"})
	posts[1].Posts = append(posts[1].Posts, Post{Id: 3, Name: "h", Text: "d"})
	fmt.Printf("%v\n", posts)

	buf := bytes.NewBuffer(nil)
	_ = gob.NewEncoder(buf).Encode(posts)

	db.Set(path, buf.Bytes())

	time.Sleep(time.Second * 1)

	b := make([]Post, 10)
	value, err := db.Get(path, treedb.FieldAny)
	if err != nil {
		fmt.Printf("ERROR %v\n", err)
		os.Exit(0)
	}

	switch value := value.(type) {
	case []byte:
		buf := bytes.NewBuffer(value)
		_ = gob.NewDecoder(buf).Decode(&b)
		fmt.Printf("%v\n", b)
	case map[string]interface{}:
		fmt.Printf("%v\n", value)
	default:
		fmt.Printf("unexpected value type: %s", reflect.TypeOf(value))

	}

}
