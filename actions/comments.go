package actions

import (
	"strconv"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
	"github.com/herwin/myprofile/models"
	"github.com/pkg/errors"
)

// CommentsCreatePost default implementation.
func CommentsCreatePost(c buffalo.Context) error {
	comment := &models.Comment{}
	user := c.Value("current_user").(*models.User)
	if err := c.Bind(comment); err != nil {
		return errors.WithStack(err)
	}
	tx := c.Value("tx").(*pop.Connection)
	comment.AuthorID = user.Username
	// postID, err := uuid.FromString(c.Param("pid"))
	// if err != nil {
	// 	return errors.WithStack(err)
	// }
	comment.PostID, _ = strconv.Atoi(c.Param("pid"))
	verrs, err := tx.ValidateAndCreate(comment)
	if err != nil {
		return errors.WithStack(err)
	}
	if verrs.HasAny() {
		c.Flash().Add("danger", "There was an error adding your comment.")
		return c.Redirect(302, "/posts/detail/%s", c.Param("pid"))
	}
	c.Flash().Add("success", "Comment added successfully.")
	return c.Redirect(302, "/posts/detail/%s", c.Param("pid"))
}

// CommentsCreate default implementation.
func CommentsCreate(c buffalo.Context) error {
	return c.Render(200, r.HTML("comments/create.html"))
}

// CommentsEdit default implementation.
func CommentsEdit(c buffalo.Context) error {
	return c.Render(200, r.HTML("comments/edit.html"))
}

// CommentsDelete default implementation.
func CommentsDelete(c buffalo.Context) error {
	return c.Render(200, r.HTML("comments/delete.html"))
}
