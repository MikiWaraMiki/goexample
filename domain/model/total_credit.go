package domain_model

type TotalCredit struct {
	credits []Credit
}

func NewTotalCredit() *TotalCredit {
	return &TotalCredit{}
}

func (this TotalCredit) AddCredit(credit *Credit) {
	this.credits = append(this.credits)
}

func (this TotalCredit) Volume() int {
	volume := 0

	for _, credit := range this.credits {
		volume += credit.Volume()
	}

	return volume
}
