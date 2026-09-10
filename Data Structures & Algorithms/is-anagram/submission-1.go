func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	bs := []byte(s) 
	ts := []byte(t)

	ms,mt := map[byte]int{},map[byte]int{}
	for _,i := range bs {
		ms[i] ++
	}
	for _,i := range ts{
		mt[i] ++
	}
	for k,v := range ms {
		if v != mt[k] {
			return false
		}
	} 
	return true         
}
