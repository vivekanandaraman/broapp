package main

import gojsonq "github.com/thedevsaddam/gojsonq"

func main() {
	const json = `  [
    {
      "Name": "Alice", "Age": 25
    },
    {
      "Name": "Baron", "Age": 25
    }
    ]`

	name := gojsonq.New().FromString(json).Find("name")
	println(name.(string)) 
}
