package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// TODO: externilize key
var secretKey = []byte("secret-key")

func createToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": username,
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		})
	tokenStr, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return tokenStr, nil
}

func verifyToken(tokenStr string) (map[string]string, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		fmt.Println("er1")
		return nil, err
	}

	if !token.Valid {
		fmt.Println("er2")
		return nil, errors.New("jwt: invalid token 1")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		fmt.Println("er3")
		return nil, errors.New("jwt: unable to assert type of claims")
	}

	// Convert jwt.MapClaims (map[string]interface{}) to map[string]string for specific fields
	result := make(map[string]string)
	if username, ok := claims["username"].(string); ok {
		result["username"] = username
		fmt.Println(result["username"])
	} else {
		return nil, errors.New("jwt: username claim is not a string")
	}

	// Optionally handle other fields, converting types as necessary
	if exp, ok := claims["exp"].(float64); ok { // JWT numeric dates are decoded as float64
		result["exp"] = time.Unix(int64(exp), 0).Format(time.RFC3339) // convert to string if needed
	} else {
		return nil, errors.New("jwt: exp claim is not a valid number")
	}

	return result, nil
}
