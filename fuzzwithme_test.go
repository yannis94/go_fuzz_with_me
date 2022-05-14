package main

import (
    "testing"
)
/*
func TestIsPrimeNumber(t *testing.T) {
    num := isPrimeNumber(3)

    if num == false {
        log.Printf("Should be a prime number")
        t.Fail()
    }
}
*/
func FuzzIsPrimeNumber(f *testing.F) {
    f.Add(3, 6, 7, 1997)

    f.Fuzz(func(t *testing.T, orig int) {
        isPrime := isPrimeNumber(orig)
        test := true

        for i := 2; i < orig; i++ {
            if orig%i == 0 {
                test = false
            }
        }

        if isPrime != test {
            var err string
            if test == true {
                err = "should be"
            } else {
                err = "shouldn't be"
            }

            t.Errorf("%d %s a prime number", orig, err)

        }
    })
}
