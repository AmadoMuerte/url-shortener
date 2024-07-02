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

type UrlInfo struct {
	Id   int64  `json:"id"`
	Name string `json:"alias"`
	Url  string `json:"url"`
}

type UrlData struct {
	Id    int64
	Url   string
	Alias string
}

func New(storagePath string) (*Storage, error) {
	const op = "storage.sqlite.New"

	db, err := sql.Open("sqlite3", storagePath)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	stmt, err := db.Prepare(`
	CREATE TABLE IF NOT EXISTS "url" (
		id INTEGER PRIMARY KEY,
		alias TEXT NOT NULL UNIQUE,
		url TEXT NOT NULL
	);
	CREATE INDEX IF NOT EXISTS idx_alias ON "url"(alias);
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
	const op = "storage.sqlite.GetIdByUrl"

	stmt, err := s.db.Prepare("SELECT id FROM url WHERE url = ?")
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	row := stmt.QueryRow(urlName)
	var id int64
	err = row.Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return id, nil
}

func (s *Storage) SaveURL(urlToSave, alias string) (int64, error) {
	const op = "storage.sqlite.SaveURL"

	stmt, err := s.db.Prepare("INSERT INTO url(url, alias) VALUES (?, ?)")
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}
	_, err = stmt.Exec(urlToSave, alias)
	if err != nil {
		var sqliteErr sqlite3.Error
		if errors.As(err, &sqliteErr) &&
			errors.Is(sqliteErr.ExtendedCode, sqlite3.ErrConstraintUnique) {
			return 0, fmt.Errorf("%s: %w", op, storage.ErrURLExists)
		}
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	id, err := s.GetIdByUrl(urlToSave)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return id, nil
}

func (s *Storage) GetUrlByAlias(alias string) (string, error) {
	const op = "storage.sqlite.GetUrlByAlias"

	stmt, err := s.db.Prepare("SELECT url FROM url WHERE alias = ?")
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	var resultUrl string
	err = stmt.QueryRow(alias).Scan(&resultUrl)
	if errors.Is(err, sql.ErrNoRows) {
		return "", storage.ErrURLNotFound
	}
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	return resultUrl, nil
}

func (s *Storage) DeleteUrl(alias string) error {
	const op = "storage.sqlite.DeleteUrl"

	stmt, err := s.db.Prepare("DELETE FROM url WHERE alias = ?")
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	var res string
	err = stmt.QueryRow(alias).Scan(&res)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return storage.ErrURLNotFound
		}
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (s *Storage) GetAllAlias() ([]UrlInfo, error) {
	const op = "storage.sqlite.GetAllAlias"

	stmt, err := s.db.Prepare("SELECT id, alias, url FROM url")
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			_ = fmt.Errorf("%s: %w", op, err)
		}
	}(stmt)

	var aliases []UrlInfo

	rows, err := stmt.Query()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			_ = fmt.Errorf("%s: %w", op, err)
		}
	}(rows)

	// Перебор результатов
	for rows.Next() {
		var a UrlInfo
		if err := rows.Scan(&a.Id, &a.Name, &a.Url); err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		aliases = append(aliases, a)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return aliases, nil
}

func (s *Storage) CheckUrlExist(id int64) (int64, error) {
	const op = "storage.sqlite.CheckUrlExist"

	stmt, err := s.db.Prepare("SELECT id FROM url WHERE id = ?")
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			_ = fmt.Errorf("%s: %w", op, err)
		}
	}(stmt)

	var resultId int64
	err = stmt.QueryRow(id).Scan(&resultId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, storage.ErrURLNotFound
		}
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return resultId, nil
}

func (s *Storage) RemoveUrl(id int64) (int64, error) {
	const op = "storage.sqlite.RemoveUrl"

	resultId, err := s.CheckUrlExist(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, storage.ErrURLNotFound
		}
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	stmt, err := s.db.Prepare("DELETE FROM url WHERE id = ?")
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			_ = fmt.Errorf("%s: %w", op, err)
		}
	}(stmt)

	_, err = stmt.Exec(id)
	if err != nil {
		var sqliteErr sqlite3.Error
		if errors.As(err, &sqliteErr) && errors.Is(sqliteErr.ExtendedCode, sqlite3.ErrConstraintUnique) {
			return 0, fmt.Errorf("%s: %w", op, storage.ErrURLExists)
		}
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return resultId, nil
}

func (s *Storage) GetUrl(id int64) (UrlData, error) {
	const op = "storage.sqlite.GetUrl"
	emptyRes := UrlData{Id: 0, Url: "", Alias: ""}

	stmt, err := s.db.Prepare("SELECT id, url, alias FROM url WHERE id = ?")
	if err != nil {
		return emptyRes, fmt.Errorf("%s: %w", op, err)
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			_ = fmt.Errorf("%s: %w", op, err)
		}
	}(stmt)

	var res UrlData
	err = stmt.QueryRow(id).Scan(&res.Id, &res.Url, &res.Alias)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return emptyRes, storage.ErrURLNotFound
		}
		return emptyRes, fmt.Errorf("%s: %w", op, err)
	}

	return res, nil
}

func (s *Storage) UpdateUrl(id int64, url, alias string) (UrlData, error) {
	const op = "storage.sqlite.UpdateUrl"
	emptyRes := UrlData{Id: 0, Url: "", Alias: ""}

	stmt, err := s.db.Prepare("UPDATE url SET url = ?, alias = ? WHERE id = ?")
	if err != nil {
		return emptyRes, fmt.Errorf("%s: %w", op, err)
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			_ = fmt.Errorf("%s: %w", op, err)
		}
	}(stmt)

	_, err = stmt.Exec(url, alias, id)
	if err != nil {
		var sqliteErr sqlite3.Error
		if errors.As(err, &sqliteErr) && errors.Is(sqliteErr.ExtendedCode, sqlite3.ErrConstraintUnique) {
			return emptyRes, fmt.Errorf("%s: %w", op, storage.ErrURLExists)
		}
		return emptyRes, fmt.Errorf("%s: %w", op, err)
	}

	return UrlData{Id: id, Url: url, Alias: alias}, nil
}
