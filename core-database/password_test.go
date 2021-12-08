package core_database_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const passwordtoHash string = "teststring"
const emptyPasswordToHash string = ""

func TestValidateAndHashPassword(t *testing.T) {
	t.Run("TestName:Passed_TestValidateAndHashPassword", ValidateAndHashPasswordValidPassword)
	t.Run("TestName:Failed_TestInValidateAndHashPassword", ValidateAndHashPasswordInValidPassword)
}

// TestComparePassword tests whether a password can be adequately compared
func TestComparePassword(t *testing.T) {
	t.Run("TestName:ComparePasswords", ComparePasswords)
}

// ValidateAndHashPasswordValidPassword Tests a valid password
func ValidateAndHashPasswordValidPassword(t *testing.T) {
	// act
	hashPassword, err := dbConn.ValidateAndHashPassword(passwordtoHash)

	// assert
	assert.Empty(t, err)
	assert.NotEmpty(t, hashPassword)
}

// ValidateAndHashPasswordInValidPassword Tests wether a given invalid hashed password
func ValidateAndHashPasswordInValidPassword(t *testing.T) {
	// act
	hashPassword, err := dbConn.ValidateAndHashPassword(emptyPasswordToHash)

	// assert
	assert.NotEmpty(t, err)
	assert.Empty(t, hashPassword)
}

// ComparePasswords Tests if a hashed password is equal to a plain counterpart
func ComparePasswords(t *testing.T) {
	// act
	hashPassword, err := dbConn.ValidateAndHashPassword(passwordtoHash)
	// assert
	assert.Empty(t, err)
	assert.NotEmpty(t, hashPassword)

	valid := dbConn.ComparePasswords(hashPassword, []byte(passwordtoHash))
	assert.True(t, valid)

	valid = dbConn.ComparePasswords(hashPassword, []byte("random string for password"))
	assert.False(t, valid)
}
