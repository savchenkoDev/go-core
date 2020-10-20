package fibo

// Num - return fibo number
func Num(n int) (int) {
    array := []int{0, 1}
    for i := 2; i < n ; i++ {
        array = append(array, array[i - 1] + array[i - 2])
    }
    return array[n - 1]
}