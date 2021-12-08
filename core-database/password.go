package core_database

import (
	"errors"
	"log"

	"golang.org/x/crypto/bcrypt"
)

// ValidateAndHashPassword validates, hashes and salts a password
func (db *DatabaseConn) ValidateAndHashPassword(password string) (string, error) {
	// check if confirmed password is not empty
	if password == "" {
		return "", errors.New("password cannot be empty")
	}

	//  hash and salt password
	hashedPassword, err := db.hashAndSalt([]byte(password))
	if err != nil {
		return "", err
	}
	return hashedPassword, nil
}

// hashAndSalt hashes and salts a password
func (db *DatabaseConn) hashAndSalt(pwd []byte) (string, error) {

	// Use GenerateFromPassword to hash & salt pwd
	// MinCost is just an integer constant provided by the bcrypt
	// package along with DefaultCost & MaxCost.
	// The cost can be any value you want provided it isn't lower
	// than the MinCost (4)
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	// GenerateFromPassword returns a byte slice so we need to
	// convert the bytes to a string and return it
	return string(hash), nil
}

// ComparePasswords compares a hashed password and a plaintext password and returns
// a boolean stating wether they are equal or not
func (db *DatabaseConn) ComparePasswords(hashedPwd string, plainPwd []byte) bool {
	// Since we'll be getting the hashed password from the DB it
	// will be a string so we'll need to convert it to a byte slice
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
