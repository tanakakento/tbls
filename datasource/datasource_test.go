package datasource

import (
	"testing"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

var tests = []struct {
	dsn           string
	tableCount    int
	relationCount int
}{
	{"my://root:mypass@localhost:33306/testdb", 8, 6},
	{"pg://postgres:pgpass@localhost:55432/testdb?sslmode=disable", 9, 7},
	{"json://../testdata/testdb.json", 7, 9},
}

func TestAnalyzeSchema(t *testing.T) {
	for _, tt := range tests {
		schema, err := Analyze(tt.dsn)
		if err != nil {
			t.Errorf("%s", err)
		}
		expected := "testdb"
		actual := schema.Name
		if actual != expected {
			t.Errorf("actual %v\nwant %v", actual, expected)
		}
	}
}

func TestAnalyzeTables(t *testing.T) {
	for _, tt := range tests {
		schema, err := Analyze(tt.dsn)
		if err != nil {
			t.Errorf("%s", err)
		}
		expected := tt.tableCount
		actual := len(schema.Tables)
		if actual != expected {
			t.Errorf("actual %v\nwant %v", actual, expected)
		}
	}
}

func TestAnalyzeRelations(t *testing.T) {
	for _, tt := range tests {
		schema, err := Analyze(tt.dsn)
		if err != nil {
			t.Errorf("%s", err)
		}
		expected := tt.relationCount
		actual := len(schema.Relations)
		if actual != expected {
			t.Errorf("actual %v\nwant %v", actual, expected)
		}
	}
}
