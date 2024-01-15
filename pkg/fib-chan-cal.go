package pkg

type Calc struct {
	Pos    int64
	First  int64
	Second int64
}

type Input chan Calc
type Output chan Calc

func FibChan(pos int64) int64 {
	in := make(Input)
	out := make(Output)
	defer close(in)
	defer close(out)

	// Seed channel with first value.
	go func() {
		calc := Calc{
			pos,
			0,
			1,
		}
		in <- calc
	}()

	in.fibChanProcessor(out)
	res := <-out

	return res.First
}

func (in Input) fibChanProcessor(out Output) {
	go func() {
		for {
			select {
			case calc := <-in:
				if calc.Pos == 0 {
					out <- calc
					return
				}
				// Because chan points to value you must point to
				// new value or you will modify the orinal.
				go func() {
					oCalc := Calc{
						calc.Pos - 1,
						calc.Second,
						calc.Second + calc.First,
					}
					in <- oCalc
				}()
			}
		}
	}()
}
