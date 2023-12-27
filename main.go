package main

import (
	"creditcard_validation/validator"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	var ccNumber int

	appBanner()

	// read creditcard number
	if len(os.Args) > 1 {
		ccNumberBuffer, err := strconv.Atoi(os.Args[1])
		if err != nil {
			log.Println("error converting input credit card numbers")
		}

		ccNumber = ccNumberBuffer

	} else {
		panic("To run this program you must enter a credit card number")
	}

	if validator.Validate(ccNumber) {
		fmt.Println("The credit card number is valid")
	} else {
		fmt.Println("The credit card number is NOT valid")
	}

}

func appBanner() {

	fmt.Println("CreditCard Validator :: made it for fun in vim!")

	fmt.Println("------------------------------------------------")
	fmt.Println("")
}
