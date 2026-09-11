func groupAnagrams(strs []string) [][]string {
	//character sort
	m := map[string][]string{}
	for _,s := range strs {
		sortedS := sortString(s)
		m[sortedS] = append(m[sortedS],s)
	}
	res := [][]string{}
	for _,v := range m {
		res = append(res,v)
	}
	return res
}

func sortString(s string) string {
	bs := []byte(s)
	sort.Slice(bs,func(i,j int)bool {
		return bs[i] < bs[j]
	})
	return string(bs)
}
