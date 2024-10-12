package mutex

type Mutex struct {
	exitCh chan interface{}
}

func New() *Mutex {
	return &Mutex{
		exitCh: make(chan interface{}),
	}
}

func (m *Mutex) Done() {
	close(m.exitCh)
}

func (m *Mutex) Wait() {
	<-m.exitCh
}
