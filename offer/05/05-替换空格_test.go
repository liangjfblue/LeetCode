/**
 *
 * @author liangjf
 * @create on 2020/7/24
 * @version 1.0
 */
package _5

import "testing"

func Test_replaceSpace(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test_replaceSpace",
			args: args{s: "We are happy."},
			want: "We%20are%20happy.",
		},
		{
			name: "Test_replaceSpace",
			args: args{s: "     "},
			want: "%20%20%20%20%20",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := replaceSpace2(tt.args.s); got != tt.want {
				t.Errorf("replaceSpace() = %v, want %v", got, tt.want)
			}
		})
	}
}
