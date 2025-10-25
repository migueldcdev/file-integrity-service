package service

import (
	"fmt"
	"github.com/migueldcdev/file-integrity-service/internal/db"
	"github.com/migueldcdev/file-integrity-service/internal/hash"
	"io/fs"
	"path/filepath"
	"time"
)

func WalkDirAndHashFiles(dirPath string) error {
	fmt.Println("Starting hashing service...")
	return filepath.WalkDir(dirPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			fmt.Printf("Error accessing %s: %v\n", path, err)
			return nil
		}

		if d.IsDir() {
			return nil
		}

		h, err := hash.ComputeFileHash(path)

		if err != nil {
			return err
		}

		i, err := d.Info()

		if err != nil {
			return err
		}

		s := i.Size()

		hf := db.HashedFile{
			Path:         path,
			Hash:         h,
			Size:         s,
			Created_at:   time.Now(),
			Last_checked: time.Now(),
		}

		err = db.SaveFileHash(hf)

		if err != nil {
			return err
		}

		return nil
	})
}
