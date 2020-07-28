/**
 *
 * @author liangjf
 * @create on 2020/7/28
 * @version 1.0
 */
package _3

import "testing"

func Test_verifyPostorder(t *testing.T) {
	type args struct {
		postorder []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		//{
		//	name: "Test_verifyPostorder",
		//	args: args{postorder: []int{1, 6, 3, 2, 5}},
		//	want: false,
		//},
		//{
		//	name: "Test_verifyPostorder",
		//	args: args{postorder: []int{1, 3, 2, 6, 5}},
		//	want: true,
		//},
		{
			name: "Test_verifyPostorder",
			args: args{postorder: []int{4, 8, 6, 12, 16, 14, 10}},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := verifyPostorder(tt.args.postorder); got != tt.want {
				t.Errorf("verifyPostorder() = %v, want %v", got, tt.want)
			}
		})
	}
}
