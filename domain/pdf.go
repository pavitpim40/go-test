package domain

import (
	"time"
)

type PDF struct {
	ID          string
	Title       string
	Content     []byte
	Path        string
	Author      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	PageCount   int
	ContentType string
}
