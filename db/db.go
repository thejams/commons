package db

type IDB interface {
	// Conn connects to a DB engine
	Conn(settings DBSettings) error

	// Close closes the DB connection
	Close() error
}
