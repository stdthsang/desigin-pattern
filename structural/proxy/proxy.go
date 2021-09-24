package proxy

type Subject interface {
	Do() string
}

type RealSubject struct{}

func (r *RealSubject) Do() string {
	return "real"
}

type Proxy struct {
	real RealSubject
}

func (p *Proxy) Do() string {
	var res string
	res += "pre:"
	res += p.real.Do()
	res += ":after"
	return res
}
