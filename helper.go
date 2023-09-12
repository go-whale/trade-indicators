package trade_indicators

import (
	"math"
	"strconv"
	"strings"
)

func roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}

func detectPrecision(val float64) uint {
	var precision uint = 2

	strFloat := strconv.FormatFloat(val, 'f', -1, 64)
	posDecimal := strings.Index(strFloat, ".")
	if len(strFloat[posDecimal+1:]) > int(precision) {
		return uint(len(strFloat[posDecimal:]))
	}

	return precision
}

func MinOf(vars ...float64) float64 {
	min := vars[0]

	for _, i := range vars {
		if min > i {
			min = i
		}
	}

	return min
}

func MaxOf(vars ...float64) float64 {
	max := vars[0]

	for _, i := range vars {
		if max < i {
			max = i
		}
	}

	return max
}
