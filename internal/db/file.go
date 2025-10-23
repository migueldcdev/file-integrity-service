package db

import (
	"time"
)

type HashedFile struct {
	Id           int
	Path         string
	Hash         string
	Size         int64
	Created_at   time.Time
	Last_checked time.Time
}

func SaveFileHash(hf HashedFile) error {
	var err error
	_, err = DB.Exec("INSERT INTO hash(path, hash, size, created_at, last_checked) VALUES($1, $2, $3, $4, $5)", hf.Path, hf.Hash, hf.Size, hf.Created_at, hf.Last_checked)

	if err != nil {
		return err
	}

	return nil
}
