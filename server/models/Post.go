package models

import "time"

type Post struct {
	ID       int       `json:"id,omitempty"`
	PostedBy int       `json:"postedby,omitempty"`
	Title    string    `json:"title,omitempty"`
	Content  string    `json:"content,omitempty"`
	PostedAt time.Time `json:"postedat,omitempty"`
}
