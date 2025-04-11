package count_symmetric_integers

import "strconv"

func countSymmetricIntegers(low int, high int) int {
    res := 0
    for i := low; i <= high; i++ {
        s := strconv.Itoa(i)

        if len(s) % 2 == 1 { // odd 
            continue
        }

        half := len(s) / 2
        leftSum := 0 
        rightSum := 0

        // Sum left half digits
        for i := 0; i < half; i++ {
            leftSum += int(s[i] - '0')
        }

        // Sum right half digits
        for i := half; i < len(s); i++ {
            rightSum += int(s[i] - '0')
        }

        if leftSum == rightSum {
            res++
        }
    }

    return res
}