func lengthOfLongestSubstring(s string) int {
    maxLen := 0

    if len(s) == 1 {
        return 1
    }

    for i:=0; i<(len(s)-1); i++ {

        symbols      := make(map[byte]int)
        symbols[s[i]] = 1

        curLen := 1

        for j:=i+1; j<len(s); j++ {
            if _, ok := symbols[s[j]]; !ok {
                symbols[s[j]] = 1
                curLen++
                if curLen > maxLen {
                    maxLen = curLen
                }
                continue
            }
            if curLen > maxLen {
                maxLen = curLen
            }
            break
        }
    }
    return maxLen
}
