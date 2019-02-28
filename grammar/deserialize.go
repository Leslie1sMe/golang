//Tips:反序列化的数据类型应该和原来序列化前的数据类型一致
package main

import (
	"encoding/json"
	"log"
	"fmt"
)

//反序列化
//将json反序列化成切片，map,struct
type JsonDeserialize struct {
	Name     string
	Age      int
	Birthday string
	Skill    string
}

func testDeserializeStruct() {
	jsonStr := "{\"Name\":\"liming\",\"Age\":18,\"Birthday\":\"2019-02-28\",\"Skill\":\"Coding\"}"
	var jsonDe JsonDeserialize
	err := json.Unmarshal([]byte(jsonStr), &jsonDe)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("name=%v age=%v birthday=%v skill=%v", jsonDe.Name, jsonDe.Age, jsonDe.Birthday, jsonDe.Skill)
}
func testDeserializeSlice() {
	jsonStr := "[{\"Name\":\"liming\",\"Age\":18,\"Birthday\":\"2019-02-28\",\"Skill\":\"Coding\"}," +
		"{\"Name\":\"leslie\",\"Age\":27,\"Birthday\":\"1999-08-19\",\"Skill\":\"Eating\"}]"
	var a [] map[string]interface{}
	err := json.Unmarshal([]byte(jsonStr), &a)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(a)
}

//反序列化不需要make，make已经集成到了Unmarshal中
func testDeserializeMap() {
	jsonStr := "{\"Name\":\"liming\",\"Age\":18,\"Birthday\":\"2019-02-28\",\"Skill\":\"Coding\"}"
	var a map[string]interface{}
	err := json.Unmarshal([]byte(jsonStr), &a)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(a)
}

func main() {
	//testDeserializeStruct()
	//testDeserializeSlice()
	//testDeserializeMap()
}
