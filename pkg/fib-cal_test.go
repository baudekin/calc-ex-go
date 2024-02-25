package pkg

import "testing"

var loopCnt = 10000000

type args struct {
	pos uint64
}

var tests = []struct {
	name string
	args args
	want uint64
}{
	{
		name: "One",
		args: args{
			pos: 1,
		},
		want: 1,
	},
	{
		name: "Two",
		args: args{
			pos: 2,
		},
		want: 1,
	},
	{
		name: "Three",
		args: args{
			pos: 3,
		},
		want: 2,
	},
	{
		name: "Fifty",
		args: args{
			pos: 50,
		},
		want: 12586269025,
	},
	{
		name: "Hundred",
		args: args{
			pos: 100,
		},
		want: 3736710778780434371,
	},
	{
		name: "Three Hundred",
		args: args{
			pos: 300,
		},
		want: 17658870469870104080,
	},
	{
		name: "Five Hundred",
		args: args{
			pos: 500,
		},
		want: 2171430676560690477,
	},
	{
		name: "One Thousand",
		args: args{
			pos: 1000,
		},
		want: 817770325994397771,
	},
}

func TestFibForLoop(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for i := 1; i < loopCnt; i++ {
				{
					if got := FibForLoop(tt.args.pos); got != tt.want {
						t.Errorf("FibForLoop() = %v, want %v", got, tt.want)
					}
				}
			}
		})
	}
}

func TestFibRecursion(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for i := 1; i < loopCnt; i++ {
				{
					if got := FibRecursion(tt.args.pos); got != tt.want {
						t.Errorf("FibForLoop() = %v, want %v", got, tt.want)
					}
				}
			}
		})
	}
}

func TestFibTailRecusion(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for i := 1; i < loopCnt; i++ {
				{
					if got := FibTailRecusion(tt.args.pos); got != tt.want {
						t.Errorf("FibForLoop() = %v, want %v", got, tt.want)
					}
				}
			}
		})
	}
}
