package pkg

import "encoding/json"

type Fruit struct {
	Id        int    `json:"id" `
	FruitName string `json:"fruit_name" binding:"required"`
}

func (i Fruit) MarshalBinary() ([]byte, error) {
	return json.Marshal(i)
}
func (i Fruit) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, &i)
}
