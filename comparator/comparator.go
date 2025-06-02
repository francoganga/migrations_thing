package comparator

import (
	"minsi/platform"
	"minsi/schema"
)

type Comparator struct {
	Platform platform.Platform
}

func (c *Comparator) columnsEqual(c1, c2 schema.Column) bool {

	panic("todo")
}

func (c *Comparator) CompareTables(oldTable, newTable schema.Table) schema.TableDiff {

	addedColumns := []schema.Column{}
	changedColumns := map[string]schema.ColumnDiff{}
	droppedColumns := []schema.Column{}
	addedIndexes := []schema.Index{}
	modifiedIndexes := []schema.Index{}
	droppedIndexes := []schema.Index{}
	renamedIndexes := []schema.Index{}
	addedForeignKeys := []schema.ForeignKeyConstraint{}
	modifiedForeignKeys := []schema.ForeignKeyConstraint{}

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
		// check if columns are equal

	}

	return schema.TableDiff{
		OldTable:            oldTable,
		AddedColumns:        addedColumns,
		ChangedColumns:      changedColumns,
		DroppedColumns:      droppedColumns,
		AddedIndexes:        addedIndexes,
		ModifiedIndexes:     modifiedIndexes,
		DroppedIndexes:      droppedIndexes,
		RenamedIndexes:      renamedIndexes,
		AddedForeignKeys:    addedForeignKeys,
		ModifiedForeignKeys: modifiedForeignKeys,
	}

}

func (c *Comparator) CompareSchemas(oldSchema, newSchema *schema.Schema) schema.SchemaDiff {

	createdSchemas := []string{}
	droppedSchemas := []string{}
	createdTables := []schema.Table{}
	alteredTables := []schema.TableDiff{}
	droppedTables := []schema.Table{}
	createdSequences := []schema.Sequence{}
	alteredSequences := []schema.Sequence{}
	droppedSequences := []schema.Sequence{}

	for _, ns := range newSchema.Namespaces {
		if !oldSchema.HasNamespace(ns) {
			createdSchemas = append(createdSchemas, ns)
		}
	}

	for _, ons := range oldSchema.Namespaces {
		if !newSchema.HasNamespace(ons) {
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

			// tableDiff := CompareTables(
			// 	oldSchema.Tables[tableName],
			// 	newSchema.Tables[tableName],
			// )

		}
	}

	for tableName, oldTable := range oldSchema.Tables {
		_, exists := newSchema.Tables[tableName]
		if !exists {
			droppedTables = append(droppedTables, oldTable)
		}
	}

	return schema.SchemaDiff{
		CreatedSchemas:   createdSchemas,
		DroppedSchemas:   droppedSchemas,
		CreatedTables:    createdTables,
		AlteredTables:    alteredTables,
		DroppedTables:    droppedTables,
		CreatedSequences: createdSequences,
		AlteredSequences: alteredSequences,
		DroppedSequences: droppedSequences,
	}
}

