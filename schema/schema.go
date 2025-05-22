package schema

import (
	"fmt"
	"strings"
)

type Column struct {
	Name          string
	Type          string
	AutoIncrement bool
	Default       string
	NotNull       bool
}

type Table struct {
	Name           string
	Columns        map[string]Column
	PrimaryKeyName string
	//ForeignKeyContraints []ForeignKeyContraint
	//UniqueConstraints []UniqueConstraint
}
type Sequence struct {
	InitialValue int
}

type Schema struct {
	Tables    map[string]Table
	Sequences map[string]Sequence
}

type SchemaDiff struct {
	TableName   string
	AddColumns  []Column
	DropColumns []string
}

func (s *Schema) AddTable(table Table) {
	s.Tables[table.Name] = table
}

func New(tables []Table) *Schema {

	schema := &Schema{}
	schema.Tables = make(map[string]Table)

	for _, t := range tables {
		schema.AddTable(t)
	}

	return schema
}

func DiffSchemas(oldSchema, newSchema Schema) []string {

	var sql []string

	for tableName, newTable := range newSchema.Tables {
		oldTable, exists := oldSchema.Tables[tableName]
		if !exists {
			// New table
			cols := []string{}
			for _, col := range newTable.Columns {
				cols = append(cols, fmt.Sprintf("%s %s", col.Name, col.Type))
			}
			sql = append(sql, fmt.Sprintf("CREATE TABLE %s (%s);", tableName, strings.Join(cols, ", ")))
			continue
		}

		// Diff columns
		for colName, newCol := range newTable.Columns {
			if oldCol, exists := oldTable.Columns[colName]; !exists {
				sql = append(sql, fmt.Sprintf("ALTER TABLE %s ADD COLUMN %s %s;", tableName, newCol.Name, newCol.Type))
			} else if oldCol.Type != newCol.Type {
				sql = append(sql, fmt.Sprintf("ALTER TABLE %s ALTER COLUMN %s TYPE %s;", tableName, newCol.Name, newCol.Type))
			}
		}

		for colName := range oldTable.Columns {
			if _, exists := newTable.Columns[colName]; !exists {
				sql = append(sql, fmt.Sprintf("ALTER TABLE %s DROP COLUMN %s;", tableName, colName))
			}
		}
	}

	// Detect dropped tables
	for tableName := range oldSchema.Tables {
		if _, exists := newSchema.Tables[tableName]; !exists {
			sql = append(sql, fmt.Sprintf("DROP TABLE %s;", tableName))
		}
	}

	return sql

}
