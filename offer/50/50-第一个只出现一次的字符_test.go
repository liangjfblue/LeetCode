/**
 *
 * @author liangjf
 * @create on 2020/8/10
 * @version 1.0
 */
package _0

import "testing"

func Test_firstUniqChar(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want byte
	}{
		{
			name: "Test_firstUniqChar",
			args: args{
				s: "abaccdeff",
			},
			want: 'b',
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := firstUniqChar(tt.args.s); got != tt.want {
				t.Errorf("firstUniqChar() = %v, want %v", got, tt.want)
			}
		})
	}
}
