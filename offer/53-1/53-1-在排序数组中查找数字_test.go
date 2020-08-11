/**
 *
 * @author liangjf
 * @create on 2020/8/10
 * @version 1.0
 */
package _3_1

import "testing"

func Test_search2(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				nums:   []int{5, 7, 7, 8, 8, 10},
				target: 8,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search2(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("search2() = %v, want %v", got, tt.want)
			}
		})
	}
}
