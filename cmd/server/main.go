package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/go-redis/redis"
	"github.com/nitishm/go-rejson"
)

var rh *rejson.Handler = rejson.NewReJSONHandler()

func setRedisClient() {
	// Set up redis client with default options
	client := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})

	// Use ReJSON with the redis client
	rh.SetGoRedisClient(client)
}

func loadData() {
	// Read starter data
	log.Println("Reading data file...")
	// Binary must be run in the top level directory for this to work (for Docker)
	data, err := ioutil.ReadFile("./data/story.json")
	if err != nil {
		log.Panicln("couldn't read data file:", err)
	}

	// Unmarshal the JSON
	var s story
	err = json.Unmarshal(data, &s)
	if err != nil {
		log.Panicln("couldn't unmarshal json:", err)
	}

	// Put data in redis
	log.Println("Loading data to Redis...")
	_, err = rh.JSONSet("story", ".", s)
	if err != nil {
		log.Panicln("couldn't load data:", err)
	}
}

func main() {
	setRedisClient()
	loadData()
	http.HandleFunc("/", getStoryData)
	log.Println("Starting server on port 8085.")
	log.Fatal(http.ListenAndServe(":8085", nil))
}
