package FizzBuzz

import "strconv"

func Run(number int) string {
  divBy3 := number % 3 == 0
  divBy5 := number % 5 == 0

  if divBy3 && divBy5 {
    return "FizzBuzz";
  } else if divBy3 {
    return "Fizz";
  } else if divBy5 {
    return "Buzz";
  } else {
    return strconv.Itoa(number);
  }
}
