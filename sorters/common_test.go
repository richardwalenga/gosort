package sorters

import "testing"

func TestAreIdentical(t *testing.T) {
	type args struct {
		x []int
		y []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Empty are identical", args{[]int{}, []int{}}, true},
		{"Non-empty identical are identical", args{[]int{1, 2, 3, 4}, []int{1, 2, 3, 4}}, true},
		{"Non-empty of different lengths are not identical", args{[]int{1, 2, 3, 4}, []int{1, 2, 3}}, false},
		{"Non-empty of same lengths but different order are not identical", args{[]int{1, 2, 3, 4}, []int{1, 2, 4, 3}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AreIdentical(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("AreIdentical() = %v, want %v", got, tt.want)
			}
		})
	}
}
