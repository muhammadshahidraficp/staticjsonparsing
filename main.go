package main

import (
	"encoding/json"
	"fmt"
)

var JSON = `{
    "name":"Mark Taylor",
    "jobtitle":"Software Developer",
    "phone":{
        "home":"123-466-799",
        "office":"564-987-654"
    },
    "email":"markt@gmail.com"
}`

func main() {

	var info map[string]interface{}
	json.Unmarshal([]byte(JSON), &info)

	fmt.Println(info["name"])
	fmt.Println(info["jobtitle"])
	fmt.Println(info["email"])
	fmt.Println(info["phone"].(map[string]interface{})["home"])
	fmt.Println(info["phone"].(map[string]interface{})["office"])
}
