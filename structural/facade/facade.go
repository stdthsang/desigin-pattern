package facade

import "fmt"

type API interface {
	Test() string
}

func NewAPI() API {
	return &apiImpl{
		a: NewAModuleAPI(),
		b: NewBModuleAPI(),
	}
}

type apiImpl struct {
	a AModuleAPI
	b BModuleAPI
}

func (impl *apiImpl) Test() string {
	aRet, bRet := impl.a.TestA(), impl.b.TestB()
	return fmt.Sprintf("%s\n%s", aRet, bRet)
}

type AModuleAPI interface {
	TestA() string
}

func NewAModuleAPI() AModuleAPI {
	return &aModuleImpl{}
}

type aModuleImpl struct{}

func (a *aModuleImpl) TestA() string {
	return "A module running"
}

type BModuleAPI interface {
	TestB() string
}

func NewBModuleAPI() BModuleAPI {
	return &bModuleImpl{}
}

type bModuleImpl struct{}

func (b *bModuleImpl) TestB() string {
	return "B module running"
}
