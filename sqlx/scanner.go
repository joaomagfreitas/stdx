package sqlx

import (
	"database/sql"
)

// Scan maps a single sql.Row into a [Record] struct.
// The mapper must return pointers to the fields of Record for passing to row.Scan.
//
// The whole point of this function is to provide a helper function that reduces boilerplate code like this:
//
//	var r User
//	err := row.Scan(&r.ID, &r.Name, &r.Email)
//	if err != nil {
//		return nil, err
//	}
//	return &r, nil
//
// Into:
//
//	user, err := Scan(row, func(u *User) []any {
//		return []any{&u.ID, &u.Name, &u.Email}
//	})
func Scan[Record any](row *sql.Row, destination func(r *Record) []any) (*Record, error) {
	var r Record
	return &r, row.Scan(destination(&r)...)
}

// ScanAll maps all sql.Rows into a slice of [Record] structs.
// The mapper must return pointers to the fields of Record for passing to row.Scan.
//
// The whole point of this function is to provide a helper function that reduces boilerplate code like this:
//
//	var users []*User
//	for rows.Next() {
//		var u User
//		if err := rows.Scan(&u.ID, &u.Name, &u.Email); err != nil {
//			return nil, err
//		}
//		users = append(users, &u)
//	}
//	return users, nil
//
// Into:
//
//	users, err := ScanAll(rows, func(u *User) []any {
//		return []any{&u.ID, &u.Name, &u.Email}
//	})
func ScanAll[Record any](rows *sql.Rows, destination func(r *Record) []any) ([]*Record, error) {
	rs := []*Record{}
	for rows.Next() {
		var r Record
		if err := rows.Scan(destination(&r)...); err != nil {
			return nil, err
		}

		rs = append(rs, &r)
	}

	return rs, nil
}
