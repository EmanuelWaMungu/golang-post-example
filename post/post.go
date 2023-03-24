package post

import (
	"errors"

	"github.com/google/uuid"
)

type Post struct {
	ID         string
	Title      string
	Desciption string
}

func NewPost(title string, desc string) (*Post, error) {

	post := Post{
		ID:         uuid.NewString(),
		Title:      title,
		Desciption: desc,
	}

	err := post.Validate()

	if err != nil {
		return nil, err
	}

	return &post, nil

}

func (c Post) Validate() error {

	if c.Title == "" {
		return errors.New("Title is required")
	}

	if c.Desciption == "" {
		return errors.New("Description is required")
	}

	return nil
}
