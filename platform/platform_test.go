package platform

import (
	"minsi/schema"
	"testing"
)

func TestMain(t *testing.T) {

	mariadb := MariadbPlatform{}

	println(mariadb.GetColumnDeclarationSQL("", schema.Column{}))

}

