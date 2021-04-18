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

/*
type Post struct {
	Id    int
	Name  string
	Text  string
	Posts []Post
}
*/

type Node struct {
	Id    int
	Name  string
	Text  string
        Left   []Node
        Right  []Node
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


	var nodes []Node
	path := "top/hyogo/tree"

	nodes = append(nodes, Node{Id: 0, Name: "a", Text: "b"})
	nodes[0].Left = append(nodes[0].Left, Node{Id: 1, Name: "c", Text: "d"})

	nodes[0].Right = append(nodes[0].Right, Node{Id: 3, Name: "h", Text: "d"})

	nodes[0].Right[0].Left = append(nodes[0].Right[0].Left, Node{Id: 5, Name: "s", Text: "S"})
	nodes[0].Right[0].Right = append(nodes[0].Right[0].Right, Node{Id: 4, Name: "x", Text: "z"})

	fmt.Printf("%v\n", nodes)

	buf := bytes.NewBuffer(nil)
	_ = gob.NewEncoder(buf).Encode(nodes)


/*
	var top Node
	path := "top/hyogo/tree"

	top =  Node{Id: 0, Name: "a", Text: "b"}
	top.Left =  Node{Id: 1, Name: "c", Text: "d"}

	top.Right =  Node{Id: 3, Name: "h", Text: "d"}
	fmt.Printf("%v\n", top)

	buf := bytes.NewBuffer(nil)
	_ = gob.NewEncoder(buf).Encode(top)
*/


	db.Set(path, buf.Bytes())

	time.Sleep(time.Second * 1)

        var b = make([]Node,10)
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
