package command

import "fmt"

type MotherBoard struct{}

func (*MotherBoard) Start() {
	fmt.Print("system starting\n")
}

func (*MotherBoard) Reboot() {
	fmt.Print("system rebooting\n")
}

type Command interface {
	Execute()
}

type StartCmd struct {
	mb *MotherBoard
}

func NewStartCmd(mb *MotherBoard) *StartCmd {
	return &StartCmd{
		mb: mb,
	}
}

func (c *StartCmd) Execute() {
	c.mb.Start()
}

type RebootCmd struct {
	mb *MotherBoard
}

func NewRebootCmd(mb *MotherBoard) *RebootCmd {
	return &RebootCmd{
		mb: mb,
	}
}

func (c *RebootCmd) Execute() {
	c.mb.Reboot()
}

type Box struct {
	button1 Command
	button2 Command
}

func NewBox(button1, button2 Command) *Box {
	return &Box{
		button1: button1,
		button2: button2,
	}
}

func (b *Box) PressButton1() {
	b.button1.Execute()
}

func (b *Box) PressButton2() {
	b.button2.Execute()
}
