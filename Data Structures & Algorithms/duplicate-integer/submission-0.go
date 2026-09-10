func hasDuplicate(nums []int) bool {
    //map range
    m := map[int]bool{}
    for _,i := range nums {
        if len(m) == 0 {
            m[i] = true
        } else {
            if m[i] {
                return true
            }
            m[i] = true
        }
    }
    return false   
}
