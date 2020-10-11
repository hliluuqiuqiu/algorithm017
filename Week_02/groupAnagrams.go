type bytes []byte

func (s bytes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s bytes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s bytes) Len() int {
	return len(s)
}
func groupAnagrams(strs []string) [][]string {
   var ret [][]string
	record := map[string]int{}

	for _, str := range strs {
		var data bytes = []byte(str)
		sort.Sort(data)
		keystr := string(data)
		if idx, ok := record[keystr]; ok {
			ret[idx] = append(ret[idx], str)
		} else {
			record[keystr] = len(ret)
			ret = append(ret, []string{str})
		}
	}
	return ret
}