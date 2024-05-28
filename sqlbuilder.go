package sqlbuilder

import (
    "fmt"
    "reflect"
    "strings"
)

type QueryBuilder struct{}

func (qb *QueryBuilder) CreateTable(tableName string, model interface{}) string {
    t := reflect.TypeOf(model)
    var columns []string

    for i := 0; i < t.NumField(); i++ {
        field := t.Field(i)
        column := fmt.Sprintf("%s %s", field.Name, field.Tag.Get("sql"))
        columns = append(columns, column)
    }

    query := fmt.Sprintf("CREATE TABLE %s (%s);", tableName, strings.Join(columns, ", "))
    return query
}

func (qb *QueryBuilder) Select(tableName string, model interface{}) string {
    t := reflect.TypeOf(model)
    var columns []string

    for i := 0; i < t.NumField(); i++ {
        field := t.Field(i)
        columns = append(columns, field.Name)
    }

    query := fmt.Sprintf("SELECT %s FROM %s;", strings.Join(columns, ", "), tableName)
    return query
}

func (qb *QueryBuilder) Insert(tableName string, model interface{}) string {
    t := reflect.TypeOf(model)
    var columns []string
    var values []string

    for i := 0; i < t.NumField(); i++ {
        field := t.Field(i)
        columns = append(columns, field.Name)
        values = append(values, "?")
    }

    query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s);", tableName, strings.Join(columns, ", "), strings.Join(values, ", "))
    return query
}

func (qb *QueryBuilder) Update(tableName string, model interface{}, condition string) string {
    t := reflect.TypeOf(model)
    var setClauses []string

    for i := 0; i < t.NumField(); i++ {
        field := t.Field(i)
        clause := fmt.Sprintf("%s = ?", field.Name)
        setClauses = append(setClauses, clause)
    }

    query := fmt.Sprintf("UPDATE %s SET %s WHERE %s;", tableName, strings.Join(setClauses, ", "), condition)
    return query
}
