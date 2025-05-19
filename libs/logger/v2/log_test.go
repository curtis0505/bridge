package logger

import "testing"

func Test_toSnakeCase(t *testing.T) {
	tests := []struct {
		args string
		want string
	}{
		// TODO: Add test cases.
		{
			args: "NumCPU",
			want: "num_cpu",
		},
		{
			args: "Primrose",
			want: "primrose",
		},
		{
			args: "GetUserID",
			want: "get_user_id",
		},
		{
			args: "bridgeId",
			want: "bridge_id",
		},
		{
			args: "CurrencyID",
			want: "currency_id",
		},
		{
			args: "KLAY",
			want: "klay",
		},
		{
			args: "currency_id",
			want: "currency_id",
		},
	}
	for _, tt := range tests {
		t.Run("Test_toSnakeCase", func(t *testing.T) {
			if got := toSnakeCase(tt.args); got != tt.want {
				t.Errorf("toSnakeCase() = %v, want %v", got, tt.want)
			}
		})
	}
}
