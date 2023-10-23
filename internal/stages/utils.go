package stages

func contains(a []string, s string) bool {
	for _, m := range a {
		if m == s {
			return true
		}
	}
	return false
}
