package main

import "fmt"

type selectQuery struct {
	fields     string
	table_name string
}

type SelectQueryBuilder struct {
	select_query selectQuery
}

func NewSelectQueryBuilder() *SelectQueryBuilder {
	return &SelectQueryBuilder{select_query: selectQuery{}}
}

func (b *SelectQueryBuilder) SELECT(fields string) *SelectQueryBuilder {
	b.select_query.fields = fields

	return b
}

func (b *SelectQueryBuilder) FROM(table_name string) *SelectQueryBuilder {
	b.select_query.table_name = table_name

	return b
}

func (b *SelectQueryBuilder) GetFinalQuery() string {
	return fmt.Sprintf("SELECT %s FROM %s;",
		b.select_query.fields, b.select_query.table_name)
}

func main() {
	squery := NewSelectQueryBuilder()
	final_query := squery.SELECT("*").FROM("books").GetFinalQuery()

	fmt.Println(final_query)
}
