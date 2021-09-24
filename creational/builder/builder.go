package builder

type Builder interface {
	Part1()
	Part2()
	Part3()
}

type Director struct {
	Builder
}

func NewBuilder(builder Builder) *Director {
	return &Director{Builder: builder}
}

func (d *Director) Construct() {
	d.Builder.Part1()
	d.Builder.Part2()
	d.Builder.Part3()
}

type Builder1 struct {
	result string
}

func (b *Builder1) Part1() {
	b.result += "1"
}

func (b *Builder1) Part2() {
	b.result += "2"
}

func (b *Builder1) Part3() {
	b.result += "3"
}

func (b *Builder1) GetResult() string {
	return b.result
}

type Builder2 struct {
	result int
}

func (b *Builder2) Part1() {
	b.result += 1
}

func (b *Builder2) Part2() {
	b.result += 2
}

func (b *Builder2) Part3() {
	b.result += 3
}

func (b *Builder2) GetResult() int {
	return b.result
}
