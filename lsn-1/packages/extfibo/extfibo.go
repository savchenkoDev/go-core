package extfibo

import (
    "errors"
    "strconv"
    "../fibo"
)

const lim = 20

// Nums - return fibo numbers with limit or out of limit message
func Nums(n []int) ([]int, error) {
    var err error
    answer := make([]int,len(n))
    for i := 0 ; i <= len(n) - 1 ; i++ {
        if n[i] > lim {
            message := "Out of limit for index " + strconv.Itoa(n[i])
            err = errors.New(message)
        }
        answer[i] = fibo.Num(n[i])
    }
    return answer, err
}