package models

import (
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
)

// Comment struct
type Comment struct {
	ID        int       `json:"id" db:"id" rw:"r"`
	Content   string    `json:"content" db:"content"`
	AuthorID  string    `json:"author" db:"author"`
	PostID    int       `json:"postid" db:"postid"`
	CreatedBy string    `json:"created_by" db:"created_by"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedBy string    `json:"updated_by" db:"updated_by"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	Author    User      `json:"-" db:"-"`
}

// TableName overrides the table name used by Pop.
func (p Comment) TableName() string {
	return "bl_comment_h"
}

// Comments array
type Comments []Comment

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
func (c *Comment) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: c.Content, Name: "Content"},
	), nil
}
