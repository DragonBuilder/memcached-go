package main

import (
	"fmt"
	"log"
	"memchached-go/lib/gob_usage"

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

/*
Because the interface{} variable r is not passed as pointer, the result isn't visible on printing,
But passing it as pointer results in an error.
*/
func main() {
	gob_usage.RegisterType(&person{})
	if err := cache.Set(&memcache.Item{
		Key: "name",
		Value: gob_usage.Serialize(&person{
			Firstname: "Ethan",
			Lastname:  "Hunt",
		}),
		Expiration: 10,
	}); err != nil {
		panic(err)
	}

	val, err := cache.Get("name")
	if err != nil {
		log.Fatalln(err.Error())
	}
	var r interface{}
	gob_usage.InterfaceDeserialise(val.Value, r)
	fmt.Println(r)

}

func saveString() {
	gob_usage.RegisterType(string("ss"))
	if err := cache.Set(&memcache.Item{
		Key:        "name",
		Value:      gob_usage.Serialize("aneesh"),
		Expiration: 10,
	}); err != nil {
		panic(err)
	}

	val, err := cache.Get("name")
	if err != nil {
		log.Fatalln(err.Error())
	}
	var r interface{}
	gob_usage.InterfaceDeserialise(val.Value, r)
	fmt.Println(r)
}
