package main

import (
	"fmt"
	"log"
	"memchached-go/lib"

	"github.com/bradfitz/gomemcache/memcache"
)

var cache = memcache.New("localhost:11211")

func main() {
	if err := cache.Set(&memcache.Item{
		Key:        "name",
		Value:      lib.Serialize("aneesh"),
		Expiration: 10,
	}); err != nil {
		panic(err)
	}

	val, err := cache.Get("name")
	if err != nil {
		log.Fatalln(err.Error())
	}
	var r string
	lib.DeSerialize(val.Value, &r)
	fmt.Println(r)
}
