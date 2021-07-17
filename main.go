package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
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
	thisAmount := 0

	switch play.TypeName {
	case "tragedy":
		thisAmount = 40000
		if performance.Audience > 30 {
			addCost := 1000 * (performance.Audience - 30)
			thisAmount += addCost
		}
	case "comedy":
		thisAmount = 30000
		if performance.Audience > 20 {
			addCost := 10000 + 500*(performance.Audience-20)
			thisAmount += addCost
		}
		thisAmount += 300 * performance.Audience
	default:
		thisAmount = 0
	}

	return thisAmount
}
func TotalAmount(performances []Performance, plays []Play) int {
	result := 0
	for _, performance := range performances {
		play, err := Find(performance.PlayID, plays)
		if err != nil {
			continue
		}
		result += AmountFor(&performance, play)
	}
	return result
}

func VolumeCreditsFor(performance Performance, play *Play) int {
	volumeCredits := int(math.Max(float64(performance.Audience-30), 0))

	if play.TypeName == "comedy" {
		volumeCredits += int(math.Trunc(float64(performance.Audience) / 5))
	}

	return volumeCredits
}

func TotalVolumeCredits(performances []Performance, plays []Play) int {
	result := 0
	for _, performance := range performances {
		play, err := Find(performance.PlayID, plays)
		if err != nil {
			continue
		}
		result += VolumeCreditsFor(performance, play)
	}

	return result
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
	totalAmount := TotalAmount(invoice.Performance, plays)
	volumeCredits := TotalVolumeCredits(invoice.Performance, plays)

	result += fmt.Sprintf("Amount owed is %v\n", Usd(totalAmount))
	result += fmt.Sprintf("You earned %v credits\n", volumeCredits)
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
