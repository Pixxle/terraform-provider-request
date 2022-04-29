package utility

func SliceContains(s []string, c string) bool {
	if len(s) == 0 {
		return false
	}
	for _, i := range s {
		if i == c {
			return true
		}
	}
	return false
}
