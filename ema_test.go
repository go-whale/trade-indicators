package trade_indicators

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalculateEMA(t *testing.T) {
	type args struct {
		prices []float64
		period int
	}
	tests := []struct {
		name string
		args args
		want func() ([]float64, error)
	}{
		{
			name: "correct",
			args: args{
				prices: []float64{
					283.46, 280.69, 285.48, 294.08, 293.90, 299.92, 301.15, 284.45, 294.09, 302.77, 301.97, 306.85, 305.02, 301.06, 291.97, 284.18, 286.48, 284.54, 276.82, 284.49, 275.01, 279.07, 277.85, 278.85, 283.76, 291.72, 284.73, 291.82, 296.74, 291.13,
				},
				period: 14,
			},
			want: func() ([]float64, error) {
				return []float64{295.35, 294.9, 293.47, 292.54, 291.47, 289.52, 288.85, 287, 285.94, 284.86, 284.06, 284.02, 285.05, 285.01, 285.92, 287.36, 287.86}, nil
			},
		},
		{
			name: "err",
			args: args{
				prices: []float64{
					30,
				},
				period: 5,
			},
			want: func() ([]float64, error) {
				return nil, errors.New("prices len must be more or equal 2 periods")
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CalculateEMA(tt.args.prices, tt.args.period)

			expectedResult, expectedErr := tt.want()
			assert.Equal(t, expectedResult, got)
			assert.Equal(t, expectedErr, err)
		})
	}
}
