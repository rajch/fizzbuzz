package fizzbuzz_test

import (
	"fmt"
	"testing"

	"github.com/rajch/fizzbuzz"
)

func TestFizz(t *testing.T) {
	t.Log("Starting Fizz test")
	_, ok := fizzbuzz.Fizzbuzz(1)
	if ok {
		t.Logf("The value %v should not have returned true", 1)
		t.FailNow()
	}

	result, ok := fizzbuzz.Fizzbuzz(3)
	if !ok {
		t.Logf("The value %v should have returned true", 3)
		t.Fail()
	}

	if result != "fizz" {
		t.Log("Result should have been fizz")
		t.Fail()
	}

	t.Log("Ending Fizz test")
}

func ExampleFizzbuzz() {
	result, _ := fizzbuzz.Fizzbuzz(3)
	fmt.Println(result)
	// Output: fizz
}
