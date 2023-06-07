package orm

import (
	"context"
	"database/sql"

	"github.com/abc-valera/flugo-api/internal/domain"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

// TODO: error validation!
// TODO: building queries with parameters ($1, $2, ...)???

type Queriers struct {
	GetQuerier
	ExecQuerier
	Transactioner
}

func NewQueriers(db *sqlx.DB) Queriers {
	return Queriers{
		&dbQuerier{db},
		&dbQuerier{db},
		NewTransactioner(db),
	}
}

// GetQuerier represents ways to get entities from database
type GetQuerier interface {
	// Get is used for SELECT queries, where only one instanse is returned!
	// Recommended to use queries with LIMIT(1)
	Get(c context.Context, data interface{}, query, op string) error

	// Select is used for SELECT queries, where many instanses are returned
	Select(c context.Context, data interface{}, query, op string) error
}

// ExecQuerier represents ways to execute queries to database
type ExecQuerier interface {
	// Exec executes the query (usually used for INSERT queries)
	Exec(c context.Context, query, op string) error

	// CheckExec executes the query and performes check for the number of affected rows
	// (usually used for UPDATE and DELETE queries)
	CheckExec(c context.Context, query, op string) error
}

type dbQuerier struct {
	db *sqlx.DB
}

func (db *dbQuerier) Get(c context.Context, data interface{}, query, op string) error {
	err := db.db.GetContext(c, data, query)
	return HandlePGErr(err, op+" dbQuerier.Get")
}

func (db *dbQuerier) Select(c context.Context, data interface{}, query, op string) error {
	err := db.db.SelectContext(c, data, query)
	return HandlePGErr(err, op+" dbQuerier.Select")
}

func (db *dbQuerier) Exec(c context.Context, query, op string) error {
	_, err := db.db.ExecContext(c, query)
	return HandlePGErr(err, op+" dbQuerier.Exec")
}

func (db *dbQuerier) CheckExec(c context.Context, query, op string) error {
	res, err := db.db.ExecContext(c, query)
	if err != nil {
		return domain.NewInternalError(op+" dbQuerier.CheckExec", err)
	}
	affected, err := res.RowsAffected()
	if err != nil || affected == 0 {
		return &domain.Error{Code: domain.CodeNotFound}
	}
	return nil
}

type txQuerier struct {
	tx *sqlx.Tx
}

func (tx *txQuerier) Exec(c context.Context, query, op string) error {
	_, err := tx.tx.ExecContext(c, query)
	return HandlePGErr(err, op+" txQuerier.Exec")
}

func (tx *txQuerier) CheckExec(c context.Context, query, op string) error {
	res, err := tx.tx.ExecContext(c, query)
	if err != nil {
		return domain.NewInternalError(op+" txQuerier.CheckExec", err)
	}
	affected, err := res.RowsAffected()
	if err != nil || affected == 0 {
		return &domain.Error{Code: domain.CodeNotFound}
	}
	return nil
}

type Transactioner interface {
	StartTx() (ExecQuerier, error)
	PerformTx(tFunc func() error) (ExecQuerier, error)
}

type transactioner struct {
	db *sqlx.DB
	tx *sqlx.Tx
}

func NewTransactioner(db *sqlx.DB) Transactioner {
	return &transactioner{
		db: db,
		tx: nil,
	}
}

func (t *transactioner) StartTx() (ExecQuerier, error) {
	tx, err := t.db.Beginx()
	if err != nil {
		return nil, err
	}
	t.tx = tx
	return &txQuerier{tx}, nil
}

func (t *transactioner) PerformTx(tFunc func() error) (ExecQuerier, error) {
	if err := tFunc(); err != nil {
		if eRollback := t.tx.Rollback(); eRollback != nil {
			return nil, domain.NewInternalError("orm.PerformTx", eRollback)
		}
		return nil, err
	}

	if eCommit := t.tx.Commit(); eCommit != nil {
		return nil, domain.NewInternalError("orm.PerformTx", eCommit)
	}

	return &dbQuerier{t.db}, nil
}

// HandlePGErr transfers pg error to domain.Error type
//
// Return Codes:
//   - AlreadyExists (if 23505 error)
//   - NotFound (if sql.ErrNoRows)
//   - Internal (all other errors)
func HandlePGErr(err error, operation string) error {
	if err != nil {
		if pgErr, ok := err.(*pq.Error); ok {
			// 23505 - unique_violation
			if pgErr.Code == "23505" {
				return &domain.Error{Code: domain.CodeAlreadyExists}
			}
		}
		if err == sql.ErrNoRows {
			return &domain.Error{Code: domain.CodeNotFound}
		}
		if _, ok := err.(*domain.Error); ok {
			return err
		}
		return &domain.Error{Code: domain.CodeInternal, Op: operation, Err: err}
	}
	return nil
}
