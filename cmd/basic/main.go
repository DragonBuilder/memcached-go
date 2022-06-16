package main

import (
	"fmt"

	"github.com/bradfitz/gomemcache/memcache"
)

var cache = memcache.New("localhost:11211")

func main() {
	cacheErr := cache.Set(&memcache.Item{Key: "name", Value: []byte("aneesh"), Expiration: 10})
	if cacheErr != nil {
		panic(cacheErr)
	}

	val, err := cache.Get("name")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(val.Value))
}
