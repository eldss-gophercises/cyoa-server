package main

import (
	"errors"
	"fmt"
	"strings"
)

// getRedisPath ensures that a given endpoint starts with "arc".
//
// If the path is valid, returns the redis style path of the endpoint argument
// (the part after "arc/"). For example, if the path is "arc/test1/test2"
// the returned string will be "test1.test2", "arc/test1" will return "test1".
func getRedisPath(path string) (string, error) {
	// Separate path segments and ensure first part is "arc"
	pathSlice := strings.Split(path, "/")

	// Check index 1 due to path starting with /, 0 == ""
	if len(pathSlice) < 2 || pathSlice[1] != "arc" {
		return "", errors.New("path does not start with 'arc'")
	}

	redisPath := fmt.Sprint(".", strings.Join(pathSlice[2:], "."))
	return redisPath, nil
}

// getArcFromDB returns story data from the Redis DB.
//
// It can return any data available as long as the redisPath returns
// valid data. For example, "." will return the whole story object, ".intro"
// will return the intro section, and ".intro.title" will return the title of
// the intro section. An invalid path will return an error.
func getArcFromDB(redisPath string) ([]byte, error) {
	// Get story arc from Redis
	// rh is global in main.go
	story, err := rh.JSONGet("story", redisPath)
	if err != nil {
		return nil, errors.New("data not found:" + err.Error())
	}

	// Convert from interface to slice of bytes or return 500
	storyBytes, ok := story.([]byte)
	if !ok {
		return nil, errors.New("problem converting data")
	}

	return storyBytes, nil
}
