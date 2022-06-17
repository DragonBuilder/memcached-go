package main

import (
	"fmt"
	"memchached-go/lib/msgpack_usage"

	"github.com/bradfitz/gomemcache/memcache"
)

var cache = memcache.New("localhost:11211")

func main() {
	err := cache.Set(&memcache.Item{
		Key:        "name",
		Value:      msgpack_usage.Marshall("aneesh"),
		Expiration: 10,
	})
	if err != nil {
		panic(err)
	}
	b, err := cache.Get("name")
	if err != nil {
		panic(err)
	}
	d := msgpack_usage.Unmarshall(b.Value)
	fmt.Println(d)
}
