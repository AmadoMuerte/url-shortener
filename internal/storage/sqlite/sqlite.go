package sqlite

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/AmadoMuerte/url-shortener/internal/storage"
	"github.com/mattn/go-sqlite3"
)

type Storage struct {
	db *sql.DB
}

func New(storagePath string) (*Storage, error) {
	const op = "storage.sqlite.New"

	db, err := sql.Open("sqlite3", storagePath)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	stmt, err := db.Prepare(`
	CREATE TABLE IF NOT EXISTS url(
	    id INTEGER PRIMARY KEY,
	    alias TEXT NOT NULL UNIQUE,
	    url TEXT NOT NULL);
	CREATE INDEX IF NOT EXISTS idx_alias ON url(alias);
	`)

	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	_, err = stmt.Exec()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &Storage{db: db}, err
}

func (s *Storage) GetIdByUrl(urlName string) (int64, error) {
	const storageOperation = "storage.sqlite.GetIdByUrl"

	stmt, err := s.db.Prepare("SELECT id FROM url WHERE url = ?")
	if err != nil {
		return 0, fmt.Errorf("%s: %w", storageOperation, err)
	}

	row := stmt.QueryRow(urlName)
	var id int64
	err = row.Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", storageOperation, err)
	}

	return id, nil
}

func (s *Storage) SaveURL(urlToSave, alias string) (int64, error) {
	const storageOperation = "storage.sqlite.SaveURL"

	stmt, err := s.db.Prepare("INSERT INTO url(url, alias) VALUES (?, ?)")
	if err != nil {
		return 0, fmt.Errorf("%s: %w", storageOperation, err)
	}
	_, err = stmt.Exec(urlToSave, alias)
	if err != nil {
		var sqliteErr sqlite3.Error
		if errors.As(err, &sqliteErr) &&
			errors.Is(sqliteErr.ExtendedCode, sqlite3.ErrConstraintUnique) {
			return 0, fmt.Errorf("%s: %w", storageOperation, storage.ErrURLExists)
		}
		return 0, fmt.Errorf("%s: %w", storageOperation, err)
	}

	id, err := s.GetIdByUrl(urlToSave)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", id, err)
	}

	return id, nil
}

func (s *Storage) GetUrl(alias string) (string, error) {
	const storageOperation = "storage.sqlite.GetUrl"

	stmt, err := s.db.Prepare("SELECT url FROM url WHERE alias = ?")
	if err != nil {
		return "", fmt.Errorf("%s: %w", storageOperation, err)
	}

	var resultUrl string
	err = stmt.QueryRow(alias).Scan(&resultUrl)
	if errors.Is(err, sql.ErrNoRows) {
		return "", storage.ErrURLNotFound
	}
	if err != nil {
		return "", fmt.Errorf("%s: %w", storageOperation, err)
	}

	return resultUrl, nil
}
