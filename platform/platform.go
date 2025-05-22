package platform



type Platform interface {
	TranslateType(typ string) string
	GenerateSQL(diff SchemaDiff) string
}

