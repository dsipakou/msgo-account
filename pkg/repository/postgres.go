package repository

import (
  "fmt"
  "log"
  "github.com/jmoiron/sqlx"
   _ "github.com/lib/pq"
)

type Config struct {
  Host      string
  Port      string
  Username  string
  Password  string
  DBName    string
  SSLMode   string
}

func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
  connStr := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode)
  log.Println(connStr)
  db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
    cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))

  if err != nil {
    return nil, err
  }

  err1 := db.Ping()
  if err1 != nil {
    return nil, err1
  }

  log.Println("Connected to DB")
  return db, nil
}
