package token

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var (
	jwtKey string

	ErrTokenInvalid = errors.New("the token is invalid")
	ErrTokenExpired = errors.New("the token is expired")
)

// getJwtKey get jwtKey.
func getJwtKey() string {
	if jwtKey == "" {
		jwtKey = "jwt_secret_openscore"
	}
	return jwtKey
}

// TokenPayload is a required payload when generates token.
type TokenPayload struct {
	Id      int64         `json:"id"`
	Role    int64         `json:"role"`
	Expired time.Duration `json:"expired"` // 有效时间（nanosecond）
}

// TokenResolve means returned payload when resolves token.
type TokenResolve struct {
	Id        int64 `json:"id"`
	Role      int64 `json:"role"`
	ExpiresAt int64 `json:"expires_at"` // 过期时间（时间戳，10位）
}

// GenerateToken generates token.
func GenerateToken(payload *TokenPayload) (string, error) {
	claims := &TokenClaims{
		Id:        payload.Id,
		Role:      payload.Role,
		ExpiresAt: time.Now().Unix() + int64(payload.Expired.Seconds()),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(getJwtKey()))
}

// ResolveToken resolves token.
func ResolveToken(tokenStr string) (*TokenResolve, error) {
	claims := &TokenClaims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(getJwtKey()), nil
	})

	if err != nil {
		log.Println("Token parsing failed because of an internal error", err)
		return nil, err
	}

	if !token.Valid {
		log.Println("Token parsing failed; the token is invalid.", err)
		return nil, ErrTokenInvalid
	}

	t := &TokenResolve{
		Id:        claims.Id,
		Role:      claims.Role,
		ExpiresAt: claims.ExpiresAt,
	}
	return t, nil
}

// GetExpiredTime get token expired time from env or config file.
func GetExpiredTime() time.Duration {
	day := 10
	return time.Hour * 24 * time.Duration(day)
}
