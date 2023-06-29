package trade_indicators

import (
	"errors"
	"math"
)

func CalculateEMA(prices []float64, period int) ([]float64, error) {
	if len(prices) < 2*period {
		return nil, errors.New("prices len must be more or equal 2 periods")
	}

	var emas []float64
	// first ema value = sma value
	var sumPrices float64
	for _, p := range prices[:period] {
		sumPrices += p
	}

	previousEma := math.Round(sumPrices/float64(period)*100) / 100
	emas = append(emas, previousEma)

	k := 2 / float64(1+period)
	for _, p := range prices[period:] {
		previousEma = emas[len(emas)-1]
		ema := (p * k) + (previousEma * (1 - k))
		emas = append(emas, math.Round(ema*100)/100)
	}

	return emas, nil
}
