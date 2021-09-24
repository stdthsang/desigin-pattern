package abstract_factory

import "fmt"

type OrderMain interface {
	SaveOrderMain()
}

type OrderDetail interface {
	SaveOrderDetail()
}

type OrderFactory interface {
	CreateOrderMain() OrderMain
	CreateOrderDetail() OrderDetail
}

type RDBMain struct{}

func (r *RDBMain) SaveOrderMain() {
	fmt.Print("rdb main save\n")
}

type RDBDetail struct{}

func (r *RDBDetail) SaveOrderDetail() {
	fmt.Print("rdb detail save\n")
}

type RDBFactory struct{}

func (r *RDBFactory) CreateOrderMain() OrderMain {
	return &RDBMain{}
}

func (r *RDBFactory) CreateOrderDetail() OrderDetail {
	return &RDBDetail{}
}

type XMLMain struct{}

func (x *XMLMain) SaveOrderMain() {
	fmt.Print("xml main save\n")
}

type XMLDetail struct{}

func (x *XMLDetail) SaveOrderDetail() {
	fmt.Print("xml detail save\n")
}

type XMLFactory struct{}

func (x *XMLFactory) CreateOrderMain() OrderMain {
	return &XMLMain{}
}

func (x *XMLFactory) CreateOrderDetail() OrderDetail {
	return &XMLDetail{}
}
