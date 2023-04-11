package query

import (
	"fmt"
	"strings"
)

type where struct {
	Field   string
	Operand string
	Value   any
}

type join struct {
	First   string
	Operand string
	Second  string
	Type    string
}

type Query struct {
	Table       string
	QueryString string
	Wheres      []where
	Joins       []join
	Bindings    []any
	Columns     []string
}

type Builder struct {
	Query *Query
}

type BuilderRepository interface {
	Update(values map[string]any)
	Read()
	Create(columns []string, values [][]string)
	Delete()
	Select(columns []string)
	Where(column string, operand string, value string)
	Join(first string, operand string, second string)
	LeftJoin(first string, operand string, second string)
	RightJoin(first string, operand string, second string)
	GetQuery() *Query
}

func NewBuilder(table string) BuilderRepository {
	query := Query{Table: table, QueryString: "", Wheres: []where{}, Joins: []join{}, Bindings: []any{}, Columns: []string{"*"}}

	return &Builder{Query: &query}
}

func (b *Builder) GetQuery() *Query {
	return b.Query
}

//CRUD

func (b *Builder) Update(values map[string]any) {
	var setString string

	for column, value := range values {
		setString += fmt.Sprintf("%s = $%d", column, len(b.Query.Bindings)+1)
		b.Query.Bindings = append(b.Query.Bindings, value)
	}

	b.Query.QueryString = fmt.Sprintf("UPDATE %s SET %s, updated_at=NOW", b.Query.Table, setString)
	addWhereConditions(b.Query)
	b.Query.QueryString += ";"
}

func (b *Builder) Read() {
	b.Query.QueryString = fmt.Sprintf("SELECT %s FROM %s ", parseSelects(b.Query), b.Query.Table)
	addJoins(b.Query)
	addWhereConditions(b.Query)
	b.Query.QueryString += ";"
}

func (b *Builder) Create(columns []string, values [][]string) {
	b.Query.QueryString = fmt.Sprintf("INSERT INTO %s (%s) VALUES", b.Query.Table, strings.Join(columns, ","))

	for i, value := range values {
		row := "("

		for j, v := range value {
			row += fmt.Sprintf("$%d", j+1)
			b.Query.Bindings = append(b.Query.Bindings, v)
		}

		row += ")"

		if i == len(values)+1 {
			row += ";"
		} else {
			row += ","
		}
	}
}

func (b *Builder) Delete() {
	b.Query.QueryString = fmt.Sprintf("DELETE FROM %s ", b.Query.Table)
	addWhereConditions(b.Query)
}

// Query Operations

func (b *Builder) Select(columns []string) {
	b.Query.Columns = columns
}

func (b *Builder) Where(column string, operand string, value string) {
	nWhere := where{Field: column, Operand: operand, Value: value}
	b.Query.Wheres = append(b.Query.Wheres, nWhere)
	b.Query.Bindings = append(b.Query.Bindings, value)
}

// Joins
func (b *Builder) Join(first string, operand string, second string) {
	nJoin := join{First: first, Operand: operand, Second: second, Type: "INNER"}
	b.Query.Joins = append(b.Query.Joins, nJoin)
}

func (b *Builder) LeftJoin(first string, operand string, second string) {
	nJoin := join{First: first, Operand: operand, Second: second, Type: "LEFT"}
	b.Query.Joins = append(b.Query.Joins, nJoin)
}

func (b *Builder) RightJoin(first string, operand string, second string) {
	nJoin := join{First: first, Operand: operand, Second: second, Type: "RIGHT"}
	b.Query.Joins = append(b.Query.Joins, nJoin)
}

func parseSelects(query *Query) string {
	var columns string

	if len(query.Columns) > 0 {
		columns = strings.Join(query.Columns, ",")
	} else {
		columns = "*"
	}

	return columns
}

func addWhereConditions(query *Query) {
	var concat string = "WHERE"

	if len(query.Wheres) > 0 {
		for i, where := range query.Wheres {
			whereString := fmt.Sprintf("%s %s %s $%d ", concat, where.Field, where.Operand, query.Bindings[i+1])
			query.QueryString += whereString
			concat = "AND"
		}
	}
}

func addJoins(query *Query) {
	if len(query.Joins) > 0 {
		for _, join := range query.Joins {
			joinString := fmt.Sprintf("%s JOIN ON %s %s %s ", join.Type, join.First, join.Operand, join.Second)
			query.QueryString += joinString
		}
	}
}
