package main

import (
    "encoding/json"
    "log"
    "os"
    "strings"
)

func main() {
  const jsonData = `
    {"Name": "Alice", "Age": 25}
    {"Name": "Bob", "Age": 22}
  `
  reader := strings.NewReader(jsonData)
  writer := os.Stdout
  dec := json.NewDecoder(reader)
  enc := json.NewEncoder(writer)
    for {
        var v map[string]interface{}
        if err := dec.Decode(&v); err != nil {
            log.Println(err)
            return
        }
        for k := range v {
            if k != "Name" {
                delete(v, k)
            }
        }
        if err := enc.Encode(&v); err != nil {
            log.Println(err)
        }
    }
}
