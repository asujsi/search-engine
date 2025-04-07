package main

import (
	"backend/api"
	"backend/search"
	"fmt"
	"log"
	"net/http"
)

func main() {
	data, err := search.LoadParquetDataFromDir("data")
	if err != nil {
		log.Fatal("Failed to load Parquet files:", err)
	}

	engine := search.NewSearchEngine(data)
	api.InitSearchEngine(engine)

	http.HandleFunc("/search", api.SearchHandler)

	fmt.Println("Listening on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
