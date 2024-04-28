package a

import "testing"

func TestExportedFunction(t *testing.T) {
	words := "tom"
	want := "my_pkg:" + words
	if got := Say(words); got != want {
		t.Errorf("ExportedFunction() = %v, want %v", got, want)
	}
}
