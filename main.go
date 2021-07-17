package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	. "github.com/MikiWaraMiki/goexample/domain/model"
	. "github.com/MikiWaraMiki/goexample/domain/service"
	"github.com/dustin/go-humanize"
)

func Find(playId string, plays []Play) (*Play, error) {
	play_service := NewPlayService(plays)

	return play_service.FetchByPlayId(playId)
}

func AmountFor(performance *Performance, play *Play) int {
	amount := NewAmount(play, performance)
	return amount.Price()
}

func VolumeCreditsFor(performance Performance, play *Play) int {
	credit := NewCredit(&performance, play)

	// TODO: Credit返す
	return credit.Volume()
}

func Usd(rawCost int) string {
	dollar := float64(rawCost / 100)
	return fmt.Sprintf("$%v", humanize.Commaf(dollar))
}

func Statement(invoice Invoice, plays []Play) (string, error) {
	result := fmt.Sprintf("Statement for %v\n", invoice.Customer)

	for _, performance := range invoice.Performance {
		play, err := Find(performance.PlayID, plays)
		if err != nil {
			return "", err
		}
		result += fmt.Sprintf("%v: %v (%v seats)\n", play.Name, Usd(AmountFor(&performance, play)), performance.Audience)
	}

	play_service := NewPlayService(plays)
	calc_service := NewInvoiceCalcService(play_service)

	report, err := calc_service.GenerateInvoiceReport(&invoice)

	if err != nil {
		return "", err
	}

	result += fmt.Sprintf("Amount owed is %v\n", Usd(report.TotalAmount()))
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
