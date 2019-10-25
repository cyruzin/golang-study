package main

import "testing"

func Product(x int, y int) int {
	return x * y
}

func TestProduct(t *testing.T) {
	// Product of positive numbers
	test1 := Product(4, 5)
	if test1 != 20 {
		t.Errorf("Product(4,5) = %d; expected 20", test1)
	}

	// Product of a positive number and 0
	test2 := Product(4, 0)
	if test2 != 0 {
		t.Errorf("Product(4,0) = %d; expected 0", test2)
	}

	// Product of a positive and a negative number
	test3 := Product(4, -5)
	if test3 != -20 {
		t.Errorf("Product(4,-5) = %d; expected -20", test3)
	}
}

func Add(x int, y int) int {
	return x + y
}

func TestAdd(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "both positive numbers",
			args: args{
				x: 1,
				y: 1,
			},
			want: 2,
		},
		{
			name: "both negative numbers",
			args: args{
				x: -2,
				y: -3,
			},
			want: -5,
		},
		{
			name: "one negative, one positive",
			args: args{
				x: -2,
				y: 6,
			},
			want: 4,
		},
		{
			name: "one negative, one positive",
			args: args{
				x: 0,
				y: 0,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Add(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}
