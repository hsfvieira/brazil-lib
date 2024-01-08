package cpf

import (
  "testing"
)

func TestFormat(t *testing.T) {
  cpf := "19246233310"
  want := "192.462.333-10"
  cpfFormatted, err := Format(cpf)
  if cpfFormatted != want {
    t.Fatalf("Format(\"%s\") = %v, %v, want %v", cpf, cpfFormatted, err, want)
  }
} 

func TestUnformat(t *testing.T) {
  cpf := "192.462.333-10"
  want := "19246233310"
  cpfUnformatted, err := Unformat(cpf)
  if cpfUnformatted != want {
    t.Fatalf("Format(\"%s\") = %v, %v, want %v", cpf, cpfUnformatted, err, want)
  }
} 
