package api

import (
	"backend/search"
	"encoding/json"
	"net/http"
)

var engine *search.SearchEngine

func InitSearchEngine(e *search.SearchEngine) {
	engine = e
}

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query")
	if query == "" {
		http.Error(w, "Query param missing", http.StatusBadRequest)
		return
	}

	results, count, duration := engine.Search(query)
	resp := map[string]interface{}{
		"query":   query,
		"matches": count,
		"time_ms": duration.Milliseconds(),
		"results": results,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
