func twoSum(nums []int, target int) []int {
    result := make([]int, 2)

    for i, numf := range nums {
        for j, numsec := range nums {
            if i != j {
                numbr := numf + numsec

                if numbr == target {
                    result[0] = i
                    result[1] = j
                    break
                }
            }
        }
    }
    return result
}
