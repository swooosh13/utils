package utils

func Contains(s []string, x string) bool {
	for _, n := range s {
		if n == x {
			return true
		}
	}
	return false
}
