package abstract_factory

func getMainAndDetail(factory OrderFactory) {
	factory.CreateOrderMain().SaveOrderMain()
	factory.CreateOrderDetail().SaveOrderDetail()
}

func ExampleRDBFactory() {
	factory := &RDBFactory{}
	getMainAndDetail(factory)
	// Output:
	// rdb main save
	// rdb detail save
}

func ExampleXMLFactory() {
	factory := &XMLFactory{}
	getMainAndDetail(factory)
	// Output:
	// xml main save
	// xml detail save
}
