package luhn

import (
	"fmt"
	"unicode"
)

func ReverseString(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}

// Valid is a card number valid?
func Valid(card string) bool {
	fmt.Println("Valid?", card)

	sum := 0
	count := 1

	// reverse it bc we count from the right
	card = ReverseString(card)
	// fmt.Println("ReversedString", card)

	// str := strings.Split(card, "")
	for _, ch := range card {
		// fmt.Println("ch", ch, string(ch))

		if unicode.IsDigit(ch) {
			//if ch <= '0' && ch >= '9' {
			n := int(ch) - int('0')

			// double every second digit
			// starting from the right
			isSecond := count%2 == 0
			if isSecond {

				//If doubling the number results greater than 9
				// then subtract 9 from the product
				prod := 2 * n
				if prod > 9 {
					prod -= 9
				}

				sum += prod
			} else {
				sum += n
			}
			// fmt.Println("digit", ch, n, "sum", sum)

			count += 1
		} else if unicode.IsSpace(ch) {
			// } else if ch == ' ' || ch == '\t' || ch == '\n' {
			continue
		} else {
			// fmt.Println("FAIL. non-digit non-space", ch)
			return false
		}
	}

	// fmt.Println("sum", sum, sum)
	// fmt.Println("count", count)

	if count <= 2 {
		return false
	}

	if sum%10 == 0 {
		return true
	}

	return false
}

func main() {
	ex1 := "4539 1488 0343 6467"
	fmt.Println("Valid", Valid(ex1))
	fmt.Println("-")

	ex2 := "059"
	fmt.Println("Valid", Valid(ex2))
	fmt.Println("-")

	// this fails
	// luhn_test.go:8: Valid(59): a simple valid SIN that becomes invalid if reversed
	//     	 Expected: true
	//     	 Got: false
	ex3 := "59"
	fmt.Println("Valid", Valid(ex3))
	fmt.Println("-")
}
