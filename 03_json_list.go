//program to parse
package main

import (
    "encoding/json"
    "fmt"
)

func main() {
  var input = `
  [
  {
    "Name": "Alice", "Age": 25
  },
  {
    "Name": "Baron", "Age": 25
  }
  ]
  `
    var list1 []map[string]interface{}
    if err := json.Unmarshal([]byte(input), &list1); err != nil {
        panic(err)
    }
    for i := range list1 {
      fmt.Println(list1[i])
      for k,v := range list1[i] {
        switch v := v.(type) {
          case string:
            fmt.Println(k, v, "(string)")
          default:
            fmt.Println(k, v, "(unknown)")
        }
      }
    }

}
