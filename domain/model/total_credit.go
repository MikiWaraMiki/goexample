package domain_model

type TotalCredit struct {
	credits []*Credit
}

func NewTotalCredit() *TotalCredit {
	return &TotalCredit{
		credits: make([]*Credit, 0),
	}
}

func (this *TotalCredit) AddCredit(credit *Credit) {
	this.credits = append(this.credits, credit)
}

func (this TotalCredit) Length() int {
	return len(this.credits)
}

func (this TotalCredit) Volume() int {
	volume := 0

	for _, credit := range this.credits {
		volume += credit.Volume()
	}

	return volume
}
