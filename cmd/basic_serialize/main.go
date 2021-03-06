package main

import (
	"fmt"
	"log"
	"memchached-go/lib/gob_usage"

	"github.com/bradfitz/gomemcache/memcache"
)

var cache = memcache.New("localhost:11211")

func main() {
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
	var r string
	gob_usage.DeSerialize(val.Value, &r)
	fmt.Println(r)
}
