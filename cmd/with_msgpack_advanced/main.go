package main

import (
	"encoding/json"
	"fmt"
	"memchached-go/lib/msgpack_usage"

	"github.com/bradfitz/gomemcache/memcache"
)

type person struct {
	Firstname string
	Lastname  string
}

func (p person) String() string {
	return p.Firstname + " " + p.Lastname
}

var cache = memcache.New("localhost:11211")

func main() {
	err := cache.Set(&memcache.Item{
		Key: "person",
		Value: msgpack_usage.Marshall(
			person{
				Firstname: "Addu",
				Lastname:  "Thoma",
			},
		),
		Expiration: 10,
	})
	if err != nil {
		panic(err)
	}
	b, err := cache.Get("person")
	if err != nil {
		panic(err)
	}
	d := msgpack_usage.Unmarshall(b.Value)
	var p person
	toStruct(d.(map[string]interface{}), &p)
	fmt.Println(p)
}

func toStruct(data interface{}, dest interface{}) {
	jb, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal(jb, dest); err != nil {
		panic(err)
	}
}
