package main

import (
	"fmt"
	"math"
)

type Performance struct {
	PlayID   string `json:"playID"`
	Audience int    `json:"audience"`
}

type Invoice struct {
	Customer     string        `json:"customer"`
	Performances []Performance `json:"performance"`
}

type Play struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type Plays map[string]Play

type NumberFormat struct {
}

func (nf NumberFormat) Format() func(v float64) string {
	return func(amount float64) string {
		return fmt.Sprintf("$%v", amount)
	}
}

func Statement(invoice Invoice, plays Plays) (string, error) {
	totalAmount := 0.0
	volumeCredits := 0
	result := fmt.Sprintf("Statement for %v\n", invoice.Customer)
	format := NumberFormat{}.Format()

	for _, perf := range invoice.Performances {
		play := plays[perf.PlayID]
		thisAmount := 0.0
		switch play.Type {
		case "tragedy":
			thisAmount = 40000
			if perf.Audience > 30 {
				thisAmount += 1000 * (float64(perf.Audience) - 30)
			}
		case "comedy":
			thisAmount = 30000
			if perf.Audience > 30 {
				thisAmount += 10000 + 500*(float64(perf.Audience)-20)
			}
			thisAmount += 300 * float64(perf.Audience)
		default:
			return "", fmt.Errorf("unknown type: %v", play.Type)
		}
		volumeCredits += int(math.Max(float64(perf.Audience)-30, 0))
		if "comedy" == play.Type {
			volumeCredits += int(math.Floor(float64(perf.Audience) / 5))
		}
		result += fmt.Sprintf("  %v: %v %v\n", play.Name, format(thisAmount/100.0), perf.Audience)
		totalAmount += thisAmount
	}
	result += fmt.Sprintf("Amount owned is %v\n", format(totalAmount/100))
	result += fmt.Sprintf("You earned %v credits\n", volumeCredits)
	return result, nil
}
