package token

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type PayloadToken struct {
	AuthId  int
	Expired time.Time
}

// SecretKey is secret key to generate token
const SecretKey = "secret"

func GenerateToken(tok *PayloadToken) (string, error) {

	// set expired 10 minute from now
	tok.Expired = time.Now().Add(10 * 60 * time.Second)

	// create map claims
	claims := jwt.MapClaims{
		"payload": tok,
	}

	// create new jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// sign token with secret key and get string
	tokenString, err := token.SignedString([]byte(SecretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateToken(tokenString string) (*PayloadToken, error) {
	// Parse the token string using the provided secret key
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Ensure the signing method is HMAC
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		// Return the secret key for validation

		return []byte(SecretKey), nil
	})

	// Check for parsing errors
	if err != nil {
		return nil, err
	}

	// Assert the token claims as a MapClaims type
	claims, ok := token.Claims.(jwt.MapClaims)
	// Verify the token is valid and the claims are of the correct type
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	// Extract the payload from the claims
	payload := claims["payload"]
	var payloadToken = PayloadToken{}
	payloadByte, _ := json.Marshal(payload)
	err = json.Unmarshal(payloadByte, &payloadToken)
	if err != nil {
		return nil, err
	}

	// Return the parsed payload token
	return &payloadToken, nil

}
