package main

import (
    "log"
    prime "go_fuzz_with_me/prime"
)

func main() {
    log.SetFlags(0)
    log.Println("PRIME NUMBER")
    number := prime.GetRandomInt()
    log.Printf("Random number picked : %d", number)
    log.Printf("Calcul...\n\n")
    test := prime.IsPrimeNumber(number)

    if test {
        log.Printf("%d is prime !!", number)
    } else {
        log.Printf("%d is not prime :(",  number)
    }
}
