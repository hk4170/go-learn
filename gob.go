package main

import (
    "encoding/gob"
    "fmt"
    "log"
    "os"
)

type Person struct {
    Name    string
    Age     int
    Email   string
    Hobbies []string
}

func main() {
    // 创建一个 Person 对象
    p := Person{
        Name:    "Alice",
        Age:     30,
        Email:   "alice@example.com",
        Hobbies: []string{"reading", "coding", "hiking"},
    }

    // 1. 保存（编码）到文件
    err := saveWithGob("person.gob", p)
    if err != nil {
        log.Fatal("保存失败:", err)
    }
    fmt.Println("✅ 数据已保存到 person.gob（使用 gob 编码）")
	v, err := loadWithGob("person.gob")
    if err != nil {
        log.Fatal("加载失败:", err)
    }

    fmt.Printf("🔁 从 gob 文件加载的数据: %+v\n", v)

}

func saveWithGob(filename string, data any) error {
    file, err := os.Create(filename)
    if err != nil {
        return err
    }
    defer file.Close()

    encoder := gob.NewEncoder(file)
    err = encoder.Encode(data)
    if err != nil {
        return err
    }

    return nil
}
func loadWithGob(filename string) (Person, error) {
    var p Person

    file, err := os.Open(filename)
    if err != nil {
        return p, err
    }
    defer file.Close()

    decoder := gob.NewDecoder(file)
    err = decoder.Decode(&p)
    if err != nil {
        return p, err
    }

    return p, nil
}



