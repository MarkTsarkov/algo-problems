func romanToInt(s string) int {
    compares := map[string]int{
        "I" : 1,
        "V" : 5,
        "X" : 10,
        "L" : 50,
        "C" : 100,
        "D" : 500,
        "M" : 1000,
    }

    integer := 0

    for i := 0; i < len(s); i++ {
        current := string(s[i])
        next := ""

        if i < (len(s)-1) {
            next = string(s[i+1])
        }

        if compares[current] < compares[next] {
            integer = integer + compares[next] - compares[current]
            i++
        } else {
            integer += compares[current]
        }
    }
    return integer
}
