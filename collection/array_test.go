/**
 * created by
 * @project tools
 * @author frankstar
 * @date 2023/2/14
 * @contact frankstarye@tencent.com
 **/

package collection

import (
	"fmt"
	"reflect"
	"testing"
)

func TestArraySplit(t *testing.T) {
	type args struct {
		arr []int64
		num int64
	}

	eles := []int64{1,2,3,4,5,6,7,8}

	results := make([][]int64,0)
	results = append(results, []int64{1,2,3})
	results = append(results, []int64{4,5,6})
	results = append(results, []int64{7,8})

	tests := []struct {
		name string
		args args
		want [][]int64
	}{
		// TODO: Add test cases.
		{
			name: "test for int64",
			args: args{
				arr: eles,
				num: 3,
			},
			want: results,
		},
		{
			name: "test for int64",
			args: args{
				arr: eles[0:3],
				num: 3,
			},
			want: results[:1],
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got [][]int64
			if got = Int64ArraySplit(tt.args.arr, tt.args.num); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ArraySplit() = %v, want %v", got, tt.want)
			}
			fmt.Println(got)
		})
	}
}
