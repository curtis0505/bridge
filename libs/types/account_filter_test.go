package types

import (
	"testing"
)

func TestAccountFilter(t *testing.T) {
	type args struct {
		chains            []string
		estimatedItems    uint
		falsePositiveRate float64
	}
	tests := []struct {
		name string
		args args
		want *AccountFilter
	}{
		{
			name: "test",
			args: args{
				chains:            []string{"eth", "kly"},
				estimatedItems:    100000,
				falsePositiveRate: 0.00001,
			},
			want: &AccountFilter{
				estimatedItems:    100000,
				falsePositiveRate: 0.00001,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewAccountFilter(tt.args.chains, tt.args.estimatedItems, tt.args.falsePositiveRate)
			//!reflect.DeepEqual(got, tt.want) {
			//	t.Errorf("NewAccountFilter() = %v, want %v", got, tt.want)
			//}
			got.AddString("eth", "0x123456")
			got.AddString("kly", "0x456789")
			if !got.IsMemberString("eth", "0x123456") {
				t.Errorf("NewAccountFilter()  failed %v", "eth")
			}
			if !got.IsMemberString("kly", "0x456789") {
				t.Errorf("NewAccountFilter()  failed %v", "eth")
			}
			if got.IsMemberString("eth", "0xabcdef") {
				t.Errorf("NewAccountFilter()  failed %v", "eth")
			}
			if got.IsMemberString("kly", "0xabcdef") {
				t.Errorf("NewAccountFilter()  failed %v", "eth")
			}

			got.AddHexString("eth", "0x3456789")
			got.AddHexString("kly", "0x567890a")
			if !got.IsMemberHexString("eth", "0x3456789") {
				t.Errorf("NewAccountFilter()  failed %v", "eth")
			}
			if !got.IsMemberHexString("eth", "3456789") {
				t.Errorf("NewAccountFilter()  failed %v", "eth")
			}
			if !got.IsMemberHexString("kly", "0x567890a") {
				t.Errorf("NewAccountFilter()  failed %v", "eth")
			}
			if got.IsMemberHexString("eth", "87650123") {
				t.Errorf("NewAccountFilter()  failed %v", "eth")
			}
			if got.IsMemberHexString("eth", "0x87650123") {
				t.Errorf("NewAccountFilter()  failed %v", "eth")
			}
			if got.IsMemberHexString("kly", "0x87650123") {
				t.Errorf("NewAccountFilter()  failed %v", "eth")
			}
		})
	}
}
