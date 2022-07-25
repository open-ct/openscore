package token

import (
	"time"
)

// TokenClaims means a claim segment in a JWT.
type TokenClaims struct {
	Id        int64 `json:"id"`
	Role      int64 `json:"role"`
	ExpiresAt int64 `json:"expires_at"` // 过期时间（时间戳，10位）
}

// Valid checks whether the token is valid.
// Used by jwt-go package.
// Because of this method, the type TokenClaims implements the interface Claims of jwt-go.
func (c TokenClaims) Valid() error {
	now := time.Now().Unix()

	if !c.VerifyExpiresAt(now, false) {
		return ErrTokenExpired
	}

	return nil
}

// VerifyExpiresAt verifies the 'ExpiresAt' field.
// If required is false, this method will return true if the value matches or is unset.
func (c *TokenClaims) VerifyExpiresAt(now int64, required bool) bool {
	if c.ExpiresAt == 0 {
		return !required
	}
	return now <= c.ExpiresAt
}
