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
		atoi(parts[2]),
		parse_query.Phrase{parts[3], atoi(parts[5]), atoi(parts[7])},
		parse_query.Phrase{parts[8], atoi(parts[10]), atoi(parts[12])},
	}
}
