package trade_indicators

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalculateSMA(t *testing.T) {
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
					10, 11, 12, 13, 11, 14, 11, 12, 10, 14,
				},
				period: 5,
			},
			want: func() ([]float64, error) {
				return []float64{11.4, 12.2, 12.2, 12.2, 11.6, 12.2}, nil
			},
		},
		{
			name: "correct len=period",
			args: args{
				prices: []float64{
					10, 11, 12, 13, 11,
				},
				period: 5,
			},
			want: func() ([]float64, error) {
				return []float64{11.4}, nil
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
				return nil, errors.New("prices len must be equal to period")
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CalculateSMA(tt.args.prices, tt.args.period)

			expectedResult, expectedErr := tt.want()
			assert.Equal(t, expectedResult, got)
			assert.Equal(t, expectedErr, err)
		})
	}
}
