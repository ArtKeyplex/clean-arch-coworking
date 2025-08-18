package domain

type Money struct {
	Amount   int64
	Currency string
}

func NewMoney(amount int64, currency string) Money {
	return Money{Amount: amount, Currency: currency}
}

func (m Money) Add(other Money) Money {
	if m.Currency != other.Currency {
		panic("currency mismatch")
	}
	return Money{Amount: m.Amount + other.Amount, Currency: m.Currency}
}
