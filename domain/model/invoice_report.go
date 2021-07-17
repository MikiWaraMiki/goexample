package domain_model

type InvoiceReport struct {
	customer    string
	totalAmount *TotalAmount
	totalCredit *TotalCredit
	currency    *Currency
}

func NewInvoiceReport(customer string, totalAmount *TotalAmount, totalCredit *TotalCredit, currency *Currency) *InvoiceReport {
	return &InvoiceReport{
		customer:    customer,
		totalAmount: totalAmount,
		totalCredit: totalCredit,
		currency:    currency,
	}
}

func (this InvoiceReport) TotalAmount() string {
	total := this.totalAmount.Price()
	return this.currency.GetUsdStr(total)
}

func (this InvoiceReport) TotalCredit() int {
	return this.totalCredit.Volume()
}

func (this InvoiceReport) InvoiceDetail() string {
	return this.totalAmount.Detail()
}