package proto

type Cloneable interface {
	Clone() Cloneable
}

type ProtoManager struct {
	proto map[string]Cloneable
}

func NewProtoManager() *ProtoManager {
	return &ProtoManager{proto: make(map[string]Cloneable)}
}

func (p *ProtoManager) Get(name string) Cloneable {
	return p.proto[name].Clone()
}

func (p *ProtoManager) Set(name string, proto Cloneable) {
	p.proto[name] = proto
}
