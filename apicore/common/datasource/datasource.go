// Package datasource provides access to the database for the settings module.
package datasource

import (
	"context"

	"encore.dev/beta/errs"
	"encore.dev/rlog"
	"encore.dev/storage/sqldb"
	"github.com/jackc/pgx/v5/pgxpool"

	"encore.app/apicore/common/datasource/db"
)

// UnitOfWork defines a function type that takes a pointer to a transactional
// queries object and returns an error. It is used to execute a series of
// database operations within a transaction.
type UnitOfWork func(qtx *db.Queries) error

// IDatasource defines the interface for the settings module datasource.
//
// It can contain SQLC generated queries for settings table but also for any
// other table used in the module. Isolation and segregation of responsabilities
// should be achieved on repositories.
type IDatasource interface {
	Database() *pgxpool.Pool
	Queries() *db.Queries
	QueriesTx(ctx context.Context, fn UnitOfWork) error
}

// Datasource provides access to the settings database.
type Datasource struct {
	database *pgxpool.Pool
	queries  *db.Queries
}

// New creates a new instance of Datasource.
func New(encoreSQLDb *sqldb.Database) *Datasource {
	pgxdb := sqldb.Driver[*pgxpool.Pool](encoreSQLDb)
	queries := db.New(pgxdb)

	return &Datasource{
		database: pgxdb,
		queries:  queries,
	}
}

// Database returns the database connection.
func (d *Datasource) Database() *pgxpool.Pool {
	return d.database
}

// Queries returns the queries for the settings datasource.
func (d *Datasource) Queries() *db.Queries {
	return d.queries
}

// QueriesTx begins and manages a database transaction, providing a pointer to a
// transactional queries object to the unit of work. It ensures that the transaction
// is either committed if the unit of work succeeds or rolled back if an error occurs.
//
// Parameters:
//   - ctx: The context for managing the transaction lifecycle.
//   - uow: The unit of work function that contains the database operations to be executed.
//
// Returns:
//   - An error if the transaction could not be started, if the unit of work fails,
//     or if the transaction could not be committed.
func (d *Datasource) QueriesTx(ctx context.Context, uow UnitOfWork) error {
	tx, err := d.database.Begin(ctx)

	if err != nil {
		rlog.Error("Unit of Work could not start transaction", "error", err)

		return errs.WrapCode(err, errs.Unavailable, "transaction_not_started")
	}

	// Ensure transaction is rolled back if the function exits before the
	// transaction is explicitly committed. This is a safety mechanism to
	// prevent leaving a transaction open, which could lead to database locks.
	defer tx.Rollback(ctx)

	// Creae ntew queries instance with transaction
	qtx := d.queries.WithTx(tx)

	// Execute unit of work callback. In case it errs rollback any changes made
	// to the database because this function is exiting before explicitly
	// calling tx.Commit.
	if err := uow(qtx); err != nil {
		return errs.WrapCode(err, errs.Internal, "unit_of_work_not_completed")
	}

	// Commit changes if no errors occurred
	if err := tx.Commit(ctx); err != nil {
		rlog.Error("Unit of work failed to commit transaction", "error", err)

		return errs.WrapCode(err, errs.Unavailable, "transaction_not_committed")
	}

	return nil
}
