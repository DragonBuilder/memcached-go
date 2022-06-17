package main

import (
	"encoding/json"
	"fmt"

	"github.com/bradfitz/gomemcache/memcache"
)

var cache = memcache.New("localhost:11211")

type person struct {
	Firstname string
	Lastname  string
}

func (p person) String() string {
	return p.Firstname + " " + p.Lastname
}

func main() {
	save("person", person{
		Firstname: "Thomas",
		Lastname:  "Shelby",
	})

	d := retrieve("person")

	var p person
	decode(d, &p)
	fmt.Println(p)

	save("primitive", "this is a primitive string value")

	d = retrieve("primitive")

	var s string
	decode(d, &s)
	fmt.Println(s)
}

func save(key string, value interface{}) {
	err := cache.Set(&memcache.Item{
		Key:        key,
		Value:      encode(value),
		Expiration: 10,
	})
	if err != nil {
		panic(err)
	}
}

func retrieve(key string) []byte {
	item, err := cache.Get(key)
	if err != nil {
		panic(err)
	}
	return item.Value
}

func encode(d interface{}) []byte {
	b, err := json.Marshal(d)
	if err != nil {
		panic(err)
	}
	return b
}

func decode(data []byte, dest interface{}) {
	err := json.Unmarshal(data, dest)
	if err != nil {
		panic(err)
	}
}
