package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

var input = "3,4,3,1,2"

func TestPart1(t *testing.T) {
	require.Equal(t, 26, Part1(input, 18))
}

func TestPart1Print(t *testing.T) {
	fmt.Println(Part1("4", 18))
}


//func TestCalculateSpawn(t *testing.T) {
//	require.Equal(t, 3, CalculateSpawn(3, 18))
//}

func TestCalculateSpawn(t *testing.T) {
	type args struct {
		init int
		days int
	}
	tests := []struct {
		args args
		want int
	}{
		{args{0, 18}, 3},
		{args{1, 18}, 3},
		{args{2, 18}, 3},
		{args{3, 18}, 3},
		{args{4, 18}, 2},
		{args{5, 18}, 2},
		{args{6, 18}, 2},
	}
		for i, tt := range tests {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			require.Equal(t, tt.want, CalculateSpawn(tt.args.init, tt.args.days))
		})
	}
}
func TestCalculateTotalSpawn(t *testing.T) {
	type args struct {
		init int
		days int
	}
	tests := []struct {
		args args
		want uint64
	}{
		{args{0, 18}, 6},
		{args{1, 18}, 6},
		{args{2, 18}, 4},
		{args{3, 18}, 4},
		{args{4, 18}, 3},
		{args{5, 18}, 3},
		{args{6, 18}, 3},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			require.Equal(t, tt.want, CalculateTotalSpawn(tt.args.init, tt.args.days))
		})
	}
}

var day6 = `1,4,3,3,1,3,1,1,1,2,1,1,1,4,4,1,5,5,3,1,3,5,2,1,5,2,4,1,4,5,4,1,5,1,5,5,1,1,1,4,1,5,1,1,1,1,1,4,1,2,5,1,4,1,2,1,1,5,1,1,1,1,4,1,5,1,1,2,1,4,5,1,2,1,2,2,1,1,1,1,1,5,5,3,1,1,1,1,1,4,2,4,1,2,1,4,2,3,1,4,5,3,3,2,1,1,5,4,1,1,1,2,1,1,5,4,5,1,3,1,1,1,1,1,1,2,1,3,1,2,1,1,1,1,1,1,1,2,1,1,1,1,2,1,1,1,1,1,1,4,5,1,3,1,4,4,2,3,4,1,1,1,5,1,1,1,4,1,5,4,3,1,5,1,1,1,1,1,5,4,1,1,1,4,3,1,3,3,1,3,2,1,1,3,1,1,4,5,1,1,1,1,1,3,1,4,1,3,1,5,4,5,1,1,5,1,1,4,1,1,1,3,1,1,4,2,3,1,1,1,1,2,4,1,1,1,1,1,2,3,1,5,5,1,4,1,1,1,1,3,3,1,4,1,2,1,3,1,1,1,3,2,2,1,5,1,1,3,2,1,1,5,1,1,1,1,1,1,1,1,1,1,2,5,1,1,1,1,3,1,1,1,1,1,1,1,1,5,5,1`
func TestPart2(t *testing.T) {
	require.EqualValues(t, 26, Part2(input, 18))
	require.EqualValues(t, 5934, Part2(input, 80))
	require.EqualValues(t, 26984457539, Part2(input, 256))
	require.EqualValues(t, 379114, Part2(day6, 80))
}