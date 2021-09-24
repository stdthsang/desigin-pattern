package builder

import "testing"

func TestBuilder1(t *testing.T) {
	builder := &Builder1{}
	director := NewDirector(builder)
	director.Construct()
	if builder.result != "123" {
		t.Fatalf("Builder1 fail expect 123 acture %s", builder.result)
	}
}

func TestBuilder2(t *testing.T) {
	builder := &Builder2{}
	director := NewDirector(builder)
	director.Construct()
	if builder.result != 6 {
		t.Fatalf("Builder2 fail expect 6 acture %d", builder.result)
	}
}
