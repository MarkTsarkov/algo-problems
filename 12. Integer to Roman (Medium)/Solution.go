func intToRoman(num int) string {
    comparesMap := map[int]string {
      1     :   "I" ,
      5     :   "V" ,
      10    :   "X" ,
      50    :   "L" ,
      100   :   "C" ,
      500   :   "D" ,
      1000  :   "M" ,
    }
    
    var roman string
    dec := 1 
    nums := strconv.Itoa(num)

    for i := (len(nums)-1); i>=0; i-- {
        current, _ := strconv.Atoi(string(nums[i]))

        if current == 5 {
            roman = comparesMap[5*dec] + roman
        } else if current == 0 {
            dec*=10
            continue
        } else if current == 4 {
            roman = comparesMap[dec] + comparesMap[dec*5] + roman
        } else if current == 9 {
            roman = comparesMap[dec] + comparesMap[dec*10] + roman
        } else {
            if current > 5 {
                number := comparesMap[dec*5]
                for j:=0; j<(current-5); j++ {
                    number = number + comparesMap[dec]
                }
                roman = number + roman
            } else { 
                for j:=0; j<current; j++ {
                    roman = comparesMap[dec] + roman
                }
            }
        }

        dec *= 10
    }


    return roman
}
