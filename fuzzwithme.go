package main

import (
    "log"
    "time"
    "math/rand"
)

func isPrimeNumber(num int) bool {
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

func getRandomInt() int {
    rand.Seed(time.Now().UnixNano())
    randomInt := rand.Intn(50000)
    return randomInt
}

func main() {
    log.SetFlags(0)
    log.Println("PRIME NUMBER")
    number := getRandomInt()
    log.Printf("Random number picked : %d", number)
    log.Printf("Calcul...\n\n")
    test := isPrimeNumber(number)

    if test {
        log.Printf("%d is prime !!", number)
    } else {
        log.Printf("%d is not prime.",  number)
    }
}
