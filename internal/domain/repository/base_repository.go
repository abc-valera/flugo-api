package repository

// BaseRepository is a set of functions that every repository contains
type BaseRepository interface {
	// WithTx starts txFunc in a transaction
	WithTx(txFunc func() error) error
}
