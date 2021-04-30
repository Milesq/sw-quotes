package getscene

import (
	"regexp"
	"strconv"
)

func parseQuery(r *regexp.Regexp, s string) Query {
	atoi := func(str string) int {
		num, err := strconv.Atoi(str)
		if err != nil {
			return -1
		}

		return num
	}

	parts := r.FindStringSubmatch(s)

	return Query{
		MovieID: atoi(parts[2]),
		BegPhrase: Phrase{
			Str:    parts[3],
			Offset: atoi(parts[5]),
			I:      atoi(parts[7]),
		},
		EndPhrase: Phrase{
			Str:    parts[8],
			Offset: atoi(parts[10]),
			I:      atoi(parts[12])},
	}
}
