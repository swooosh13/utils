package xutils

func Contains(s []string, x string) bool {
	for _, n := range s {
		if n == x {
			return true
		}
	}
	return false
}

func ContainsInt(a []int, x int) bool {
	for _, n := range a {
		if n == x {
			return true
		}
	}
	return false
}

