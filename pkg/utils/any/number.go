package any

import "strconv"

func ParseInt(a any) int {
	s, ok := a.(string)
	if !ok {
		return 0
	}

	i, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return 0
	}

	return int(i)
}
