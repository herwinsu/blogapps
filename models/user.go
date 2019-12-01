package models

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

// User represents user data.
type User struct {
	ID              int       `db:"id" rw:"r"`
	Name            string    `db:"name" form:"name"`
	Username        string    `db:"username" form:"username"`
	Email           string    `db:"email" form:"email"`
	RoleID          int       `db:"roleid" form:"roleid"`
	Password        string    `db:"password"`
	CreatedAt       time.Time `db:"created_at"`
	UpdatedAt       time.Time `db:"updated_at"`
	PasswordHash    string    `db:"-"`
	PasswordConfirm string    `db:"-"`
}

// TableName overrides the table name used by Pop.
func (u User) TableName() string {
	return "xx_user_h"
}

// String is not required by pop and may be deleted
func (u User) String() string {
	ju, _ := json.Marshal(u)
	return string(ju)
}

// Users is not required by pop and may be deleted
type Users []User

// String is not required by pop and may be deleted
func (u Users) String() string {
	ju, _ := json.Marshal(u)
	return string(ju)
}

// Create validates and creates a new User.
func (u *User) Create(tx *pop.Connection) (*validate.Errors, error) {
	u.Email = strings.ToLower(u.Email)
	pwdHash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return validate.NewErrors(), errors.WithStack(err)
	}
	u.PasswordHash = u.Password
	u.Password = string(pwdHash)

	verrs, _ := u.Validate(tx)
	if verrs.HasAny() {
		return verrs, nil
	}

	// email := TableField{Name: "Email", Value: u.Email}
	err = u.EmailIsNotTaken(tx)

	if err == nil {
		verrs.Add(validators.GenerateKey("Email"), fmt.Sprintf("The email %s is not available.", u.Email))

		return verrs, nil
	}

	// usrname := TableField{Name: "Username", Value: u.Username}
	err = u.UserNameIsNotTaken(tx)

	if err == nil {
		verrs.Add(validators.GenerateKey("Username"), fmt.Sprintf("The username %s is not available.", u.Username))

		return verrs, nil
	}

	err = tx.Create(u)
	if err != nil {
		return validate.NewErrors(), errors.WithStack(err)
	}

	return verrs, err //tx.Create(u)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (u *User) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: u.Name, Name: "Name"},
		&validators.StringIsPresent{Field: u.Username, Name: "Username"},
		&validators.StringIsPresent{Field: u.Email, Name: "Email"},
		&validators.EmailIsPresent{Name: "Email", Field: u.Email},
		&validators.StringIsPresent{Field: u.Username, Name: "Username"},
		&validators.StringIsPresent{Field: u.Password, Name: "Password"},
		&validators.StringsMatch{Name: "Password", Field: u.PasswordHash, Field2: u.PasswordConfirm, Message: "Passwords do not match."},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (u *User) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (u *User) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// EmailIsNotTaken gets run every time you call "EmailIsNotTaken(tx)" method.
func (u *User) EmailIsNotTaken(tx *pop.Connection) error {
	query := tx.Where("email = ?", u.Email)
	queryUser := User{}
	err := query.First(&queryUser)
	return err
}

// UserNameIsNotTaken gets run every time you call "UserNameIsNotTaken(tx)" method.
func (u *User) UserNameIsNotTaken(tx *pop.Connection) error {
	query := tx.Where("username = ?", u.Username)
	queryUser := User{}
	err := query.First(&queryUser)
	return err
}

// FindByUsername gets run every time you call "UserNameIsNotTaken(tx)" method.
func (u *User) FindByUsername(tx *pop.Connection) error {
	query := tx.Where("username = ?", u.Username)
	queryUser := User{}
	err := query.First(&queryUser)
	return err
}

// Authorize checks user's password for logging in
func (u *User) Authorize(tx *pop.Connection) error {
	password := u.Password
	err := tx.Where("email = ?", strings.ToLower(u.Email)).First(u)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			// couldn't find an user with that email address
			return errors.New("User not found")
		}
		return errors.WithStack(err)
	}
	// confirm that the given password matches the hashed password from the db
	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		return errors.New("Invalid Password")
	}
	return nil
}
