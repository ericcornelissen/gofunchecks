package walk

func includes(vs []string, x string) bool {
	for _, v := range vs {
		if v == x {
			return true
		}
	}

	return false
}
