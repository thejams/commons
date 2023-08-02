package db

import "time"

type DBSettings struct {
	DBAddress   string
	DBUser      string
	DBPassword  string
	DBName      string
	DBPort      string
	DbPagLimit  int
	PoolSize    int
	PoolTimeout time.Duration
	IdleTimeout time.Duration
	MaxConnAge  time.Duration
}
