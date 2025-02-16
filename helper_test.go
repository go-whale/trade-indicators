package trade_indicators

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_detectPrecision(t *testing.T) {
	type args struct {
		val float64
	}
	tests := []struct {
		name string
		args args
		want uint
	}{
		{
			name: "success without decimal part",
			args: args{
				val: 1000,
			},
			want: 2,
		},
		{
			name: "success with decimal part",
			args: args{
				val: 1000.00,
			},
			want: 2,
		},
		{
			name: "success with 3 decimal places",
			args: args{
				val: 0.001,
			},
			want: 4,
		},
		{
			name: "success with 5 decimal places",
			args: args{
				val: 0.00112,
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, detectPrecision(tt.args.val), "detectPrecision(%v)", tt.args.val)
		})
	}
}

func Test_roundFloat(t *testing.T) {
	type args struct {
		val       float64
		precision uint
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "success precision 6",
			args: args{
				val:       0.0023458,
				precision: 6,
			},
			want: 0.002346,
		},
		{
			name: "success 1 places precision",
			args: args{
				val:       0.28511,
				precision: 1,
			},
			want: 0.3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, roundFloat(tt.args.val, tt.args.precision), "roundFloat(%v, %v)", tt.args.val, tt.args.precision)
		})
	}
}
