package querybuilder

import (
	"fmt"
	"strings"

	"github.com/aklinker1/miasma/internal/server"
	fmt2 "github.com/aklinker1/miasma/internal/server/fmt"
)

type insertBuilder struct {
	table   string
	columns []string
	values  []string
	args    []any
	logger  server.Logger
}

func Insert(table string, record map[string]any) *insertBuilder {
	columns := []string{}
	values := []string{}
	args := []any{}
	for column, value := range record {
		columns = append(columns, column)
		args = append(args, value)
		values = append(values, "?")
	}
	return &insertBuilder{
		table:   table,
		columns: columns,
		args:    args,
		values:  values,
		logger:  &fmt2.Logger{},
	}
}

func (b *insertBuilder) ToSQL() (sql string, args []any) {
	args = b.args
	columns := strings.Join(b.columns, ", ")
	values := strings.Join(b.values, ", ")
	sql = fmt.Sprintf(`INSERT INTO %s (%s) VALUES (%s)`, b.table, columns, values)
	b.logger.V("SQLite Insert: %s %v", sql, args)
	return sql, args
}
