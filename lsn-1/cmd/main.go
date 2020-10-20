package main

import (
    "../packages/extfibo"
    "fmt"
    "flag"
    "strings"
    "strconv"
)

type numbers []int

func (i *numbers) String() string {
    return fmt.Sprint(*i)
}

func (i *numbers) Set(value string) error {
    for _, num := range strings.Split(value, ",") {
        flag, err := strconv.Atoi(num)
        if err != nil {
            return err
        }
        *i = append(*i, flag)
    }
    return nil
}

var numsFlag numbers

func init() {
    flag.Var(&numsFlag, "n", "fibo numbers list")
}

const limit = 20

func main() {
    flag.Parse()
    nums, err := extfibo.Nums(numsFlag)
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println(nums)
    }
}

