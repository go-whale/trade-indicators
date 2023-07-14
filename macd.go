package trade_indicators

import (
	"errors"
	"fmt"
)

const macdFastPeriod = 12
const macdSlowPeriod = 26
const signalPeriod = 9

// CalculateMACD with default period settings: MACD = EMA(12) - EMA(26), signal EMA(9)
func CalculateMACD(prices []float64) (macd []float64, signal []float64, err error) {
	if len(prices) < 2*macdSlowPeriod {
		return nil, nil, errors.New("prices len must be more or equal 2 periods")
	}

	ema26, _ := CalculateEMA(prices, macdSlowPeriod)
	ema12, _ := CalculateEMA(prices, macdFastPeriod)

	roundPrecision := detectPrecision(ema12[0])

	ema12 = ema12[len(ema12)-len(ema26):]

	macd = make([]float64, 0)
	for i := 0; i < len(ema26); i++ {
		macd = append(macd, roundFloat(ema12[i]-ema26[i], roundPrecision))
	}

	signal, err = CalculateEMA(macd, signalPeriod)
	if err != nil {
		return nil, nil, fmt.Errorf("can't calculate signal ema: %w", err)
	}

	return macd, signal, nil
}
