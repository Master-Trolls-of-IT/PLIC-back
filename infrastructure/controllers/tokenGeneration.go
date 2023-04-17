package controllers

import (
	_ "github.com/golang-jwt/jwt/v4"
	"time"
)

// Define a function to generate a new refresh token with a secret key
func generateRefreshToken(secretKey []byte) (string, error) {
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

// Define a function to generate a new access token that depends on the refresh token and a secret key
func generateAccessToken(refreshToken string, secretKey []byte) (string, error) {
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

// Define a function to verify if an access token is still valid
func verifyAccessToken(accessToken string) (bool, error) {
	// Parse the access token
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

// Define a function to verify if a refresh token is still valid
func verifyRefreshToken(refreshToken string) (bool, error) {
	// Parse the refresh token
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
