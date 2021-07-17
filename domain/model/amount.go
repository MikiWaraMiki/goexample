package domain_model

type Amount struct {
	price int
}

func NewAmount(play *Play, performance *Performance) *Amount {
	price := calcAmount(play, performance)
	return &Amount{
		price: price,
	}
}

func (this Amount) Price() int {
	return this.price
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
