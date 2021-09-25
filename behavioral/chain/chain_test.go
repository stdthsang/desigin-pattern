package chain

func ExampleChain() {
	cashier := &Cashier{}

	//Set next for medical department
	medical := &Medical{}
	medical.SetNext(cashier)

	//Set next for doctor department
	doctor := &Doctor{}
	doctor.SetNext(medical)

	//Set next for reception department
	reception := &Reception{}
	reception.SetNext(doctor)

	patient := &Patient{Name: "abc"}
	//Patient visiting
	reception.Execute(patient)

	// Output:
	// Reception registering patient
	// Doctor checking patient
	// Medical giving medicine to patient
	// Cashier getting money from patient patient
}
