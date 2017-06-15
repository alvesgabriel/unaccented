package unaccented

import (
	"regexp"
)

// Return a string without accent
func Unaccented(str string) string {
	replace := []string{"[áàãâä]", "[çĉ]", "[éèẽêë]", "[ǵĝ]", "[ĥḧ]", "[íìĩîï]", "[ĵ]", "[ḱ]", "[ĺ]", "[ḿ]", "[ńǹñ]", "[óòõôö]", "[ṕ]", "[ŕ]", "[śŝ]", "[ẗ]", "[úùũûü]", "[ǘǜṽ]", "[ẃẁŵẅ]", "[ẍ]", "[ýỳỹŷÿ]", "[źẑ]",
		"[ÁÀÂÃÄ]", "[ÇĈ]", "[ÉÈẼÊË]", "[ǴĜ]", "[ĤḦ]", "[ÍÌÎĨÏ]", "[Ĵ]", "[Ḱ]", "[Ĺ]", "[Ḿ]", "[ŃǸÑ]", "[ÓÒÕÔÖ]", "[Ṕ]", "[Ŕ]", "[ŚŜ]", "[ÚÙŨÛÜ]", "[ǗǛṼ]", "[ẂẀŴẄ]", "[Ẍ]", "[ÝỲỸŶŸ]", "[ŹẐ]"}
	letters := []string{"a", "c", "e", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "r", "s", "t", "u", "v", "w", "x", "y", "z",
		"A", "C", "E", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "R", "S", "U", "V", "W", "X", "Y", "Z"}

	for i := 0; i < len(replace); i++ {
		re := regexp.MustCompile(replace[i])
		str = re.ReplaceAllString(str, letters[i])
	}

	return str
}
