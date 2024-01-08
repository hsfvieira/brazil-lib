package cpf

import (
  "regexp"
  "errors"
)

func Format(cpf string) (string, error) {
  cpfNumbers, err := Unformat(cpf)
  if err != nil {
    return "", err
  }

  re := regexp.MustCompile("^([0-9]{3})([0-9]{3})([0-9]{3})([0-9]{2})$")
  cpfFormatted := re.ReplaceAllString(cpfNumbers, "$1.$2.$3-$4")
  return cpfFormatted, nil
}

func Unformat(cpf string) (string, error) {
  re := regexp.MustCompile("[^0-9]+") 
  numbers := re.ReplaceAllString(cpf, "")

  if len(numbers) != 11 {
    return "", errors.New("Tamanho do CPF incorreto")
  } 

  return numbers, nil
}
