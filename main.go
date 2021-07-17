package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	. "github.com/MikiWaraMiki/goexample/domain/model"
	. "github.com/MikiWaraMiki/goexample/domain/service"
	. "github.com/MikiWaraMiki/goexample/presentation/plain_text"
)

func Statement(invoice Invoice, plays []Play) (*InvoiceReport, error) {

	currency := NewUsd()
	play_service := NewPlayService(plays)
	calc_service := NewInvoiceCalcService(play_service)

	report, err := calc_service.GenerateInvoiceReport(&invoice, currency)

	if err != nil {
		return nil, err
	}
	return report, nil
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
		report, err := Statement(invoice, plays)

		if err != nil {
			log.Fatal(err)
			panic("report generate failed")
		}

		plain_render := NewInvoicePlainText(report)

		fmt.Println(plain_render.Render())
	}
	fmt.Println("exit")
}
