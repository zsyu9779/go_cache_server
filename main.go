package main

import (
	"go_cache_server/TCP"
	"go_cache_server/cache"
)

func main() {
	//c :=cache.New("inmemory")
	//http.New(c).Listen()
	c :=cache.New("inmemory")
	TCP.New(c).Listen()
}
