package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type car struct {
	Name  string
	Price int
	Color string
}

func main() {
	car1 := car{"BMW", 100000, "red"}
	car2 := car{"Audi", 200000, "blue"}

	cars := []car{car1, car2}

	json, err := json.Marshal(cars)
	if err != nil {
		log.Panic(err)
	}

	fmt.Println(string(json))
}
