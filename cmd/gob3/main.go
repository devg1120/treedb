package main

import (
    "bytes"
    "time"
    "fmt"
    "encoding/gob"
)

// サンプル用の構造体
// 何となくブログ風味
type Entry struct {
    ID         int64 `datastore:"-"`
    Title      string
    Contents   string
    CreateAt   time.Time
    UpdateAt   time.Time
    DeleteAt   time.Time
}

// シリアライズ用のメソッド
// レシーバ(e)の値をシリアライズしてbyte配列にする
func (e *Entry) GobEncode() ([]byte, error) {
    // エンコーダを作る.
    // io.Writerであれば何でも良いのでここではbytes.Bufferを使う
    w := new(bytes.Buffer)
    encoder := gob.NewEncoder(w)

    // 先頭から順番にエンコードしていく
    // 必ず先頭から1つずつエンコードを実行していく必要があるらしい

    if err := encoder.Encode(e.ID); err != nil {
        return nil, err
    }
    if err := encoder.Encode(e.Title); err != nil {
        return nil, err
    }
    if err := encoder.Encode(e.Contents); err != nil {
        return nil, err
    }
    if err := encoder.Encode(e.CreateAt); err != nil {
        return nil, err
    }
    if err := encoder.Encode(e.UpdateAt); err != nil {
        return nil, err
    }
    if err := encoder.Encode(e.DeleteAt); err != nil {
        return nil, err
    }
    return w.Bytes(), nil
}

// デシリアライズするためのデコーダ
// byte配列を渡してレシーバへ値を入れる
func (e *Entry) GobDecode(buf []byte) error {
    // デコーダを作る.
    // io.Readerであれば何でも良いのでここではbytes.Bufferを使う
    r := bytes.NewBuffer(buf)
    decoder := gob.NewDecoder(r)

    // 先頭から順番にデコードしていく
    // 必ず先頭から1つずつデコードを実行していく必要があるらしい

    if err := decoder.Decode(&e.ID); err != nil {
        return err
    }
    if err := decoder.Decode(&e.Title); err != nil {
        return err
    }
    if err := decoder.Decode(&e.Contents); err != nil {
        return err
    }
    if err := decoder.Decode(&e.CreateAt); err != nil {
        return err
    }
    if err := decoder.Decode(&e.UpdateAt); err != nil {
        return err
    }
    if err := decoder.Decode(&e.DeleteAt); err != nil {
        return err
    }
    return nil
}

func main() {
    // 記録するデータ
    e1 := &Entry{
        ID: 1,
        Title: "title",
        Contents: "contents",
        CreateAt: time.Now(),
        UpdateAt: time.Now(),
        DeleteAt: time.Now(),
    }
    // bにシリアライズして入れる
    b, err := e1.GobEncode()
    if err != nil {
        fmt.Println("Error occured.", err)
    }

    // デシリアライズするために用意
    e2 := &Entry{}

    // bからe2へデシリアライズ
    if err = e2.GobDecode(b); err != nil {
        fmt.Println("Error occured.", err)
    }

    // デシリアライズしたものを表示
    fmt.Println(e2.ID)
    fmt.Println(e2.Title)
    fmt.Println(e2.Contents)
    fmt.Println(e2.CreateAt)
    fmt.Println(e2.UpdateAt)
    fmt.Println(e2.DeleteAt)
}
