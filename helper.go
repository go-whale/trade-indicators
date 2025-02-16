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
	// WARNING: default precision needs for int prices, because division takes place in formulas
	var defaultPrecision uint = 2

	strFloat := strconv.FormatFloat(val, 'f', -1, 64)
	posDecimal := strings.Index(strFloat, ".")

	if posDecimal == -1 {
		return defaultPrecision
	}

	decimalSignsCnt := len(strFloat[posDecimal+1:])
	if decimalSignsCnt > int(defaultPrecision) {
		// WARNING: added 1 to better precision for small values like 0.0000078, because calculate functions could divide small price
		return uint(decimalSignsCnt) + 1
	}

	return defaultPrecision
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
