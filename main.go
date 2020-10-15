package main

import "fmt"

func main() {
    getFibo(5)
}

func getFibo(n int) (int, error) {
    var array [20]int
    array[0] = 0
    array[1] = 1
    for i := 2; i < n; i++ {
        array[i] = array[i - 1] + array[i - 2]
    }

    return fmt.Println("Fibo number with index ", n, ": ", array[n - 1])
}