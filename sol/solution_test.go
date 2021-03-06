package sol

import "testing"

func BenchmarkTest(b *testing.B) {
	nums := []int{2, 7, 9, 3, 1}
	for idx := 0; idx < b.N; idx++ {
		rob(nums)
	}
}
func Test_rob(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "nums = [1,2,3,1]",
			args: args{nums: []int{1, 2, 3, 1}},
			want: 4,
		},
		{
			name: "nums = [2,7,9,3,1]",
			args: args{nums: []int{2, 7, 9, 3, 1}},
			want: 12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rob(tt.args.nums); got != tt.want {
				t.Errorf("rob() = %v, want %v", got, tt.want)
			}
		})
	}
}
