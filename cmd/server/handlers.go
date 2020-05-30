package main

import "net/http"

// getStoryArc retrieves a part of the choose your own adventure story based
// on the URL path.
//
// The path must start with /arc/. Subsequent path segments are treated as nested
// keys for the JSON formatted story arc. For example /arc/intro gets the
// introduction arc and /arc/intro/title will get the title of the introduction arc.
func getStoryData(w http.ResponseWriter, r *http.Request) {
	// Ensure valid path and method or return 400
	path, err := getRedisPath(r.URL.Path)
	if err != nil || r.Method != http.MethodGet {
		http.Error(w, "Bad Request: Invalid Path", http.StatusBadRequest)
		return
	}

	// Get Data
	story, err := getArcFromDB(path)
	if err != nil {
		if err.Error() == "data not found" {
			http.NotFound(w, r)
			return
		}

		if err.Error() == "problem converting data" {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	}

	// Send successful response
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(story)
}
