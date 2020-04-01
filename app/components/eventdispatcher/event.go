package eventdispatcher

type EventInterface interface {
	GetName() string
	GetData() interface{}
}
