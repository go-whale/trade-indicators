package trade_indicators

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalculateRSI(t *testing.T) {
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
				return []float64{55.37, 50.07, 51.55, 50.2, 45.14, 50.48, 44.69, 47.47, 46.71, 47.45, 51.05, 56.29, 51.12, 55.58, 58.41, 54.17}, nil
			},
		},
		{
			name: "correct const price",
			args: args{
				prices: []float64{
					41, 41, 41, 41, 41, 41, 41, 41, 41, 41, 41, 41, 41, 41, 41, 41, 41, 41, 41, 41, 41, 41, 41, 41, 41, 41, 41, 41, 41, 41, 41,
				},
				period: 14,
			},
			want: func() ([]float64, error) {
				return []float64{50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50}, nil
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
				return nil, errors.New("prices len must be more or equal 2 periods +1")
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CalculateRSI(tt.args.prices, tt.args.period)

			expectedResult, expectedErr := tt.want()
			assert.Equal(t, expectedResult, got)
			assert.Equal(t, expectedErr, err)
		})
	}
}
