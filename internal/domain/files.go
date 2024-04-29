package domain

import "time"

type FileListItem struct {
	ID        string
	Name      string
	Type      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type File struct {
	FileListItem
	Data []byte
}
