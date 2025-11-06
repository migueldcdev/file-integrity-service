package api

import (
	"encoding/json"
	"fmt"
	"github.com/migueldcdev/file-integrity-service/internal/db"
	"github.com/migueldcdev/file-integrity-service/internal/hash"
	"net/http"
)

type UpdateFileRequest struct {
	Path string `json:"path"`
}

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

func checkFileIntegrity(w http.ResponseWriter, req *http.Request) {
	path := req.URL.Query().Get("path")

	if path == "" {
		badRequestError := http.StatusBadRequest
		http.Error(w, "The path to the file is required.", badRequestError)
	}

	h, err := db.GetFileHashByPath(path)

	if err != nil {
		internalError := http.StatusInternalServerError
		http.Error(w, err.Error(), internalError)
	}

	ch, err := hash.ComputeFileHash(path)

	if err != nil {
		internalError := http.StatusInternalServerError
		http.Error(w, err.Error(), internalError)
	}

	if h == ch {
		fmt.Fprintln(w, "The file is intact and verified.")
	} else {
		fmt.Fprintln(w, "Integrity check failed for file. Possible tampering detected!")
	}
}

func updateFileHash(w http.ResponseWriter, req *http.Request) {
	var reqBody UpdateFileRequest
	if err := json.NewDecoder(req.Body).Decode(&reqBody); err != nil {
		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
		return
	}

	path := reqBody.Path

	if path == "" {
		badRequestError := http.StatusBadRequest
		http.Error(w, "The path to the file is required.", badRequestError)
		return
	}

	h, err := hash.ComputeFileHash(path)

	if err != nil {
		internalError := http.StatusInternalServerError
		http.Error(w, err.Error(), internalError)
		return
	}

	err = db.UpdateFileHash(path, h)

	if err != nil {
		internalError := http.StatusInternalServerError
		http.Error(w, err.Error(), internalError)
		return
	}

	fmt.Fprintln(w, "File hash succesfuly updated")
}

func RunAPIServer() {
	fmt.Println("File Integrity API server running on localhost:8090")

	http.HandleFunc("/hashedfiles", getHashedFiles)
	http.HandleFunc("/check-file-integrity", checkFileIntegrity)
	http.HandleFunc("/update-file-hash", updateFileHash)

	err := http.ListenAndServe(":8090", nil)

	if err != nil {
		fmt.Printf("API server error: %v\n", err)
	}

}
