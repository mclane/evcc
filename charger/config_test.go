package charger

import (
	"testing"

	"github.com/andig/evcc/util/test"
)

func TestChargers(t *testing.T) {
	acceptable := []string{
		"invalid plugin type: ...",
		"mqtt not configured",
		"invalid charger type: nrgkick-bluetooth",
		"invalid pin:",
		"connect: no route to host",
		"connect: connection refused",
	}

	for _, tmpl := range test.ConfigTemplates("charger") {
		_, err := NewFromConfig(tmpl.Type, tmpl.Config)
		if err != nil && !test.Acceptable(err, acceptable) {
			t.Logf("%s", tmpl.Name)
			t.Error(err)
		}
	}
}
