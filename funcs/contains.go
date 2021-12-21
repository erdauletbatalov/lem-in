package funcs

// contains
func contains(s []*Vertex, k string) bool {
	for _, v := range s {
		if k == v.Key {
			return true
		}
	}
	return false
}
