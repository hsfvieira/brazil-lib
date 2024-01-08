package cpf

import (
  "testing"
)

func TestGenerate(t *testing.T) {
  cpf := Generate()
  want := 11
  if len(cpf) != want {
    t.Fatalf("Generate() = %v, len == %d, want len == %d", cpf, len(cpf), want)
  }
} 
