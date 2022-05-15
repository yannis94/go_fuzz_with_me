package prime

import (
    "time"
    "math/rand"
)

func IsPrimeNumber(num int) bool {
    var isPrime bool = true

    if num < 2 {
        isPrime = false
    } else {
        for i := 2; i < num; i++ {
            if num % i == 0 {
                isPrime = false
                break
            }
        }
    }

    return isPrime
}

func GetRandomInt() int {
    rand.Seed(time.Now().UnixNano())
    randomInt := rand.Intn(50000)
    return randomInt
}

