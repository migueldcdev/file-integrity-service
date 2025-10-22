package hash

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func ComputeFileHash(filePath string) (string, error) {

	file, err := os.Open(filePath)

	if err != nil {
		return "", err
	}

	defer file.Close()

	h := sha256.New()

	if _, err := io.Copy(h, file); err != nil {
		return "", err
	}

	return hex.EncodeToString(h.Sum(nil)), nil
}

func visit(path string, f os.FileInfo, err error) error {

	i, err := os.Stat(path)

	if err != nil {
		fmt.Printf("Path does no exist")
	}
	if !i.IsDir() {

		fmt.Printf("Visited: %s\n", path)
	}

	return nil
}

func WalkDirAndHashFiles(dirPath string) {
	err := filepath.Walk(dirPath, visit)
	if err != nil {
		fmt.Printf("filepath.Walk() returned %v\n", err)
	}
}
