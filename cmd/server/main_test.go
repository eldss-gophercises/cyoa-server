package main

import (
	"bytes"
	"testing"

	"github.com/go-redis/redis"
)

func setRedisClient() {
	// Set up redis client with default options
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	// Use ReJSON with the redis client
	rh.SetGoRedisClient(client)
}

func TestValidatePath_HappyPath(t *testing.T) {
	testCases := []struct {
		in  string
		out string
	}{
		{"/arc/", "."},
		{"/arc/test", ".test"},
		{"/arc/test1/test2", ".test1.test2"},
	}

	for _, test := range testCases {
		result, err := getRedisPath(test.in)
		if err != nil {
			t.Errorf("For in=%s, expected %s, got error %v", test.in, test.out, err)
		}
		if result != test.out {
			t.Errorf("For in=%s, expected %s, got %s", test.in, test.out, result)
		}
	}
}

func TestValidatePath_BadPaths(t *testing.T) {
	testCases := []string{
		"/",
		"",
		"/notarc",
		"/notarc/with/path",
	}

	for _, test := range testCases {
		result, err := getRedisPath(test)
		if err == nil {
			t.Errorf("Bad path %s should have returned an error, got %s", test, result)
		}
	}
}

func TestGetArcFromDB_HappyPaths(t *testing.T) {
	setRedisClient()

	testCases := []struct {
		in       string
		contains []byte
	}{
		{".intro.title", []byte("The Little Blue Gopher")},
		{".intro.options", []byte("That story about the Sticky Bandits")},
	}

	for _, test := range testCases {
		result, err := getArcFromDB(test.in)
		if !bytes.Contains(result, test.contains) {
			t.Errorf(
				"For in=%s, expected result to contain %s, but did not. Result=%q",
				test.in,
				test.contains,
				result,
			)
		}

		if err != nil {
			t.Errorf("For in=%s, encountered error %v.", test.in, err)
		}
	}
}
