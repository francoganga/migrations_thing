package platform

type MysqlPlatform struct {
	BasePlatform
}

func (n MysqlPlatform) BooleanTypeDeclaration() string {
	return "TINYINT(1)"
}
