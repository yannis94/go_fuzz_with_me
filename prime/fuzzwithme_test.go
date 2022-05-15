package prime

import (
    //fuzz "github.com/google/gofuzz"
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

/*
func Fuzz(data []byte) int {
    var i int
    fuzz.NewFromGoFuzz(data).Fuzz(&i)
    IsPrimeNumber(i)
    return 0
}
*/

func FuzzIsPrimeNumber(f *testing.F) {

    testcases := []int{3, 5, 6, 7, 1997}

    for _, tc := range testcases {
        f.Add(tc)
    }

    f.Fuzz(func(t *testing.T, orig int) {
        isPrime := IsPrimeNumber(orig)
        test := true

        if orig < 2 {
            test = false
        }
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
