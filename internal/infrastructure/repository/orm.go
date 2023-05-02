package repository

import (
	"context"
	"database/sql"

	"github.com/abc-valera/flugo-api/internal/domain"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

// Contains boilerplate database queries (and error handling?)

// handlePGErr transfers pg error to domain.Error type
//
// Return Codes:
//   - AlreadyExists (if 23505 error)
//   - NotFound (if sql.ErrNoRows)
//   - Internal (all other errors)
func handlePGErr(err error, operation string) error {
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

// baseExecDB is used for INSERT queries
func baseExecDB(
	c context.Context,
	db *sqlx.DB,
	query,
	op string,
) error {
	_, err := db.ExecContext(c, query)
	return handlePGErr(err, op+" repository.CreateEntity")
}

// TODO: error validation!
// TODO: building queries with parameters ($1, $2, ...)???

// execCheckDB is usually used for UPDATE and DELETE queries.
// It performes check for the number of affected rows.
//
// Returned codes:
//   - NotFound
//   - Internal
func execCheckDB(
	c context.Context,
	db *sqlx.DB,
	query,
	op string,
) error {
	res, err := db.ExecContext(c, query)
	if err != nil {
		return domain.NewInternalError(op+" repository.updateEntityField", err)
	}
	affected, err := res.RowsAffected()
	if err != nil || affected == 0 {
		return &domain.Error{Code: domain.CodeNotFound}
	}
	return nil
}

// execWithCheckDB is used for SELECT queries, where only one instanse is returned!
// Recommended to use queries with LIMIT(1)
func getDB(
	c context.Context,
	db *sqlx.DB,
	data interface{},
	query,
	op string,
) (interface{}, error) {
	err := db.GetContext(c, data, query)
	return data, handlePGErr(err, op+" repository.getEntityByField")
}

// execWithCheckDB is used for SELECT queries, where many instanses are returned
func selectDB(
	c context.Context,
	db *sqlx.DB,
	data interface{},
	query,
	op string,
) (interface{}, error) {
	err := db.SelectContext(c, data, query)
	return data, handlePGErr(err, op+" repository.getEntities")
}
