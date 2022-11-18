package main

import (
	"bytes"
	"context"
	"encoding/gob"
	"os"
	"redis_postgres_MC/models"
	"time"

	"github.com/go-redis/redis/v8"
)

type Client struct {
	client *redis.Client
}

func NewRedis() (*Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:        os.Getenv("REDIS"),
		DB:          0,
		DialTimeout: 100 * time.Millisecond,
		ReadTimeout: 100 * time.Millisecond,
	})

	if _, err := client.Ping(context.Background()).Result(); err != nil {
		return nil, err
	}

	return &Client{
		client: client,
	}, nil
}

func (c *Client) GetName(ctx context.Context, nconst string) (models.Name, error) {
	cmd := c.client.Get(ctx, nconst)

	cmbd, err := cmd.Bytes()
	if err != nil {
		return models.Name{}, err
	}

	b := bytes.NewReader(cmbd)

	var result models.Name

	if err := gob.NewDecoder(b).Decode(&result); err != nil {
		return models.Name{}, err
	}

	return result, nil
}

func (c *Client) SetName(ctx context.Context, n models.Name) error {
	var b bytes.Buffer

	if err := gob.NewEncoder(&b).Encode(n); err != nil {
		return err
	}

	return c.client.Set(ctx, n.NConst, b.Bytes(), 25*time.Second).Err()
}
