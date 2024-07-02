package storage

import "errors"

var (
	ErrURLNotFound = errors.New("url.ts not found")
	ErrURLExists   = errors.New("url.ts exists")
)
