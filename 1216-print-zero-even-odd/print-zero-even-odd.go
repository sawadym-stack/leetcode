type ZeroEvenOdd struct {
	n      int
	zeroCh chan struct{}
	oddCh  chan struct{}
	evenCh chan struct{}
}

func NewZeroEvenOdd(n int) *ZeroEvenOdd {
	z := &ZeroEvenOdd{
		n:      n,
		zeroCh: make(chan struct{}, 1),
		oddCh:  make(chan struct{}, 1),
		evenCh: make(chan struct{}, 1),
	}
	z.zeroCh <- struct{}{} // start with zero
	return z
}

func (z *ZeroEvenOdd) Zero(printNumber func(int)) {
	for i := 1; i <= z.n; i++ {
		<-z.zeroCh
		printNumber(0)

		if i%2 == 1 {
			z.oddCh <- struct{}{}
		} else {
			z.evenCh <- struct{}{}
		}
	}
}

func (z *ZeroEvenOdd) Odd(printNumber func(int)) {
	for i := 1; i <= z.n; i += 2 {
		<-z.oddCh
		printNumber(i)
		z.zeroCh <- struct{}{}
	}
}

func (z *ZeroEvenOdd) Even(printNumber func(int)) {
	for i := 2; i <= z.n; i += 2 {
		<-z.evenCh
		printNumber(i)
		z.zeroCh <- struct{}{}
	}
}