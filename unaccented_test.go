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
	{"fantástico", "fantastico"},
	{"hã", "ha"},
	{"ícaro", "icaro"},
	{"já", "ja"},
	{"key", "key"},
	{"ĺĩôà", "lioa"},
	{"mãe", "mae"},
	{"el niño", "el nino"},
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
