package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
)

type prefixStructure struct {
	Prefix uuid.UUID `json:"prefix"`
}

func prefix(w http.ResponseWriter, r *http.Request) {
	prefix := prefixStructure{
		Prefix: uuid.New(),
	}

	if err := json.NewEncoder(w).Encode(prefix); err != nil {
		panic(err)
	}
}

func main() {
	http.HandleFunc("/prefix", prefix)

	fmt.Println("Started on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
