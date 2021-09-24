package simple_factory

import "fmt"

type API interface {
	Say(name string) string
}

func NewAPI(t int) API {
	switch t {
	case 1:
		return &HiAPI{}
	case 2:
		return &HelloAPI{}
	}
	return nil
}

type HiAPI struct{}

func (h *HiAPI) Say(name string) string {
	return fmt.Sprintf("Hi, %s", name)
}

type HelloAPI struct{}

func (h *HelloAPI) Say(name string) string {
	return fmt.Sprintf("Hello, %s", name)
}
