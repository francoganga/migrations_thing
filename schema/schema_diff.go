package schema

type ColumnDiff struct{}
type Index struct{}
type ForeignKeyConstraint struct{}

type TableDiff struct {
	OldTable            Table
	AddedColumns        []Column
	ChangedColumns      map[string]ColumnDiff
	DroppedColumns      []Column
	AddedIndexes        []Index
	ModifiedIndexes     []Index
	DroppedIndexes      []Index
	RenamedIndexes      []Index
	AddedForeignKeys    []ForeignKeyConstraint
	ModifiedForeignKeys []ForeignKeyConstraint
}

type SchemaDiff struct {
	CreatedSchemas   []string
	DroppedSchemas   []string
	CreatedTables    []Table
	AlteredTables    []TableDiff
	DroppedTables    []Table
	CreatedSequences []Sequence
	AlteredSequences []Sequence
	DroppedSequences []Sequence
}

