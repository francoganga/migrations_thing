package schema

import (
	"fmt"
	"strings"
)

type Column struct {
	Name             string
	Type             string
	AutoIncrement    bool
	Default          string
	NotNull          bool
	ColumnDefinition string
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
	Asset
	Namespaces map[string]string
	Tables     map[string]Table
	Sequences  map[string]Sequence
}

type Asset struct {
	Namespace string
	Name      string
}

func New(tables []Table) *Schema {

	schema := &Schema{}
	schema.Tables = make(map[string]Table)

	for _, t := range tables {
		schema.AddTable(t)
	}

	return schema
}

func (s *Schema) AddTable(table Table) {
	s.Tables[table.Name] = table
}

func (s *Schema) NormalizeName(asset Schema) string {

	name := asset.Name
	if asset.Namespace != "" {
		name = fmt.Sprintf("%s.%s", s.Name, name)
	}

	return strings.ToLower(name)
}

func (s *Schema) HasNamespace(name string) bool {
	if _, ok := s.Namespaces[name]; ok {
		return true
	}

	return false
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

func CompareTables(oldTable, newTable Table) TableDiff {

	addedColumns := []Column{}
	changedColumns := map[string]ColumnDiff{}
	droppedColumns := []Column{}
	addedIndexes := []Index{}
	modifiedIndexes := []Index{}
	droppedIndexes := []Index{}
	renamedIndexes := []Index{}
	addedForeignKeys := []ForeignKeyConstraint{}
	modifiedForeignKeys := []ForeignKeyConstraint{}

	for _, nc := range newTable.Columns {
		_, exists := oldTable.Columns[nc.Name]

		if !exists {
			addedColumns = append(addedColumns, nc)
		}
	}

	for _, oc := range oldTable.Columns {
		_, exists := newTable.Columns[oc.Name]

		if !exists {
			droppedColumns = append(droppedColumns, oc)
			continue
		}

		// TODO: check if its ok to index by old name
		// newCol := newTable.Columns[oc.Name]

	}

	return TableDiff{
		oldTable,
		addedColumns,
		changedColumns,
		droppedColumns,
		addedIndexes,
		modifiedIndexes,
		droppedIndexes,
		renamedIndexes,
		addedForeignKeys,
		modifiedForeignKeys,
	}

}

func CompareSchemas(oldSchema, newSchema Schema) SchemaDiff {

	createdSchemas := []string{}
	droppedSchemas := []string{}
	createdTables := []Table{}
	alteredTables := []TableDiff{}
	droppedTables := []Table{}
	createdSequences := []Sequence{}
	alteredSequences := []Sequence{}
	droppedSequences := []Sequence{}

	for _, ns := range newSchema.Namespaces {
		if !oldSchema.hasNamespace(ns) {
			createdSchemas = append(createdSchemas, ns)
		}
	}

	for _, ons := range oldSchema.Namespaces {
		if !newSchema.hasNamespace(ons) {
			droppedSchemas = append(droppedSchemas, ons)
		}
	}

	for tableName, newTable := range newSchema.Tables {
		_, exists := oldSchema.Tables[tableName]
		if !exists {
			// new table
			createdTables = append(createdTables, newTable)
		} else {
			// diff tables

		}
	}

	return SchemaDiff{
		createdSchemas,
		droppedSchemas,
		createdTables,
		alteredTables,
		droppedTables,
		createdSequences,
		alteredSequences,
		droppedSequences,
	}
}

