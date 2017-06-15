// Package to tranform text accent in unaccent.
package unaccented

import (
	"regexp"
)

// Unacented returns a string unaccent.
func Unaccented(str string) string {
	replace := map[string]string{
		"[áàãâä]": "a",
		"[çĉ]":    "c",
		"[éèẽêë]": "e",
		"[ǵĝ]":    "g",
		"[ĥḧ]":    "h",
		"[íìĩîï]": "i",
		"[ĵ]":     "j",
		"[ḱ]":     "k",
		"[ĺ]":     "l",
		"[ḿ]":     "m",
		"[ńǹñ]":   "n",
		"[óòõôö]": "o",
		"[ṕ]":     "p",
		"[ŕ]":     "r",
		"[śŝ]":    "s",
		"[ẗ]":     "t",
		"[úùũûü]": "u",
		"[ǘǜṽ]":   "v",
		"[ẃẁŵẅ]":  "w",
		"[ẍ]":     "x",
		"[ýỳỹŷÿ]": "y",
		"[źẑ]":    "z",
		"[ÁÀÂÃÄ]": "A",
		"[ÇĈ]":    "C",
		"[ÉÈẼÊË]": "E",
		"[ǴĜ]":    "G",
		"[ĤḦ]":    "H",
		"[ÍÌÎĨÏ]": "I",
		"[Ĵ]":     "J",
		"[Ḱ]":     "K",
		"[Ĺ]":     "L",
		"[Ḿ]":     "M",
		"[ŃǸÑ]":   "N",
		"[ÓÒÕÔÖ]": "O",
		"[Ṕ]":     "P",
		"[Ŕ]":     "R",
		"[ŚŜ]":    "S",
		"[ÚÙŨÛÜ]": "U",
		"[ǗǛṼ]":   "V",
		"[ẂẀŴẄ]":  "W",
		"[Ẍ]":     "X",
		"[ÝỲỸŶŸ]": "Y",
		"[ŹẐ]":    "Z",
	}

	for key, value := range replace {
		re := regexp.MustCompile(key)
		str = re.ReplaceAllString(str, value)
	}

	return str
}
