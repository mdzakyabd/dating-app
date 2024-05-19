package config

import (
	"os"

	"github.com/joho/godotenv"
)

func ConfigJWT() (string, error) {
	var err error

	mu.Lock()
	defer mu.Unlock()

	once.Do(func() {
		err = godotenv.Load()
	})

	if err != nil {
		return "", err
	}

	secret := os.Getenv("JWT_SECRET")

	return secret, nil
}
