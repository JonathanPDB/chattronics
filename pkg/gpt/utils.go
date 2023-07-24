package gpt

import (
	"github.com/JonathanPDB/chattronics/pkg/logging"
	"github.com/adrg/exrates"
)

func getBRLfromUSD(usdValue float64) float64 {
	rates, err := exrates.Latest("USD", []string{"BRL"})
	if err != nil {
		logging.Warn("Failed to get exchange rates")
	}
	return rates.Values["BRL"] * usdValue
}
