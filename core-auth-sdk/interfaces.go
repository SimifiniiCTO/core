package core_auth_sdk

import (
	jose "gopkg.in/square/go-jose.v2"
	jwt "gopkg.in/square/go-jose.v2/jwt"
)

// JWKProvider Provides a JSON Web Key from a Key ID
// Wanted to use function signature from go-jose.v2
// but that would make us lose error information
type JWKProvider interface {
	Key(kid string) ([]jose.JSONWebKey, error)
}

// JWTClaimsExtractor Extracts verified in-built claims from a jwt idToken
type JWTClaimsExtractor interface {
	GetVerifiedClaims(idToken string) (*jwt.Claims, error)
}

// AuthService exposes the interface contract the authentication service client adheres to
type AuthService interface {
	// GetAccount Get a user account
	GetAccount(id string) (*Account, error)
	// Update Updates the username associated with a user account
	Update(id, username string) error
	// LockAccount Locks a user account
	LockAccount(id string) error
	// UnlockAccount Unlocks a user account
	UnlockAccount(id string) error
	// ArchiveAccount Archives a user account
	ArchiveAccount(id string) error
	// ImportAccount Creates a new user account
	ImportAccount(username, password string, locked bool) (int, error)
	// ExpirePassword Expires the password associated with a user account
	ExpirePassword(id string) error
	// LoginAccount Authenticates a user account
	LoginAccount(username, password string) (string, error)
	// SignupAccount Signs up a user account
	SignupAccount(username, password string) (string, error)
	// LogOutAccount Remove a session associated with a given user account
	LogOutAccount() error
}
