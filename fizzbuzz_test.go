package fizzbuzz

import "testing"

func TestFizzBuzzOneShouldGetOne(t *testing.T){
	result := FizzBuzz(1)
	expected := "1"
	if result != expected  {
		t.Errorf("It should return %s but got %s", expected, result)
	}
}

func TestFizzBuzzTwoShouldGetTwo(t *testing.T){
	result := FizzBuzz(2)
	expected := "2"
	if result != expected {
		t.Errorf("It should return %s but got %s", expected, result)
	}
}
