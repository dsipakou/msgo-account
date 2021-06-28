package db

import "database/sql"
import _ "github.com/go-sql-driver/postgres"

type Transaction struct {
  UserId int
  Category string
  Amount int
}
