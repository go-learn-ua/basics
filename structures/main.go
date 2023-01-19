package main

import "fmt"

type wallet struct {
	creditCards []creditCard
}

type creditCard struct {
	number         int
	expirationDate string
	cvvCode        int
	holder         string
}

func main() {
	ivankoCreditCard := creditCard{
		number:         4224580604397698,
		expirationDate: "20 липня 2031р",
		cvvCode:        230,
		holder:         "Іванко",
	}

	ivankoMonoCreditCard := creditCard{
		number:         4114391864397698,
		expirationDate: "20 липня 2025р",
		cvvCode:        892,
		holder:         "Іванко",
	}

	ivankoPrivatCreditCard := creditCard{
		number:         3911391723597698,
		expirationDate: "8 вересня 2025р",
		cvvCode:        777,
		holder:         "Іванко",
	}

	myWallet := wallet{
		creditCards: []creditCard{
			ivankoMonoCreditCard,
			ivankoPrivatCreditCard,
			ivankoCreditCard,
		},
	}

	payWithWallet(myWallet, 100)

	cardNumber := 4224580604397698
	cardExpirationDate := "20 липня 2031р"
	cardCvvCode := 230
	cardHolderName := "Іванко"
	amountToPay := 100

	pay(cardNumber, cardExpirationDate, cardCvvCode, cardHolderName, amountToPay)

	ivankoMonobankCreditCardNumber := 4114391864397698
	ivankoMonobankCardExpirationDate := "20 липня 2025р"
	ivankoMonobankCardCvvCode := 892
	ivankoMonobankCardHolderName := "Іванко"

	pay(
		ivankoMonobankCreditCardNumber,
		ivankoMonobankCardExpirationDate,
		ivankoMonobankCardCvvCode,
		ivankoMonobankCardHolderName,
		amountToPay,
	)

	ivankoPrivatBankCreditCardNumber := 3911391723597698
	ivankoPrivatBankCardExpirationDate := "8 вересня 2025р"
	ivankoPrivatBankCardCvvCode := 777
	ivankoPrivatBankCardHolderName := "Іванко"

	pay(
		ivankoPrivatBankCreditCardNumber,
		ivankoPrivatBankCardExpirationDate,
		ivankoPrivatBankCardCvvCode,
		ivankoPrivatBankCardHolderName,
		amountToPay,
	)

}

func payWithCreditCard(c creditCard, amount int) {
	fmt.Println("Інформація по картці")
	fmt.Println("Номер картки: ", c.number)
	fmt.Println("Tермін дії: ", c.expirationDate)
	fmt.Println("Cvv код: ", c.cvvCode)
	fmt.Println("Iмя власника: ", c.holder)
	fmt.Println("Cума до сплати: ", amount)
	fmt.Println()
}

func pay(cardNumber int, cardExpirationDate string, cardCvvCode int, cardHolderName string, amount int) {
	fmt.Println("Інформація по картці")
	fmt.Println("Номер картки: ", cardNumber)
	fmt.Println("Tермін дії: ", cardExpirationDate)
	fmt.Println("Cvv код: ", cardCvvCode)
	fmt.Println("Iмя власника: ", cardHolderName)
	fmt.Println("Cума до сплати: ", amount)
	fmt.Println()
}

func payWithWallet(w wallet, amount int) {
	for _, card := range w.creditCards {
		payWithCreditCard(card, amount)
	}
}
