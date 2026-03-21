package sqlx_test

import (
	"database/sql"
	"reflect"
	"testing"

	"github.com/joaomagfreitas/stdx/slicesx"
	"github.com/joaomagfreitas/stdx/sqlx"
)

func TestScan(t *testing.T) {
	db, err := sql.Open(testDriver, "")
	if err != nil {
		t.Fatal(err)
	}

	defer db.Close()

	row := db.QueryRow(query)
	foo, err := sqlx.Scan(row, func(f *Foo) []any {
		return []any{&f.Id, &f.Bar}
	})

	if !reflect.DeepEqual(
		*foo,
		Foo{Id: 1, Bar: "Foo"},
	) {
		t.Fatal(foo)
	}
}

func TestScanAll(t *testing.T) {
	db, err := sql.Open(testDriver, "")
	if err != nil {
		t.Fatal(err)
	}

	defer db.Close()

	rows, err := db.Query(query)
	if err != nil {
		t.Fatal(err)
	}

	foos, err := sqlx.ScanAll(rows, func(f *Foo) []any {
		return []any{&f.Id, &f.Bar}
	})

	if !reflect.DeepEqual(
		slicesx.Map(foos, func(foo *Foo) Foo { return *foo }),
		[]Foo{{Id: 1, Bar: "Foo"}, {Id: 2, Bar: "Bar"}, {Id: 3, Bar: "Baz"}},
	) {
		t.Fatal(foos)
	}
}
