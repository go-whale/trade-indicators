package trade_indicators

import (
	"errors"
	"math"
)

// CalculateRSI Calculate RSI for period
func CalculateRSI(prices []float64, period int) ([]float64, error) {
	if len(prices) < (2*period + 1) {
		return nil, errors.New("prices len must be more or equal 2 periods +1")
	}

	var gains, loses []float64
	var previousPrice float64
	p := float64(period)
	for _, currentPrice := range prices {
		if previousPrice == 0 {
			previousPrice = currentPrice
			continue
		}

		changePrice := currentPrice - previousPrice
		if changePrice > 0 {
			gains = append(gains, changePrice)
			loses = append(loses, 0)
		} else {
			loses = append(loses, math.Abs(changePrice))
			gains = append(gains, 0)
		}

		previousPrice = currentPrice
	}

	var avgGains, avgLoses []float64 // средние значения по убытку и профиту
	var previousAvgGain, previousAvgLose float64
	for _, g := range gains[:period] {
		previousAvgGain = previousAvgGain + g
	}

	previousAvgGain = previousAvgGain / p
	avgGains = append(avgGains, previousAvgGain)

	for _, l := range loses[:period] {
		previousAvgLose = previousAvgLose + l
	}

	previousAvgLose = previousAvgLose / p
	avgLoses = append(avgLoses, previousAvgLose)

	for _, g := range gains[period:] {
		avgGain := (previousAvgGain*(p-1) + g) / p
		avgGains = append(avgGains, avgGain)
		previousAvgGain = avgGain
	}

	for _, l := range loses[period:] {
		avgLose := (previousAvgLose*(p-1) + l) / p
		avgLoses = append(avgLoses, avgLose)
		previousAvgLose = avgLose
	}

	var rsiArr []float64
	for i, gain := range avgGains {
		var rs float64 = 1
		if avgLoses[i] != 0 {
			rs = gain / avgLoses[i]
		}

		rsi := 100 - 100/(1+rs)
		rsiArr = append(rsiArr, math.Round(rsi*100)/100)
	}

	return rsiArr, nil
}
