package horologa

import "testing"

func TestGcd(t *testing.T) {
  var expectedResult int = 10

  result := gcd(10, 10)

  if result != expectedResult {
    t.Fatalf("Expected %d but got %d", result, expectedResult)
  }
}

func TestGcdOne(t *testing.T) {
  var expectedResult int = 1

  result := gcd(7, 9)

  if result != expectedResult {
    t.Fatalf("Expected %d but got %d", result, expectedResult)
  }
}
