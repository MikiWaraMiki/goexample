package domain_model

import "fmt"

type Amount struct {
	play     *Play
	seat     int
	price    int
	currency *Currency
}

func NewAmount(play *Play, performance *Performance, currency *Currency) *Amount {
	price := calcAmount(play, performance)
	return &Amount{
		play:     play,
		seat:     performance.Audience,
		price:    price,
		currency: currency,
	}
}

func (this Amount) Price() int {
	return this.price
}

func (this Amount) FormattedPrice() string {
	return this.currency.GetUsdStr(this.price)
}

func (this Amount) Seat() int {
	return this.seat
}

func (this Amount) PlayName() string {
	return this.play.Name
}

func (this Amount) PlainText() string {
	return fmt.Sprintf("%v: %v (%v seats)\n", this.play.Name, this.currency.GetUsdStr(this.price), this.seat)
}

func calcAmount(play *Play, performance *Performance) int {
	if play.IsTragedy() {
		amount := 40000
		if performance.Audience > 30 {
			addCost := 1000 * (performance.Audience - 30)
			amount += addCost
		}
		return amount
	}

	if play.IsComedy() {
		amount := 30000
		if performance.Audience > 20 {
			addCost := 10000 + 500*(performance.Audience-20)
			amount += addCost
		}
		amount += 300 * performance.Audience
		return amount
	}
	return 0
}
