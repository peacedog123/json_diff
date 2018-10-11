package main

import (
  "encoding/json"
  "flag"
  "log"
  "os"

  "github.com/peacedog123/json_diff/data_type"
)

var (
  file_path_1 string
  file_path_2 string
)

func read_file(file_path string) []byte {
  file,err := os.Open(file_path)
  if err != nil {
    log.Fatal(err)
    return nil
  }
  defer file.Close()

  fileinfo,err := file.Stat()
  if err != nil {
    log.Fatal(err)
    return nil
  }
  fileSize := fileinfo.Size()
  buffer := make([]byte,fileSize)
  bytesread,err := file.Read(buffer)
  if err != nil {
    log.Fatal(err)
    return nil
  }

  if (int64(bytesread) != fileSize) {
    log.Println("file read error! file_path: ", file_path)
    return nil
  }

  return buffer
}

func init() {
  flag.StringVar(&file_path_1, "f1", "", "file path 1")
  flag.StringVar(&file_path_2, "f2", "", "file path 2")
}

func main() {
  flag.Parse()

  if (file_path_1 == "" || file_path_2 == "") {
    log.Println("file name is empty, exit...")
    return
  }

  data1 := read_file(file_path_1)
  res1 := []data_type.RecallResult{}
  json.Unmarshal(data1, &res1)

  data2 := read_file(file_path_2)
  res2 := []data_type.RecallResult{}
  json.Unmarshal(data2, &res2)

  cm := data_type.CompareManager {
    P1: res1,
    P2: res2 }

  log.Println("data1 == data2 : ", cm.Match())
}

