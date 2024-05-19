package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/pusher/pusher-http-go"
)

func ConfigPusher() (*pusher.Client, error) {
	var err error

	mu.Lock()
	defer mu.Unlock()

	once.Do(func() {
		err = godotenv.Load()
	})

	if err != nil {
		return nil, err
	}

	p := &pusher.Client{
		AppID:   os.Getenv("PUSHER_APP_ID"),
		Key:     os.Getenv("PUSHER_KEY"),
		Secret:  os.Getenv("PUSHER_SECRET"),
		Cluster: os.Getenv("PUSHER_CLUSTER"),
		Secure:  true,
	}

	return p, nil
}
