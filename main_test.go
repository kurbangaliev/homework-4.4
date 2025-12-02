package main

import "testing"

func Test_sumAll(t *testing.T) {
	type args struct {
		numbers []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"sumAll_1", args{[]int{1, 2, 3, 4, 5}}, 15},
		{"sumAll_2", args{[]int{-1, -2, -3, -4, 10}}, 0},
		{"sumAll_3", args{[]int{1, -1, 2, -2, 0}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumAll(tt.args.numbers...); got != tt.want {
				t.Errorf("sumAll() = %v, want %v", got, tt.want)
			}
		})
	}
}
