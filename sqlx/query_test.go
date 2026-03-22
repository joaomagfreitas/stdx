package sqlx_test

import (
	"database/sql"
	"reflect"
	"testing"

	"github.com/joaomagfreitas/stdx/sqlx"
)

func TestSingle(t *testing.T) {
	db, err := sql.Open(testDriver, "")
	if err != nil {
		t.Fatal(err)
	}

	defer db.Close()

	foo, err := sqlx.Single(t.Context(), db, query, func(f *Foo) []any {
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

func TestAll(t *testing.T) {
	db, err := sql.Open(testDriver, "")
	if err != nil {
		t.Fatal(err)
	}

	defer db.Close()

	foos, err := sqlx.All(t.Context(), db, query, func(f *Foo) []any {
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

func TestAllErrorResetsResult(t *testing.T) {
	db, err := sql.Open(testDriver, "")
	if err != nil {
		t.Fatal(err)
	}

	defer db.Close()

	foos, err := sqlx.All(t.Context(), db, query, func(f *Foo) []any {
		return []any{&f.Id, &f.Bar}
	}, errQuery)

	if err == nil {
		t.FailNow()
	}

	if foos != nil {
		t.Fatal(foos)
	}
}
