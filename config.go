package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
	Charset  string
}

func (c *Config) DSN() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=require",
		c.Host, c.Port, c.User, c.Password, c.Database)
}

func LoadConfig() *Config {
	config := &Config{
		Host:     "localhost",
		Port:     "5432",
		User:     "postgres",
		Password: "postgres",
		Database: "volleyball",
	}
	// Можно использовать окружение
	if user := os.Getenv("DB_USER"); user != "" {
		config.User = user
	}
	if pass := os.Getenv("DB_PASSWORD"); pass != "" {
		config.Password = pass
	}
	if db := os.Getenv("DB_NAME"); db != "" {
		config.Database = db
	}
	return config
}

func OpenDB() *sql.DB {
	config := LoadConfig()
	db, err := sql.Open("postgres", config.DSN())
	if err != nil {
		panic(err)
	}
	return db
}
