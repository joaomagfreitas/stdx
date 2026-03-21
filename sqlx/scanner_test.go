package sqlx_test

import (
	"database/sql"
	"errors"
	"reflect"
	"testing"

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

	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(
		foo,
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

	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(
		foos,
		[]Foo{{Id: 1, Bar: "Foo"}, {Id: 2, Bar: "Bar"}, {Id: 3, Bar: "Baz"}},
	) {
		t.Fatal(foos)
	}
}

func TestScanAllErrorResetsResult(t *testing.T) {
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
		return []any{"", &f.Bar}
	})

	if err == nil {
		t.FailNow()
	}

	if foos != nil {
		t.Fatal(foos)
	}
}

func TestScanAllCloseErrorIsNotIgnored(t *testing.T) {
	db, err := sql.Open(testDriver, "")
	if err != nil {
		t.Fatal(err)
	}

	defer db.Close()

	rows, err := db.Query(query, errRowsClose.Error())
	if err != nil {
		t.Fatal(err)
	}

	_, serr := sqlx.ScanAll(rows, func(f *Foo) []any {
		return []any{&f.Id, &f.Bar}
	})

	if serr == nil {
		t.FailNow()
	}

	if !errors.Is(serr, errRowsClose) {
		t.Fatal(serr)
	}
}
