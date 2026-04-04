type FooBar struct {
	n     int
	fooCh chan struct{}
	barCh chan struct{}
}

func NewFooBar(n int) *FooBar {
	fb := &FooBar{
		n:     n,
		fooCh: make(chan struct{}, 1),
		barCh: make(chan struct{}, 1),
	}
	fb.fooCh <- struct{}{} // start with Foo
	return fb
}

func (fb *FooBar) Foo(printFoo func()) {
	for i := 0; i < fb.n; i++ {
		<-fb.fooCh          // wait for turn
		printFoo()
		fb.barCh <- struct{}{} // signal Bar
	}
}

func (fb *FooBar) Bar(printBar func()) {
	for i := 0; i < fb.n; i++ {
		<-fb.barCh          // wait for turn
		printBar()
		fb.fooCh <- struct{}{} // signal Foo
	}
}