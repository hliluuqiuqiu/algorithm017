func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	smap := map[rune]int{}
	for _, val := range s {
		smap[val]++
	}
	for _, val := range t {
		smap[val]--
		if smap[val] < 0 {
			return false
		} else if smap[val] == 0 {
			delete(smap, val)
		}

	}
	return len(smap) == 0
}