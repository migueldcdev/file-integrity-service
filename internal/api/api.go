package api

import (
	"fmt"
	"net/http"
)

func getHashedFiles(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "hi\n")
}

func RunAPIServer() {

	fmt.Println("File Integrity API server running on localhost:8090")
	http.HandleFunc("/", getHashedFiles)
	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		fmt.Printf("API server error: %v\n", err)
	}

}
