package domain_model

type InvoiceReport struct {
	customer    string
	totalAmount *TotalAmount
	totalCredit *TotalCredit
}

func NewInvoiceReport(customer string, totalAmount *TotalAmount, totalCredit *TotalCredit) *InvoiceReport {
	return &InvoiceReport{
		customer:    customer,
		totalAmount: totalAmount,
		totalCredit: totalCredit,
	}
}

func (this InvoiceReport) TotalAmount() int {
	return this.totalAmount.Price()
}

func (this InvoiceReport) TotalCredit() int {
	return this.totalCredit.Volume()
}
