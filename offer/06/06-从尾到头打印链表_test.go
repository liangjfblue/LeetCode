/**
 *
 * @author liangjf
 * @create on 2020/7/24
 * @version 1.0
 */
package _6

import (
	"offer/model"
	"reflect"
	"testing"
)

func Test_reversePrint(t *testing.T) {
	type args struct {
		head *model.ListNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Test_reversePrint",
			args: args{head: &model.ListNode{
				Val: 1,
				Next: &model.ListNode{
					Val:  2,
					Next: nil,
				},
			}},
			want: []int{2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reversePrint2(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reversePrint() = %v, want %v", got, tt.want)
			}
		})
	}
}
