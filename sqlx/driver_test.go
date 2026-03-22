package sqlx_test

import (
	"database/sql"
	"database/sql/driver"
	"errors"
	"fmt"
	"io"
	"testing"
)

const (
	testDriver = "TestScan"
	query      = "SELECT id, bar FROM Foo"
)

var (
	rows = [][]driver.Value{
		{1, "Foo"},
		{2, "Bar"},
		{3, "Baz"},
	}
	cols = []string{"id", "bar"}

	errQuery     = errors.New("connection dropped mid query")
	errRowsClose = errors.New("connection dropped mid close")
)

type Driver struct {
	Data map[string][][]driver.Value
	Cols map[string][]string
}

type Conn struct {
	driver *Driver
}

type Stmt struct {
	driver *Driver
	query  string
}

type Rows struct {
	errClose error
	columns  []string
	data     [][]driver.Value
	idx      int
}

type Foo struct {
	Bar string
	Id  int64
}

func (d *Driver) Open(name string) (driver.Conn, error) {
	return &Conn{driver: d}, nil
}

func (c *Conn) Prepare(query string) (driver.Stmt, error) {
	return &Stmt{
		query:  query,
		driver: c.driver,
	}, nil
}

func (c *Conn) Close() error              { return nil }
func (c *Conn) Begin() (driver.Tx, error) { return nil, nil }

func (s *Stmt) Close() error  { return nil }
func (s *Stmt) NumInput() int { return -1 }

func (s *Stmt) Exec(args []driver.Value) (driver.Result, error) {
	return driver.RowsAffected(0), nil
}

func (s *Stmt) Query(args []driver.Value) (driver.Rows, error) {
	for _, arg := range args {
		if arg == errQuery.Error() {
			return nil, errQuery
		}
	}

	d, ok := s.driver.Data[s.query]
	if !ok {
		return nil, fmt.Errorf("no fake data for query: %s", s.query)
	}

	cs := s.driver.Cols[s.query]
	if cs == nil {
		return nil, fmt.Errorf("no columns defined for query: %s", s.query)
	}

	rs := Rows{
		columns: cs,
		data:    d,
	}

	for _, arg := range args {
		if arg == errRowsClose.Error() {
			rs.errClose = errRowsClose
		}
	}

	return &rs, nil
}

func (r *Rows) Columns() []string { return r.columns }
func (r *Rows) Close() error      { return r.errClose }

func (r *Rows) Next(dest []driver.Value) error {
	if r.idx >= len(r.data) {
		return io.EOF
	}
	copy(dest, r.data[r.idx])
	r.idx++
	return nil
}

func TestMain(m *testing.M) {
	d := &Driver{
		Data: map[string][][]driver.Value{
			query: rows,
		},
		Cols: map[string][]string{
			query: cols,
		},
	}

	sql.Register(testDriver, d)
	m.Run()
}
