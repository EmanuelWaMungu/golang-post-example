package post

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPost(t *testing.T) {

	title := "Is Platform Engineer a real job?"
	desc := "Put your thoughts here dude"

	post, err := NewPost(title, desc)

	assert.Nil(t, err)
	assert.NotNil(t, post)
	assert.Equal(t, title, post.Title)
	assert.Equal(t, desc, post.Desciption)

}
