package pkg

import (
	"fmt"
	"fruits_crud/server/db"
)

func FruitInsert(reqbody Fruit) bool {

	sql := `INSERT INTO fruits(fruit_name)VALUES($1)`

	_, err := db.DB.Exec(sql, reqbody.FruitName)

	if err != nil {
		return false
	}
	return true
}

func FruitUpdate(reqbody Fruit) bool {

	SQL := `UPDATE fruits SET  fruit_name=$1 WHERE id=$2`

	_, err2 := db.DB.Exec(SQL, reqbody.FruitName, reqbody.Id)

	if err2 != nil {
		//log.Fatal("ERror in update: ", err2)
		return false
	}

	return true
}

func GetFruits() []Fruit {
	data := []Fruit{}
	sql := `SELECT id,fruit_name FROM "fruits"`

	rows, err := db.DB.Query(sql)
	fmt.Println("errrrrr", err)

	row := Fruit{}

	for rows.Next() {

		rows.Scan(&row.Id, &row.FruitName)

		data = append(data, row)
	}

	return data

}

func FruitDelete(id int) bool {
	SQL := `DELETE FROM "fruits"  WHERE id=$1`

	_, err2 := db.DB.Exec(SQL, id)

	if err2 != nil {
		//log.Fatal("ERror in update: ", err2)
		return false
	}

	return true
}
