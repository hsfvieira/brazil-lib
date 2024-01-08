package cpf

import (
  "math/rand"
  "strconv"
  "strings"
)

func Generate() string {
  numbers := make([]int, 9)
  for i := range numbers {
    numbers[i] = rand.Intn(10)
  }

  first := firstDigit(numbers)
  second := secondDigit(append(numbers, first))

  cpfNumbers := append(numbers, first, second)
  cpfStringValues := make([]string, len(cpfNumbers))

  for i, v := range cpfNumbers {
    cpfStringValues[i] = strconv.Itoa(v)
  }

  cpf := strings.Join(cpfStringValues, "")

  return cpf
}
