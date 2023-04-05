package printer

import (
	"bytes"
	"testing"

	"github.com/malaupa/dbus-codegen-go/token"
)

func TestPrintClient(t *testing.T) {
	t.Parallel()

	var buf bytes.Buffer
	if err := Print(&buf, []*token.Interface{
		{
			Name:       "foo.org",
			Methods:    []*token.Method{},
			Properties: []*token.Property{},
			Signals:    []*token.Signal{},
		},
	}); err != nil {
		t.Fatal(err)
	}

	// TODO: test something
}

func TestIfaceName(t *testing.T) {
	p := &context{}
	for name, want := range map[string]string{
		"camel_case_name": "Camel_Case_Name",
	} {
		if have := p.tplIfaceType(&token.Interface{
			Name: name,
		}); have != want {
			t.Errorf("ifaceType(%q) = %q, want %q", name, have, want)
		}
	}
}
