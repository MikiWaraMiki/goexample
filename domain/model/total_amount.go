package domain_model

type TotalAmount struct {
	amounts []Amount
}

func NewTotalAmount() *TotalAmount {
	return &TotalAmount{}
}

func (this TotalAmount) AddAmount(amount *Amount) {
	this.amounts = append(this.amounts)
}

func (this TotalAmount) Price() int {
	price := 0
	for _, amount := range this.amounts {
		price += amount.Price()
	}
	return price
}
