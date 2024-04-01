package password

import (
	"golang.org/x/crypto/bcrypt"
)

var Encoder = NewBcryptPasswordEncoder(bcrypt.DefaultCost)

// PasswordEncoder interface defines the methods for encoding and checking passwords.
type PasswordEncoder interface {
	Encode(password string) (string, error)
	Matches(rawPassword, encodedPassword string) bool
}

// BcryptPasswordEncoder is an implementation of PasswordEncoder using the bcrypt hashing algorithm.
type BcryptPasswordEncoder struct {
	cost int
}

// NewBcryptPasswordEncoder creates a new BcryptPasswordEncoder with the given cost.
func NewBcryptPasswordEncoder(cost int) *BcryptPasswordEncoder {
	return &BcryptPasswordEncoder{
		cost: cost,
	}
}

// Encode hashes the password using bcrypt.
func (b *BcryptPasswordEncoder) Encode(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), b.cost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// Matches checks if the given raw password matches the encoded password.
func (b *BcryptPasswordEncoder) Matches(rawPassword, encodedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(encodedPassword), []byte(rawPassword))
	return err == nil
}
