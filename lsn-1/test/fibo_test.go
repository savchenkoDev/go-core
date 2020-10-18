package fibotest

import (
    "testing"
    "../packages/fibo"
)

func TestFibo(t *testing.T) {
    want := 13
    have := fibo.Num(8)
    if want != have {
        t.Errorf("Error: want %d, have %d", want, have)
    }
}
