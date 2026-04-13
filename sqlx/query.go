package sqlx

import (
	"context"
	"database/sql"
)

// Single executes the query and scans a single row into a value of type [Record].
//
// It is a thin wrapper around [QueryRowContext] and [Scan]. The dest function
// must return pointers to the fields of Record for passing to row.Scan.
//
// If the query returns no rows, the returned error will be sql.ErrNoRows.
//
// Example:
//
//	user, err := dbutil.Single[User](
//		ctx,
//		db,
//		`SELECT id, name FROM users WHERE id = ?`,
//		func(u *User) []any {
//			return []any{&u.ID, &u.Name}
//		},
//		42,
//	)
//	if err != nil {
//		if errors.Is(err, sql.ErrNoRows) {
//			// handle not found
//		}
//		return err
//	}
//
//	fmt.Println(user.Name)
func Single[Record any](
	ctx context.Context,
	db *sql.DB,
	query string,
	dest func(record *Record) []any,
	args ...any,
) (Record, error) {
	return Scan(db.QueryRowContext(ctx, query, args...), dest)
}

// SingleTx executes the query in a transaction and scans a single row into a value of type [Record].
//
// It is a thin wrapper around [QueryRowContext] and [Scan]. The dest function
// must return pointers to the fields of Record for passing to row.Scan.
//
// If the query returns no rows, the returned error will be sql.ErrNoRows.
//
// Example:
//
//	user, err := dbutil.SingleTx[User](
//		ctx,
//		db,
//		`SELECT id, name FROM users WHERE id = ?`,
//		func(u *User) []any {
//			return []any{&u.ID, &u.Name}
//		},
//		42,
//	)
//	if err != nil {
//		if errors.Is(err, sql.ErrNoRows) {
//			// handle not found
//		}
//		return err
//	}
//
//	fmt.Println(user.Name)
func SingleTx[Record any](
	ctx context.Context,
	tx *sql.Tx,
	query string,
	dest func(record *Record) []any,
	args ...any,
) (Record, error) {
	return Scan(tx.QueryRowContext(ctx, query, args...), dest)
}

// All executes the query and scans a all rows into values of type [Record].
//
// It is a thin wrapper around [QueryRowsContext] and [ScanAll]. The dest function
// must return pointers to the fields of Record for passing to row.Scan.
//
// If the query returns no rows, the returned error will be sql.ErrNoRows.
//
// Example:
//
//	users, err := dbutil.All[User](
//		ctx,
//		db,
//		`SELECT id, name FROM users WHERE id = ? LIMIT 5`,
//		func(u *User) []any {
//			return []any{&u.ID, &u.Name}
//		},
//		42,
//	)
//	if err != nil {
//		if errors.Is(err, sql.ErrNoRows) {
//			// handle not found
//		}
//		return err
//	}
//
//	fmt.Println(len(users))
func All[R any](
	ctx context.Context,
	db *sql.DB,
	query string,
	dest func(record *R) []any,
	args ...any,
) ([]R, error) {
	rows, err := db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	return ScanAll(rows, dest)
}

// AllTx executes the query in a transaction and scans a all rows into values of type [Record].
//
// It is a thin wrapper around [QueryRowsContext] and [ScanAll]. The dest function
// must return pointers to the fields of Record for passing to row.Scan.
//
// If the query returns no rows, the returned error will be sql.ErrNoRows.
//
// Example:
//
//	users, err := dbutil.AllTx[User](
//		ctx,
//		tx,
//		`SELECT id, name FROM users WHERE id = ? LIMIT 5`,
//		func(u *User) []any {
//			return []any{&u.ID, &u.Name}
//		},
//		42,
//	)
//	if err != nil {
//		if errors.Is(err, sql.ErrNoRows) {
//			// handle not found
//		}
//		return err
//	}
//
//	fmt.Println(len(users))
func AllTx[R any](
	ctx context.Context,
	tx *sql.Tx,
	query string,
	dest func(record *R) []any,
	args ...any,
) ([]R, error) {
	rows, err := tx.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	return ScanAll(rows, dest)
}
