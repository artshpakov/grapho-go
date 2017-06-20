package models

import (
	"time"

	"github.com/go-pg/pg"
)

type Post struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	Title     string    `json:"title"`
	Text      string    `json:"text"`
	Slug      string    `json:"slug"`
	CreatedAt time.Time `json:"created_at"`
}

func FetchPosts(db *pg.DB, limit int, ids []interface{}, userID interface{}) ([]Post, error) {
	posts := make([]Post, 0)
	query := db.Model(&posts)
	if len(ids) > 0 {
		query = query.WhereIn("id IN (?)", ids...)
	}
	if userID != nil {
		query = query.WhereIn("user_id = ?", userID)
	}
	err := query.Order("created_at DESC").Limit(limit).Select()
	return posts, err
}

func PostsFindByAuthor(db *pg.DB, userID int) ([]Post, error) {
	posts := make([]Post, 0)
	err := db.Model(&posts).Where("user_id = ?", userID).Select()
	return posts, err
}
