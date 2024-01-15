package pkg

func FibForLoop(pos int64) int64 {
	var res, ind, first, second int64

	for ind, first, second = 0, 0, 1; ind <= pos; ind, first, second = ind+1, first+second, first {
		if ind == pos {
			res = first
		}
	}

	return res
}

func FibRecursion(pos int64) int64 {
	if pos < 2 {
		return pos
	}
	return FibRecursion(pos-2) + FibRecursion(pos-1)
}

func FibTailRecusion(pos int64) int64 {
	return fibTail(pos, 0, 1)
}

func fibTail(pos int64, first int64, second int64) int64 {
	if pos == 0 {
		return first
	}
	return fibTail(pos-1, second, first+second)
}
