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
