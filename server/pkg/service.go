package pkg

import (
	"encoding/json"
	"fmt"
	"fruits_crud/server/db"
	"log"
	"time"
)

func FruitInsertService(reqbody Fruit) bool {

	res := FruitInsert(reqbody)

	if res == true {
		data := GetFruits()
		jsonData, _ := json.Marshal(data)
		r := db.Client.Set("all", jsonData, time.Minute*20)

		fmt.Println("test", r)
	}

	return res
}

func GetFruitsService() (f []Fruit) {

	val, err := db.Client.Get("all").Bytes()
	json.Unmarshal(val, &f)
	if err != nil {
		log.Println("error while unmarshalling data: ", err)
	}
	fmt.Println("data from redis: ", f, err)
	if len(f) > 0 {
		return f
	}

	// expire, _ := db.Client.TTL("all").Result()
	// fmt.Println(expire)

	data := GetFruits()
	return data
}

func UpdateFruitsService(reqbody Fruit) bool {

	res := FruitUpdate(reqbody)

	return res
}

func DeleteFruitsService(id int) bool {

	res := FruitDelete(id)

	return res
}
