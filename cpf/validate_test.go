package cpf

import (
  "testing"
)

func TestValidateValid(t *testing.T) {
  cpf := "19246233310"
  want := true
  msg, err := Validate(cpf)
  if msg != want {
    t.Fatalf("Validate(\"%s\") = %v, %v, want %v", cpf, msg, err, want)
  }
} 

func TestValidateInvalid(t *testing.T) {
  cpf := "19246233319"
  want := false
  msg, err := Validate(cpf)
  if msg != want {
    t.Fatalf("Validate(\"%s\") = %v, %v, want %v", cpf, msg, err, want)
  }
} 

func TestValidateIncorrect(t *testing.T) {
  cpf := "1924623331"
  want := false
  msg, err := Validate(cpf)
  if err == nil {
    t.Fatalf("Validate(\"%s\") = %v, %v, want %v", cpf, msg, err, want)
  }
} 
