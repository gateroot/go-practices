package fizzbuzz

func FizzBuzz(num int) string {
	if num%15 == 0 {
		return "fizzbuzz"
	} else if num%3 == 0 {
		return "fizz"
	} else {
		return "buzz"
	}
}
