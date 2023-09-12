package trade_indicators

import (
	"errors"
	"math"
)

// CalculateStochastic Stochastic oscillator
func CalculateStochastic(high []float64, low []float64, close []float64, period int) (k []float64, d []float64, err error) {
	validLen := len(high) == len(low) && len(close) == len(low)
	if !validLen {
		return nil, nil, errors.New("high, low and close array prices must be same len")
	}

	min := MinOf(low[:period]...)
	max := MaxOf(high[:period]...)

	for i := period; i <= len(close)-1; i++ {
		min = math.Min(min, low[i])
		max = math.Max(max, high[i])

		k = append(k, roundFloat((close[i]-min)/(max-min)*100, 2))
	}

	d, err = CalculateSMA(k, 3)
	if err != nil {
		return nil, nil, err
	}

	return k, d, nil
}
