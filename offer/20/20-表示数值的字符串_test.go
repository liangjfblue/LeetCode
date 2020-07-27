/**
 *
 * @author liangjf
 * @create on 2020/7/27
 * @version 1.0
 */
package _0

import "testing"

func Test_isNumber(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test_isNumber",
			args: args{"123"},
			want: true,
		},
		{
			name: "Test_isNumber",
			args: args{"5e2"},
			want: true,
		},
		{
			name: "Test_isNumber",
			args: args{"1a3.14"},
			want: false,
		},
		{
			name: "Test_isNumber",
			args: args{"-123"},
			want: true,
		},
		{
			name: "Test_isNumber",
			args: args{"+-123"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isNumber(tt.args.s); got != tt.want {
				t.Errorf("isNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
