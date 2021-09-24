package adapter

type Adaptee interface {
	SpecificRequest() string
}

func NewAdaptee() Adaptee {
	return &adapterImpl{}
}

type adapterImpl struct{}

func (a *adapterImpl) SpecificRequest() string {
	return "adaptee method"
}

type Target interface {
	Request() string
}

type adapter struct {
	Adaptee
}

func NewAdapter(adaptee Adaptee) Target {
	return &adapter{Adaptee: adaptee}
}

func (a *adapter) Request() string {
	return a.SpecificRequest()
}
