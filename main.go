package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
)

type Performance struct {
	PlayID   string `json:"playID"`
	Audience int    `json:"audience"`
}

type Invoice struct {
	Customer    string        `json:"customer"`
	Performance []Performance `json:"performances"`
}

type Play struct {
	PlayID   string `json:"playID"`
	Name     string `json:"name"`
	TypeName string `json:"typeName"`
}

func Find(playId string, plays []Play) (*Play, error) {
	for i := range plays {
		if plays[i].PlayID == playId {
			return &plays[i], nil
		}
	}
	return nil, errors.New("not found")
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

func VolumeCreditsFor(performance Performance, play *Play) int {
	volumeCredits := int(math.Max(float64(performance.Audience-30), 0))

	if play.TypeName == "comedy" {
		volumeCredits += int(math.Trunc(float64(performance.Audience) / 5))
	}

	return volumeCredits
}

func Statement(invoice Invoice, plays []Play) (string, error) {
	totalAmount := 0
	volumeCredits := 0

	result := fmt.Sprintf("Statement for %v\n", invoice.Customer)

	for _, performance := range invoice.Performance {
		play, err := Find(performance.PlayID, plays)
		if err != nil {
			return "", err
		}
		thisDollar := float64(AmountFor(&performance, play) / 100)
		result += fmt.Sprintf("%v: $%.2f (%v seats)\n", play.Name, thisDollar, performance.Audience)
		totalAmount += AmountFor(&performance, play)
		volumeCredits += VolumeCreditsFor(performance, play)
	}

	totalDollar := float64(totalAmount / 100)
	result += fmt.Sprintf("Amount owed is $%.2f\n", totalDollar)
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
