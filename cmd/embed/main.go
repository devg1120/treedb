package main

import (
	"os"
	"fmt"
        "reflect"
	"github.com/kezhuw/treedb"
	//"github.com/kezhuw/treedb/cmd/treedb/cmds"
	//"github.com/peterh/liner"
        "bytes"
        "encoding/gob"
)

type Hoge struct {
      F1 string
      F2 int64
}


func encode() []byte {
      h := Hoge{F1: "hoge", F2: 123}
      buf := bytes.NewBuffer(nil)
      _ = gob.NewEncoder(buf).Encode(&h)
      return buf.Bytes()
}

func decode(data []byte) *Hoge {
      var h Hoge
      buf := bytes.NewBuffer(data)
      _ = gob.NewDecoder(buf).Decode(&h)
      return &h
}

func P(t interface{}) {
    fmt.Println(reflect.TypeOf(t))
}


func get(db *treedb.DB, path string) {

  	value, err := db.Get(path, treedb.FieldAny)
  	//value, err := db.Get("top/hyogo", treedb.FieldAny)
	if err != nil {
             os.Exit(0)
	}

        P(value)
//        printValue(value)

        fmt.Printf("%v\n", value)

       	switch value := value.(type) {
	case []byte:
		fmt.Printf("%s\n",string(value))
	case map[string]interface{}:
		fmt.Printf("%v\n",value)
	default:
		fmt.Printf("unexpected value type: %s", reflect.TypeOf(value))
	}

}    

func get_decode(db *treedb.DB, path string) {

  	value, err := db.Get(path, treedb.FieldAny)
  	//value, err := db.Get("top/hyogo", treedb.FieldAny)
	if err != nil {
             fmt.Printf("ERROR\n")
             os.Exit(0)
	}

        P(value)
//        printValue(value)
       	switch value := value.(type) {
	case []byte:
                decoded := decode(value)
                fmt.Printf("decoded: %+v\n", decoded)
                fmt.Printf("%s %d\n", decoded.F1, decoded.F2)
	case map[string]interface{}:
		fmt.Printf("%v\n",value)
	default:
		fmt.Printf("unexpected value type: %s", reflect.TypeOf(value))
	}



}    

func main() {

  //_, _ = treedb.Dial("tcp4://:3456", nil)
  client, err := treedb.Dial("tcp4://:3456", nil)
  //_, err := treedb.Dial("tcp4://:3456", nil)
  if err != nil {
	fmt.Printf("%v\n",err)
        os.Exit(0)
  }

  db, err := client.OpenDB("gusadb", nil)
  if err != nil {
        os.Exit(0)
  }

  	get(db,"top/hyogo/kobe")

  	get(db,"top/hyogo")

  	db.Set("top/hyogo/miki", []byte("OK"))
  	get(db,"top/hyogo/miki")




    encoded := encode()
    fmt.Printf("encoded: %d bytes\n", len(encoded))

   // decoded := decode(encoded)
   // fmt.Printf("decoded: %+v\n", decoded)

    db.Set("top/hyogo/enco", encoded)

    get_decode(db,"top/hyogo/enco")
    //get(db,"top/hyogo/enco")


}
