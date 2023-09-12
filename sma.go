package trade_indicators

import "errors"

// CalculateSMA Calculate simple moving average
func CalculateSMA(prices []float64, period int) ([]float64, error) {
	if len(prices) < period {
		return nil, errors.New("prices len must be equal to period")
	}

	roundPrecision := detectPrecision(prices[0])

	var sma []float64
	for i := 0; i+period <= len(prices); i++ {
		var sum float64
		for _, p := range prices[i : i+period] {
			sum += p
		}

		sma = append(sma, roundFloat(sum/float64(period), roundPrecision))
	}

	return sma, nil
}
