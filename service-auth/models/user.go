package models

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"time"
)

// ActivateTokens contain activation tokens
type ActivateTokens struct {
	ID        uint       `gorm:"primary_key" json:"-"`
	Activated bool       `gorm:"not null;default:false" json:"-"`
	Token     string     `json:"-"`
	Email     string     `json:"-"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"-"`
}

// PasswordRestoreTokens contain password restore tokens
type PasswordRestoreTokens struct {
	ID        uint       `gorm:"primary_key" json:"-"`
	Activated bool       `gorm:"not null;default:false" json:"-"`
	Token     string     `json:"-"`
	Email     string     `json:"-"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"-"`
}

// User model
type User struct {
	ID           uint      `gorm:"primary_key" json:"id"`
	FirstName    string    `gorm:"not null" json:"first_name"`
	LastName     string    `gorm:"not null" json:"last_name"`
	Nickname     string    `gorm:"unique;not null;index:nickname" json:"nickname"`
	Email        string    `gorm:"unique;not null;index:email;type:varchar(100)" json:"email"`
	Activated    bool      `gorm:"not null;default:false" json:"-"`
	Role         string    `gorm:"not null;default:user" json:"role"`
	Sex          int       `gorm:"not null;default:2" json:"sex"` // 1 – female; 2 – male.
	BDate        time.Time `json:"b_date"`
	Picture      string    `json:"picture"`
	Desc         string    `json:"desc"`
	Status       string    `json:"status"`
	Badges       []Badges  `json:"badges"`
	PasswordHash string    `gorm:"not null" json:"-"`

	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"-"`
}

// Badges struct
type Badges struct {
	Name string `gorm:"not null"`
	Icon string `gorm:"not null"`
}

// View return user with private settings
func (u *User) View() User {
	return User{
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Nickname:  u.Nickname,
		Email:     u.Email,
		Role:      u.Role,
		Status:    u.Status,
		Badges:    u.Badges,
		Sex:       u.Sex,
		Picture:   u.Picture,
		Desc:      u.Desc,
		BDate:     u.BDate,
	}
}

// SetPassword crypt and set password to current user
func (u *User) SetPassword(password string) error {
	if len(password) == 0 {
		return errors.New("password should not be empty")
	}
	bytePassword := []byte(password)
	// Make sure the second param `bcrypt generator cost` between [4, 32)
	passwordHash, _ := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	u.PasswordHash = string(passwordHash)
	return nil
}

// CheckPassword compare current password hash and password string
// Database will only save the hashed string, you should check it by util function.
// 	if err := serModel.checkPassword("password0"); err != nil { password error }
func (u *User) CheckPassword(password string) error {
	bytePassword := []byte(password)
	byteHashedPassword := []byte(u.PasswordHash)
	return bcrypt.CompareHashAndPassword(byteHashedPassword, bytePassword)
}
