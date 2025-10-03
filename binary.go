package main

import (
    "bytes"
    "encoding/binary"
    "fmt"
    "log"
    "os"
)

// 要保存的数据结构
type MyData struct {
    ID     int32
    Name   string
    Score  float32
    Active bool
}

// 保存到二进制文件
func saveToFile(filename string, data MyData) error {
    file, err := os.Create(filename)
    if err != nil {
        return err
    }
    defer file.Close()

    buf := new(bytes.Buffer)

    // 写入 ID (int32)
    err = binary.Write(buf, binary.LittleEndian, data.ID)
    if err != nil {
        return err
    }

    // 写入 Name 长度 (uint16)，然后是 Name 的字节
    nameBytes := []byte(data.Name)
    nameLen := uint16(len(nameBytes))
    err = binary.Write(buf, binary.LittleEndian, nameLen)
    if err != nil {
        return err
    }
    err = binary.Write(buf, binary.LittleEndian, nameBytes)
    if err != nil {
        return err
    }

    // 写入 Score (float32)
    err = binary.Write(buf, binary.LittleEndian, data.Score)
    if err != nil {
        return err
    }

    // 写入 Active (bool，其实可以转成 uint8)
    activeByte := uint8(0)
    if data.Active {
        activeByte = 1
    }
    err = binary.Write(buf, binary.LittleEndian, activeByte)
    if err != nil {
        return err
    }

    // 全部写入文件
    _, err = file.Write(buf.Bytes())
    return err
}

// 从二进制文件读取
func loadFromFile(filename string) (MyData, error) {
    file, err := os.Open(filename)
    if err != nil {
        return MyData{}, err
    }
    defer file.Close()

    var data MyData

    buf := make([]byte, 1024) // 或者使用 bufio.NewReader 等更灵活方式
    n, err := file.Read(buf)
    if err != nil {
        return MyData{}, err
    }

    reader := bytes.NewReader(buf[:n])

    // 读取 ID
    err = binary.Read(reader, binary.LittleEndian, &data.ID)
    if err != nil {
        return MyData{}, err
    }

    // 读取 Name 长度
    var nameLen uint16
    err = binary.Read(reader, binary.LittleEndian, &nameLen)
    if err != nil {
        return MyData{}, err
    }

    // 读取 Name
    nameBytes := make([]byte, nameLen)
    err = binary.Read(reader, binary.LittleEndian, &nameBytes)
    if err != nil {
        return MyData{}, err
    }
    data.Name = string(nameBytes)

    // 读取 Score
    err = binary.Read(reader, binary.LittleEndian, &data.Score)
    if err != nil {
        return MyData{}, err
    }

    // 读取 Active
    var activeByte uint8
    err = binary.Read(reader, binary.LittleEndian, &activeByte)
    if err != nil {
        return MyData{}, err
    }
    data.Active = (activeByte == 1)

    return data, nil
}

func main() {
    // 准备数据
    original := MyData{
        ID:     1001,
        Name:   "MyApp Data",
        Score:  95.5,
        Active: true,
    }

    // 保存到文件
    err := saveToFile("data.bin", original)
    if err != nil {
        log.Fatal("保存失败:", err)
    }
    fmt.Println("数据已保存到 data.bin")

    // 从文件读取
    loaded, err := loadFromFile("data.bin")
    if err != nil {
        log.Fatal("读取失败:", err)
    }
    fmt.Printf("从文件读取的数据: %+v\n", loaded)
}
