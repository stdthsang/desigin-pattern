package factory_method

type Operator interface {
	SetA(int)
	SetB(int)
	Result() int
}

type OperatorFactory interface {
	Create() Operator
}

type BaseOperator struct {
	a, b int
}

func (o *BaseOperator) SetA(a int) {
	o.a = a
}

func (o *BaseOperator) SetB(b int) {
	o.b = b
}

type PlusOperatorFactory struct{}

func (p *PlusOperatorFactory) Create() Operator {
	return &PlusOperator{BaseOperator: &BaseOperator{}}
}

type PlusOperator struct {
	*BaseOperator
}

func (o *PlusOperator) Result() int {
	return o.a + o.b
}

type MinusOperatorFactory struct{}

func (m *MinusOperatorFactory) Create() Operator {
	return &MinusOperator{BaseOperator: &BaseOperator{}}
}

type MinusOperator struct {
	*BaseOperator
}

func (m *MinusOperator) Result() int {
	return m.a - m.b
}
