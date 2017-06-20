package models

import (
	"fmt"

	"github.com/go-pg/pg"
)

type User struct {
	ID      int         `json:"id"`
	Email   string      `json:"email"`
	Slug    string      `json:"slug"`
	Profile UserProfile `json:"profile"`
}

type UserProfile struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Location  string `json:"location"`
	Bio       string `json:"bio"`
}

func (u *User) Name() string {
	if u.Profile.Firstname != "" && u.Profile.Lastname != "" {
		return fmt.Sprintf("%v %v", u.Profile.Firstname, u.Profile.Lastname)
	} else {
		return u.Email
	}
}

func (u *User) Posts(db *pg.DB) ([]Post, error) {
	return PostsFindByAuthor(db, u.ID)
}

func UsersFetch(db *pg.DB, ids []interface{}) ([]User, error) {
	users := make([]User, 0)
	query := db.Model(&users)
	if len(ids) > 0 {
		query = query.WhereIn("id IN (?)", ids...)
	}
	err := query.Select()
	return users, err
}

func UserFindById(db *pg.DB, id int) (user User, err error) {
	err = db.Model(&user).Where("id = ?", id).Limit(1).Select()
	return user, err
}
