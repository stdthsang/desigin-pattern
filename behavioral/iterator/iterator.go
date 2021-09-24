package iterator

import "fmt"

type Aggregate interface {
	Iterator() Iterator
}

type Iterator interface {
	First()
	IsDone() bool
	Next() interface{}
}

type Numbers struct {
	start int
	end   int
}

func NewNumbers(start, end int) *Numbers {
	return &Numbers{
		start: start,
		end:   end,
	}
}

func (n *Numbers) Iterator() Iterator {
	return &NumberIterator{
		numbers: n,
		next:    n.start,
	}
}

type NumberIterator struct {
	numbers *Numbers
	next    int
}

func (i *NumberIterator) First() {
	i.next = i.numbers.start
}

func (i *NumberIterator) IsDone() bool {
	return i.next > i.numbers.end
}

func (i *NumberIterator) Next() interface{} {
	if i.IsDone() {
		return nil
	}
	next := i.next
	i.next++
	return next
}

func IteratorPrint(i Iterator) {
	for i.First(); !i.IsDone(); {
		c := i.Next()
		fmt.Printf("%#v\n", c)
	}
}
