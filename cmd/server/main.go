package main

import (
	"log"
	"net/http"

	"github.com/go-redis/redis"
	"github.com/nitishm/go-rejson"
)

var rh *rejson.Handler = rejson.NewReJSONHandler()

func main() {
	// Set up redis client with default options
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	// Use ReJSON with the redis client
	rh.SetGoRedisClient(client)

	http.HandleFunc("/", getStoryArc)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
