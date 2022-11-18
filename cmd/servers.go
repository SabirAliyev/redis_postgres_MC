package main

import "log"

func InitializePostgres() *PostgreSQL {
	db, err := NewPostgreSQL()
	if err != nil {
		log.Fatalf("Could not initialize database connection %s", err)
	}

	return db
}

func InitializeRedis() *Client {
	redis, err := NewRedis()
	if err != nil {
		log.Fatalf("Could not initialize Redis client %s", err)
	}

	return redis
}
