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

func TestFizzBuzzThreeShouldGetFizz(t *testing.T){
	result := FizzBuzz(3)
	expected := "Fizz"
	if result != expected {
		t.Errorf("It should return %s but got %s", expected, result)
	}
}

func TestFizzBuzzFourShouldGetFour(t *testing.T){
	result := FizzBuzz(4)
	expected := "4"
	if result != expected {
		t.Errorf("It should return %s but got %s", expected, result)
	}
}

func TestFizzBuzzFiveShouldGetBuzz(t *testing.T){
	result := FizzBuzz(5)
	expected := "Buzz"
	if result != expected {
		t.Errorf("It should return %s but got %s", expected, result)
	}
}

func TestFizzBuzzSixShouldGetFizz(t *testing.T){
	result := FizzBuzz(6)
	expected := "Fizz"
	if result != expected {
		t.Errorf("It should return %s but got %s", expected, result)
	}
}

func TestFizzBuzzNineShouldGetFizz(t *testing.T){
	result := FizzBuzz(9)
	expected := "Fizz"
	if result != expected {
		t.Errorf("It should return %s but got %s", expected, result)
	}
}
func TestFizzBuzzTenShouldGetBuzz(t *testing.T){
	result := FizzBuzz(10)
	expected := "Buzz"
	if result != expected {
		t.Errorf("It should return %s but got %s", expected, result)
	}
}

func TestFizzBuzzTwelveShouldGetFizz(t *testing.T){
	result := FizzBuzz(12)
	expected := "Fizz"
	if result != expected {
		t.Errorf("It should return %s but got %s", expected, result)
	}
}

func TestFizzBuzzFifthenShouldGetFizzBuzz(t *testing.T){
	result := FizzBuzz(15)
	expected := "FizzBuzz"
	if result != expected {
		t.Errorf("It should return %s but got %s", expected, result)
	}
}

func TestFizzBuzzTwentyShouldGetBuzz(t *testing.T){
	result := FizzBuzz(20)
	expected := "Buzz"
	if result != expected {
		t.Errorf("It should return %s but got %s", expected, result)
	}
}

func TestFizzBuzzThirtyShouldGetFizzBuzz(t *testing.T){
	result := FizzBuzz(30)
	expected := "FizzBuzz"
	if result != expected {
		t.Errorf("It should return %s but got %s", expected, result)
	}
}
