package movie

import (
	"regexp"
	"strconv"

	parse_query "github.com/milesq/sw-quotes/src/movie/parse-query"
)

func parseQuery(r *regexp.Regexp, s string) parse_query.Query {
	atoi := func(str string) int {
		num, err := strconv.Atoi(str)
		if err != nil {
			return 0
		}

		return num
	}

	parts := r.FindStringSubmatch(s)

	return parse_query.Query{
		MovieID: atoi(parts[2]),
		BegPhrase: parse_query.Phrase{
			Str:    parts[3],
			Offset: atoi(parts[5]),
			I:      atoi(parts[7]),
		},
		EndPhrase: parse_query.Phrase{
			Str:    parts[8],
			Offset: atoi(parts[10]),
			I:      atoi(parts[12])},
	}
}
