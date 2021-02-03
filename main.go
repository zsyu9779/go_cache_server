package main

import (
	http "go_cache_server/HTTP"
	"go_cache_server/cache"
)

func main() {
	//c :=cache.New("inmemory")
	//http.New(c).Listen()

	c :=cache.New("inmemory")
	http.New(c).Listen()
}
