package comparator

import (
	"minsi/platform"
	"minsi/schema"
	"testing"
)

func TestComparator(t *testing.T) {

	cmp := Comparator{}
	cmp.Platform = &platform.MariadbPlatform{}

	schema1 := schema.New("production")

	schema1.AddTable(schema.Table{
		Name: "user",
	})

	schema2 := schema.New("production-v2")

	schema2.AddTable(schema.Table{
		Name: "products",
		Columns: schema.Columns{
			"price": {
				Type: "int",
			},
			"sku": {
				Type: "int",
			},
		},
	})

	diff := cmp.CompareSchemas(schema1, schema2)

	if len(diff.DroppedTables) == 0 {
		t.Error("expected to ahve new tables")
	}

	if len(diff.CreatedTables) != 1 {
		t.Error("should have a new table")
	}

	if len(diff.CreatedTables[0].Cols()) != 2 {
		t.Error("table should have 2 columns")
	}

}

