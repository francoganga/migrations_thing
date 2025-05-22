package platform

import "minsi/schema"

// type ColumnDefinition struct {
// 	Length    int
// 	Default   string
// 	NotNull   bool
// 	Charset   string
// 	Collation string
// 	ColumnDef string
// }

type Platform interface {
	TranslateType(typ string) string
	GetColumnDeclarationSQL(name string, column schema.Column) string
	columnsEqual(column1, column2 schema.Column) bool
}

