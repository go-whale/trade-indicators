package trade_indicators

import (
	"errors"
	"math"
)

func CalculateVolatility(prices []float64, period int) (float64, error) {
	if len(prices) < period {
		return 0, errors.New("period can't be less than prices len")
	}

	var previousPrice, dispersionSum float64
	for _, p := range prices[len(prices)-period-1:] {
		if previousPrice == 0 {
			previousPrice = p
			continue
		}

		diff := (p - previousPrice) / previousPrice * 100
		dispersionSum += math.Pow(diff, 2)

		previousPrice = p
	}

	return math.Round(math.Sqrt(dispersionSum/float64(period))*100) / 100, nil
}
