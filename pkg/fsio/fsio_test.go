package fsio

import (
	"testing"

	"github.com/TheDevtop/exm/pkg/data"
)

func TestWriteTable(t *testing.T) {
	// Setup keys and valus
	var (
		aa = []string{"Trivago!", "Hilton!"}
		ab = "Hotel?"
		ac = "test.db"
		ad = make(data.Table)
	)

	// Construct table
	ad[ab] = aa

	// Execute function
	ba := WriteTable(ac, ad)

	// Check test
	if ba != nil {
		t.Fail()
	}
}
