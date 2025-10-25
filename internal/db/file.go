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

func GetAllHashedFiles() ([]HashedFile, error) {
	var err error
	rows, err := DB.Query("SELECT id, path, hash, size, created_at, last_checked FROM hash")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var hashedFiles []HashedFile

	for rows.Next() {
		var hf HashedFile
		err := rows.Scan(&hf.Id, &hf.Path, &hf.Hash, &hf.Size, &hf.Created_at, &hf.Last_checked)
		if err != nil {
			return nil, err
		}

		hashedFiles = append(hashedFiles, hf)

	}

	return hashedFiles, nil
}

func UpdateFileHash(hfId int, hash string) error {
	var err error
	_, err = DB.Exec("UPDATE hash SET hash=$1, created_at=$2 WHERE id=$3", hash, hfId, time.Now())

	if err != nil {
		return err
	}

	return nil
}
