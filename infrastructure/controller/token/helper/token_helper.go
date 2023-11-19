package helper

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

func GenerateRefreshToken(secretKey []byte) (string, error) {
	// Create a new claims object
	claims := jwt.MapClaims{}

	// Set the claims
	claims["exp"] = time.Now().Add(time.Hour * 24 * 7).Unix()

	// Create a new token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	refreshToken, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return refreshToken, nil
}

func CheckRefreshToken(refreshToken string) (bool, error) {
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(refreshToken, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("your-secret-key"), nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return false, nil
		}
		return false, err
	}

	// Check if the token has expired
	exp := int64(claims["exp"].(float64))
	if time.Now().Unix() > exp {
		return false, nil
	}

	return true, nil
}

func GenerateAccessToken(refreshToken string, secretKey []byte) (string, error) {
	// Parse the refresh token
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(refreshToken, claims, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return "", err
	}

	// Create a new claims object
	claims = jwt.MapClaims{}

	// Set the claims
	claims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	claims["refresh_token"] = refreshToken

	// Create a new token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	accessToken, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return accessToken, nil
}

func CheckAccessToken(accessToken string) (bool, error) {
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(accessToken, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("your-secret-key"), nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return false, nil
		}
		return false, err
	}

	// Check if the token has expired
	exp := int64(claims["exp"].(float64))
	if time.Now().Unix() > exp {
		return false, nil
	}

	return true, nil
}
