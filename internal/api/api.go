package api

import (
	"encoding/json"
	"fmt"
	"github.com/migueldcdev/file-integrity-service/internal/db"
	"net/http"
)

func getHashedFiles(w http.ResponseWriter, req *http.Request) {
	hashedFiles, err := db.GetAllHashedFiles()

	if err != nil {
		internalError := http.StatusInternalServerError
		http.Error(w, err.Error(), internalError)
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(hashedFiles); err != nil {
		internalError := http.StatusInternalServerError
		http.Error(w, err.Error(), internalError)
	}
}

func RunAPIServer() {

	fmt.Println("File Integrity API server running on localhost:8090")
	http.HandleFunc("/hashedfiles", getHashedFiles)
	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		fmt.Printf("API server error: %v\n", err)
	}

}
