package cpf

import (
  "strings"
  "strconv"
)

var weights = [10]int{
  11,
  10,
  9,
  8,
  7,
  6,
  5,
  4,
  3,
  2,
}

func checkTotal(total int) int {
  rest := total % 11

  if rest < 2 {
    return 0
  }

  return 11 - rest
}

func firstDigit(numbers []int) int {
  total := 0
  
  for i, v := range numbers {
    total += v * weights[i + 1]
  }

  return checkTotal(total)
}

func secondDigit(numbers []int) int {
  total := 0
  
  for i, v := range numbers {
    total += v * weights[i]
  }

  return checkTotal(total)
}

func Validate(cpf string) (bool, error) {
  cpfNumbers, err := Unformat(cpf)
  if err != nil {
    return false, err
  }

  cpfSplitted := strings.Split(cpfNumbers, "")

  numbers := make([]int, 9)
  for i, _ := range numbers {
    numbers[i], _ = strconv.Atoi(cpfSplitted[i])
  }
  first, _ := strconv.Atoi(cpfSplitted[9])
  second, _ := strconv.Atoi(cpfSplitted[10])

  firstDigitGenerated := firstDigit(numbers)
  secondDigitGenerated := secondDigit(append(numbers, firstDigitGenerated))

  if firstDigitGenerated == first && secondDigitGenerated == second {
    return true, nil
  } else {
    return false, nil
  }
}
