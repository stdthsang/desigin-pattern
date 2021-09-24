package simple_factory

import "testing"

const (
	hiAPI    = 1
	helloAPI = 2
)

func TestHiAPI_Say(t *testing.T) {
	api := NewAPI(hiAPI)
	s := api.Say("Tom")
	if s != "Hi, Tom" {
		t.Fatal("TestHiAPI_Say test fail")
	}
}

func TestHelloAPI_Say(t *testing.T) {
	api := NewAPI(helloAPI)
	s := api.Say("Tom")
	if s != "Hello, Tom" {
		t.Fatal("TestHelloAPI_Say test fail")
	}
}
