package eventdispatcher

type Listener struct {
	async    bool
	onListen func(event EventInterface) error
}

//新建listener
func NewListener(onListen func(event EventInterface) error, async bool) *Listener {
	return &Listener{async: async, onListen: onListen}
}
