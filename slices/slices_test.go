package main

import (
	"reflect"
	"testing"
)

func Test_concat(t *testing.T) {
	type args struct {
		s1 []string
		s2 []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "concat",
			args: args{
				[]string{"A", "B"}, []string{"C", "D", "E"},
			},
			want: []string{"A", "B", "C", "D", "E"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := concat(tt.args.s1, tt.args.s2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("concat() = %v, want %v", got, tt.want)
			}
		})
	}
}
