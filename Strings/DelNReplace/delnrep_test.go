package delnrep

import (
	"reflect"
	"testing"
)

func Test_delnrep(t *testing.T) {
	type args struct {
		arr [10]int
	}
	tests := []struct {
		name string
		args args
		want [10]int
	}{
		{"Input 1", args{[10]int{1, 2, 1, 3}}, [10]int{4, 4, 4, 4, 3}},
		{"Input 2", args{[10]int{1, 2, 3, 1}}, [10]int{4, 4, 3, 4, 4}},
		{"Input 3", args{[10]int{2, 1, 3, 1}}, [10]int{4, 4, 3, 4, 4}},
		{"Input 4", args{[10]int{6, 7, 8, 3, 2}}, [10]int{6, 7, 8, 3}},
		{"Input 5", args{[10]int{1, 3, 1, 1}}, [10]int{4, 4, 3, 4, 4, 4, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := delnrep(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("delnrep() = %v, want %v", got, tt.want)
			}
		})
	}
}
