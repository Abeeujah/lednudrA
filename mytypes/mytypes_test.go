package mytypes_test

import (
	"mytypes"
	"testing"
)

func TestTwice(t *testing.T) {
	t.Parallel()
	input := mytypes.MyInt(9)

	want := mytypes.MyInt(18)

	got := input.Twice()

	if want != got {
		t.Errorf("Expected %d, got %d\n", want, got)
	}
}

func TestMyStringLen(t *testing.T) {
	t.Parallel()

	input := mytypes.MyString("Abraham")

	want := 7

	got := input.Len()

	if want != got {
		t.Errorf("Expected %d, got %d\n", want, got)
	}
}

func TestMyBuilder(t *testing.T) {
	t.Parallel()

	var mb mytypes.MyBuilder

	mb.Contents.WriteString("Hello, ")
	mb.Contents.WriteString("Gophers!")

	want := "Hello, Gophers!"

	got := mb.Contents.String()

	if want != got {
		t.Errorf("Expected %q, got %q\n", want, got)
	}

	wantLen := 15
	gotLen := mb.Contents.Len()

	if wantLen != gotLen {
		t.Errorf("%q: Expected %d, got %d\n", mb.Contents.String(), wantLen, gotLen)
	}
}

func TestMyBuilderLen(t *testing.T) {
	t.Parallel()

	var mb mytypes.MyBuilder
	mb.Contents.WriteString("Hello, Gophers!")

	want := 15
	got := mb.Contents.Len()

	if want != got {
		t.Errorf("Expected %d, got %d\n", want, got)
	}
}

func TestMyBuilderHello(t *testing.T) {
	t.Parallel()

	var mb mytypes.MyBuilder
	want := "Hello, Gophers!"
	got := mb.Hello()

	if want != got {
		t.Errorf("Expected %q, got %q\n", want, got)
	}
}

func TestStringUpperCaser(t *testing.T) {
	var uc mytypes.StringUpperCaser
	uc.Content.WriteString("hello")

	want := "HELLO"
	got := uc.ToUpper()

	if want != got {
		t.Errorf("Expected %q, got %q\n", want, got)
	}
}

func TestDouble(t *testing.T) {
	t.Parallel()

	x := mytypes.MyInt(12)
	want := mytypes.MyInt(24)

	p := &x
	p.Double()

	if want != x {
		t.Errorf("Expected %d, got %d", want, p)
	}
}
