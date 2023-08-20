package differ

import "testing"

func TestRun(t *testing.T) {
	differ := NewDiffer("testdata/grammar.g4", "relational_table")
	if err := differ.Run(); err != nil {
		t.Fatal(err)
	}
}
