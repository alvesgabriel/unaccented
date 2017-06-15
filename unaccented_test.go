package unaccented

import (
	"testing"
)

type test struct {
	value string
	unac  string
}

var tests = []test{
	{"ágora", "agora"},
	{"biblíco", "biblico"},
	{"canção", "cancao"},
	{"dragão", "dragao"},
	{"exercício", "exercicio"},
	{"el niño", "el nino"},
	{"fantástico", "fantastico"},
	{"gabriel", "gabriel"},
	{"hã", "ha"},
	{"ícaro", "icaro"},
	{"já", "ja"},
	{"key", "key"},
	{"ĺĩôà", "lioa"},
	{"mãe", "mae"},
}

// Return a string without accent
func TestUnaccented(t *testing.T) {
	for _, pair := range tests {
		v := Unaccented(pair.value)
		if v != pair.unac {
			t.Error(
				"For", pair.value,
				"expected", pair.unac,
				"got", v,
			)
		}
	}
}
