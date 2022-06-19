package sqlb

import (
	"fmt"
	"strings"

	"github.com/aklinker1/miasma/internal/server"
	fmt2 "github.com/aklinker1/miasma/internal/server/fmt"
	"github.com/samber/lo"
)

type selectBuilder struct {
	table             string
	columns           []string
	scanDest          []any
	args              []any
	where             []string
	limit             string
	offset            string
	order             *server.Sort
	includeSoftDelete bool
	pagination        *server.Pagination
	logger            server.Logger
}

func Select(table string, columns map[string]any) *selectBuilder {
	selectColumns := []string{}
	scanDest := []any{}
	for name, target := range columns {
		selectColumns = append(selectColumns, name)
		scanDest = append(scanDest, target)
	}
	return &selectBuilder{
		table:    table,
		columns:  selectColumns,
		scanDest: scanDest,
		logger:   &fmt2.Logger{},
	}
}

func (b *selectBuilder) Where(condition string, args ...any) *selectBuilder {
	b.where = append(b.where, condition)
	b.args = append(b.args, args...)
	return b
}

func (b *selectBuilder) IncludeSoftDeleted() *selectBuilder {
	b.includeSoftDelete = true
	return b
}

func (b *selectBuilder) OrderBy(sort server.Sort) *selectBuilder {
	b.order = &sort
	return b
}

func (b *selectBuilder) Paginate(pagination server.Pagination) *selectBuilder {
	b.pagination = &pagination
	return b
}

func (b *selectBuilder) ToSQL() (sql string, args []any) {
	args = b.args
	columns := strings.Join(b.columns, ", ")

	var where string
	wheres := b.where
	if !b.includeSoftDelete && lo.Contains(b.columns, "deleted_at") {
		wheres = append(wheres, "deleted_at IS NULL")
	}
	if len(b.where) > 0 {
		where = fmt.Sprintf(" WHERE %s", strings.Join(wheres, " AND "))
	}

	var order string
	if b.order != nil {
		dir := "ASC"
		if strings.ToUpper(b.order.Direction) == "DESC" {
			dir = "DESC"
		}
		order = fmt.Sprintf(" ORDER BY %s %s", b.order.Field, dir)
	}

	var limitOffset string
	if b.pagination != nil {
		limitOffset = " LIMIT ? OFFSET ?"
		args = append(args, b.pagination.Limit(), b.pagination.Offset())
	}
	sql = fmt.Sprintf(`SELECT %s FROM %s%s%s%s`, columns, b.table, where, order, limitOffset)
	b.logger.V("SQL Query: %s %v", sql, args)
	return sql, args
}

func (b *selectBuilder) ScanDest() []any {
	return b.scanDest
}
