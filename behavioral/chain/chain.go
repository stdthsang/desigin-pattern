package chain

import "fmt"

type Department interface {
	Execute(*Patient)
	SetNext(Department)
}

type Patient struct {
	Name              string
	RegistrationDone  bool
	DoctorCheckUpDone bool
	MedicineDone      bool
	PaymentDone       bool
}

type Reception struct {
	next Department
}

func (r *Reception) Execute(patient *Patient) {
	if patient.RegistrationDone {
		fmt.Print("Patient registration already done")
		r.next.Execute(patient)
		return
	}
	fmt.Print("Reception registering patient\n")
	patient.RegistrationDone = true
	r.next.Execute(patient)
}

func (r *Reception) SetNext(next Department) {
	r.next = next
}

type Doctor struct {
	next Department
}

func (d *Doctor) Execute(patient *Patient) {
	if patient.DoctorCheckUpDone {
		fmt.Print("Doctor checkup already done")
		d.next.Execute(patient)
		return
	}
	fmt.Print("Doctor checking patient\n")
	patient.DoctorCheckUpDone = true
	d.next.Execute(patient)
}

func (d *Doctor) SetNext(next Department) {
	d.next = next
}

type Medical struct {
	next Department
}

func (m *Medical) Execute(patient *Patient) {
	if patient.MedicineDone {
		fmt.Print("Medicine already given to patient")
		m.next.Execute(patient)
		return
	}
	fmt.Print("Medical giving medicine to patient\n")
	patient.MedicineDone = true
	m.next.Execute(patient)
}

func (m *Medical) SetNext(next Department) {
	m.next = next
}

type Cashier struct {
	next Department
}

func (c *Cashier) Execute(patient *Patient) {
	if patient.PaymentDone {
		fmt.Print("Payment Done")
		return
	}
	fmt.Print("Cashier getting money from patient patient\n")
}

func (c *Cashier) SetNext(next Department) {
	c.next = next
}
