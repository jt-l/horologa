package horologa

import "testing"

func TestLcmSame(t *testing.T) {
  var expectedResult int = 10

  result, err := lcm(10, 10)

  if err != nil {
    t.Fatalf("unexpected error LCM")
  }

  if result != expectedResult {
    t.Fatalf("Expected %d but got %d", result, expectedResult)
  }
}

func TestLcmDiff(t *testing.T) {
  var expectedResult int = 15

  result, err := lcm(3, 5)

  if err != nil {
    t.Fatalf("unexpected error LCM")
  }

  if result != expectedResult {
    t.Fatalf("Expected %d but got %d", result, expectedResult)
  }
}

func TestLcmPrime(t *testing.T) {
  var expectedResult int = 7*11

  result, err := lcm(7, 11)

  if err != nil {
    t.Fatalf("unexpected error LCM")
  }

  if result != expectedResult {
    t.Fatalf("Expected %d but got %d", result, expectedResult)
  }
}
