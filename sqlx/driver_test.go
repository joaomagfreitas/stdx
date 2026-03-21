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

	errRowsClose = errors.New("connection dropped mid close")
)

type Driver struct {
	Data map[string][][]driver.Value
	Cols map[string][]string
}

func (d *Driver) Open(name string) (driver.Conn, error) {
	return &Conn{driver: d}, nil
}

type Conn struct {
	driver *Driver
}

func (c *Conn) Prepare(query string) (driver.Stmt, error) {
	return &Stmt{
		query:  query,
		driver: c.driver,
	}, nil
}

func (c *Conn) Close() error              { return nil }
func (c *Conn) Begin() (driver.Tx, error) { return nil, nil }

type Stmt struct {
	query  string
	driver *Driver
}

func (s *Stmt) Close() error  { return nil }
func (s *Stmt) NumInput() int { return -1 }

func (s *Stmt) Exec(args []driver.Value) (driver.Result, error) {
	return driver.RowsAffected(0), nil
}

func (s *Stmt) Query(args []driver.Value) (driver.Rows, error) {
	data, ok := s.driver.Data[s.query]
	if !ok {
		return nil, fmt.Errorf("no fake data for query: %s", s.query)
	}

	cols := s.driver.Cols[s.query]
	if cols == nil {
		return nil, fmt.Errorf("no columns defined for query: %s", s.query)
	}

	rs := Rows{
		columns: cols,
		data:    data,
	}

	for _, arg := range args {
		if arg == errRowsClose.Error() {
			rs.errClose = errRowsClose
		}
	}

	return &rs, nil
}

type Rows struct {
	columns  []string
	data     [][]driver.Value
	idx      int
	errScan  error
	errClose error
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

type Foo struct {
	Id  int64
	Bar string
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
