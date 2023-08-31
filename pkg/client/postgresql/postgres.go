package postgresql

import (
	"fmt"

	_ "github.com/lib/pq"

	"github.com/jmoiron/sqlx"
)

const (
	UsersTable        = "users"
	SegmentTable      = "storage"
	UsersSegmentTable = "userssegment"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
}

func NewClient(cfg Config) (*sqlx.DB, error) {
	dataSource := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Database)
	conn, err := sqlx.Connect("postgres", dataSource)
	if err != nil {
		return nil, err
	}

	err = conn.Ping()
	if err != nil {
		return nil, err
	}

	return conn, nil

}
