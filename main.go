package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	. "github.com/MikiWaraMiki/goexample/domain/model"
	. "github.com/MikiWaraMiki/goexample/domain/service"
)

func Statement(invoice Invoice, plays []Play) (string, error) {
	result := fmt.Sprintf("Statement for %v\n", invoice.Customer)

	currency := NewUsd()
	play_service := NewPlayService(plays)
	calc_service := NewInvoiceCalcService(play_service)

	report, err := calc_service.GenerateInvoiceReport(&invoice, currency)

	if err != nil {
		return "", err
	}
	result += report.InvoiceDetail()
	result += fmt.Sprintf("Amount owed is %v\n", report.TotalAmount())
	result += fmt.Sprintf("You earned %v credits\n", report.TotalCredit())
	return result, nil
}

func main() {
	jsonPlays, err := ioutil.ReadFile("./datas/plays.json")
	if err != nil {
		log.Fatal(err)
		panic("json load error")
	}

	var plays []Play
	err = json.Unmarshal(jsonPlays, &plays)
	if err != nil {
		log.Fatal(err)
		panic("json convert error")
	}

	for _, play := range plays {
		fmt.Printf("%v", play.PlayID)
	}

	jsonInvoices, err := ioutil.ReadFile("./datas/invoices.json")
	if err != nil {
		log.Fatal(err)
		panic("json load error")
	}

	var invoices []Invoice
	err = json.Unmarshal(jsonInvoices, &invoices)
	if err != nil {
		log.Fatal(err)
		panic("json convert error")
	}

	if len(invoices) == 0 {
		fmt.Println("Nothing data")
		os.Exit(0)
	}
	for _, invoice := range invoices {
		result, err := Statement(invoice, plays)
		if err != nil {
			log.Fatal(err)
			panic("runtime error")
		}
		fmt.Print(result)
	}
	fmt.Println("exit")
}
