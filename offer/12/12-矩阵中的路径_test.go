/**
 *
 * @author liangjf
 * @create on 2020/7/25
 * @version 1.0
 */
package _2

import "testing"

/**
[["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]]
"ABCCED"
*/
func Test_exist(t *testing.T) {
	type args struct {
		board [][]byte
		word  string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test_exist",
			args: args{
				board: [][]byte{
					{'A', 'B', 'C', 'E'},
					{'S', 'F', 'C', 'S'},
					{'A', 'D', 'E', 'E'},
				},
				word: "ABCCED",
			},
			want: true,
		},
		{
			name: "Test_exist",
			args: args{
				board: [][]byte{
					{'a', 'b'},
					{'d', 'd'},
				},
				word: "abcd",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := exist(tt.args.board, tt.args.word); got != tt.want {
				t.Errorf("exist() = %v, want %v", got, tt.want)
			}
		})
	}
}
