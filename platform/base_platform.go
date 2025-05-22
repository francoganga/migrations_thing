package platform

import "minsi/schema"

type BasePlatform struct {
}

func (b *BasePlatform) columnsEqual(column1, column2 schema.Column) bool {

	panic("todo")
}

func (b *BasePlatform) GetColumnDeclarationSQL(name string, column schema.Column) string {

	var declaration string

	if column.ColumnDefinition != "" {
		declaration = column.ColumnDefinition
	} else {

		//default_ := column.Default

	}

	panic("todo")
	return name + " " + declaration
}

func (b *BasePlatform) GetDefaultValueDeclarationSQL(column schema.Column) string {

	if column.Default == "" {
		if !column.NotNull {
			return " DEFAULT NULL"
		} else {
			return ""
		}
	}

	// default_ := column.Default

	panic("todo")
}

