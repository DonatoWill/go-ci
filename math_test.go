package main
import "testing"

func TestSoma(t *testing.T) {
	if Soma(10, 10) != 20 {
		t.Error("Soma deve ser 20")
	}
}