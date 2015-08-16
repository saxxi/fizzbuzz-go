package main

import (
  "fmt"
  "github.com/saxxi/fizzbuzz-go/FizzBuzz"
)

func main() {

  for i := 1; i <= 100; i++ {
    fmt.Println(FizzBuzz.Run(i))
  }

}
