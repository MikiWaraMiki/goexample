package domain_model

import "fmt"

type InvoiceReport struct {
	customer    string
	totalAmount int
	totalCredit int
	currency    Currency
}

func NewAmount(customer string, totalAmount int, totalCredit int, currency Currency) *InvoiceReport {
	return &InvoiceReport{
		customer:    customer,
		totalAmount: totalAmount,
		totalCredit: totalCredit,
		currency:    currency,
	}
}

func (this InvoiceReport) TotalAmountReport() string {
	totalAmountUsd := this.currency.GetUsdStr(this.totalAmount)

	return fmt.Sprintf("Amount owed is %v\n", totalAmountUsd)
}
