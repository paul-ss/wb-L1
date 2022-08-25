package main

import (
	"encoding/json"
	"fmt"
)

type Data struct {
	Id     int
	Email  string
	Rating int
}

func main() {
	// client code
	a := NewAdapter()

	data, err := a.GetData()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("id = %d\nemail = %s\nrating = %d\n", data.Id, data.Email, data.Rating)

}

// Adapter

func NewAdapter() *Adapter {
	return &Adapter{
		db: NewDatabase(),
	}
}

type Adapter struct {
	db *Database
}

func (a *Adapter) GetData() (*Data, error) {
	var data Data
	err := json.Unmarshal(a.db.GetData(), &data)
	return &data, err
}

// Adaptable

func NewDatabase() *Database {
	return &Database{
		data: []byte(`{"Id": 1, "Email": "123@mail.ru", "Rating": 123}`),
	}
}

type Database struct {
	data []byte
}

func (d *Database) GetData() []byte {
	return d.data
}
