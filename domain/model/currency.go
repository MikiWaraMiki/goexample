package domain_model

import (
	"fmt"

	"github.com/dustin/go-humanize"
)

type Currency struct {
	val string
}

func NewUsd() *Currency {
	return &Currency{
		val: "usd",
	}
}

func (this Currency) GetUsdStr(price int) string {
	dollar := float64(price / 100)
	return fmt.Sprintf("$%v", humanize.Commaf(dollar))
}
