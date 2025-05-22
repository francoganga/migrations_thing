package platform

import "minsi/schema"

type MariadbPlatform struct {
	BasePlatform
}

var _ Platform = (*MariadbPlatform)(nil)

func (m *MariadbPlatform) TranslateType(typ string) string {

	switch typ {
	case "string":
		return "VARCHAR(255)"
	case "int":
		return "INT"
	case "bool":
		return "TINYINT(1)"
	default:
		return "TEXT"
	}
}

// TODO: just duplicate common logic in all platforms dont do oop things
func (m *MariadbPlatform) GetColumnDeclarationSQL(name string, column schema.Column) string {

	res := m.BasePlatform.GetColumnDeclarationSQL("", schema.Column{})

	return res + "mariadb"
}

