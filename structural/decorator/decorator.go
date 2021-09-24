package decorator

type Component interface {
	Calc() int
}

type ConcreteComponent struct{}

func (c *ConcreteComponent) Calc() int {
	return 0
}

type MulDecorator struct {
	Component
	num int
}

func (m *MulDecorator) Calc() int {
	return m.Component.Calc() * m.num
}

type AddDecorator struct {
	Component
	num int
}

func (a *AddDecorator) Calc() int {
	return a.Component.Calc() + a.num
}

func WarpMulDecorator(c Component, num int) Component {
	return &MulDecorator{
		Component: c,
		num:       num,
	}
}

func WarpAddDecorator(c Component, num int) Component {
	return &AddDecorator{
		Component: c,
		num:       num,
	}
}
