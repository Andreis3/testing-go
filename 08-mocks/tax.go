package tax

import "errors"

type Repository interface {
	SaveTax(amount float64) error
}

func CalculateTaxAndSave(amount float64, repository Repository) error {
	tax := CalculateTax2(amount)
	return repository.SaveTax(tax)
}

func CalculateTax(amount float64) (float64, error) {
	if amount <= 0 {
		return 0, errors.New("amount must be greater than zero")
	}
	if amount >= 1000 && amount <= 20000 {
		return 10, nil
	}
	if amount > 20000 {
		return 20, nil
	}
	return 5, nil
}

func CalculateTax2(amount float64) float64 {
	if amount <= 0 {
		return 0
	}
	if amount >= 1000 && amount <= 20000 {
		return 10
	}
	if amount > 20000 {
		return 20
	}
	return 5
}
