package sql

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/dogmatiq/projectionkit/sql/mysql"
	"github.com/dogmatiq/projectionkit/sql/postgres"
	"github.com/dogmatiq/projectionkit/sql/sqlite"
)

// Driver is an interface for database-specific projection drivers.
type Driver interface {
	// UpdateVersion updates the version for a specific handler and resource.
	UpdateVersion(
		ctx context.Context,
		tx *sql.Tx,
		h string,
		r, c, n []byte,
	) (bool, error)

	// QueryVersion returns the version for a specific handler and resource.
	QueryVersion(
		ctx context.Context,
		db *sql.DB,
		h string,
		r []byte,
	) ([]byte, error)

	// DeleteResource removes the version for a specific handler and resource.
	DeleteResource(
		ctx context.Context,
		db *sql.DB,
		h string,
		r []byte,
	) error
}

// NewDriver returns the appropriate driver implementation to use with the given
// database.
func NewDriver(db *sql.DB) (Driver, error) {
	if mysql.IsCompatibleWith(db) {
		return &mysql.Driver{}, nil
	}

	if postgres.IsCompatibleWith(db) {
		return &postgres.Driver{}, nil
	}

	if sqlite.IsCompatibleWith(db) {
		return &sqlite.Driver{}, nil
	}

	return nil, fmt.Errorf(
		"can not deduce the appropriate SQL projection driver for %T",
		db.Driver(),
	)
}