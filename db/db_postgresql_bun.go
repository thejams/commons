package db

import (
	"database/sql"
	"fmt"
	"sync"
	"time"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

var once sync.Once

type PostgreSQL struct {
	DB *bun.DB
}

// NewPostgreSQLDB returns a PostgreSQL db object
func NewPostgreSQLDB() IDB {
	return &PostgreSQL{}
}

// Conn connects to a postgresql database using Bun ORM
func (p *PostgreSQL) Conn(s DBSettings) error {

	once.Do(func() {
		dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", s.DBUser, s.DBPassword, s.DBAddress, s.DBPort, s.DBName)

		pgconn := pgdriver.NewConnector(
			pgdriver.WithDSN(dsn),
			pgdriver.WithTimeout(5*time.Second),
			pgdriver.WithDialTimeout(5*time.Second),
			pgdriver.WithReadTimeout(5*time.Second),
			pgdriver.WithWriteTimeout(5*time.Second),
			//pgdriver.WithTLSConfig(&tls.Config{InsecureSkipVerify: true}),
			/* pgdriver.WithConnParams(map[string]interface{}{
				"search_path": "my_search_path",
			}), */
		)

		// db, err := sql.Open("postgres", "postgres://{user}:{password}@{hostname}:{port}/{database-name}?sslmode=disable")
		sqldb := sql.OpenDB(pgconn)
		sqldb.SetMaxOpenConns(100)
		sqldb.SetConnMaxIdleTime(10 * time.Second)
		sqldb.SetMaxIdleConns(2)

		p.DB = bun.NewDB(sqldb, pgdialect.New())
	})

	return p.DB.Ping()
}

// Close closes the postgresql database connection
func (p *PostgreSQL) Close() error {
	return p.DB.Close()
}
