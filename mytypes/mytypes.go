package mytypes

import "strings"

type MyInt int
type MyString string
type MyBuilder struct {
	Contents strings.Builder
}
type StringUpperCaser struct {
	Content strings.Builder
}

func (i MyInt) Twice() MyInt {
	return i * 2
}

func (s MyString) Len() int {
	return len(s)
}

func (mb MyBuilder) Len() int {

	return mb.Contents.Len()
}

func (mb MyBuilder) Hello() string {
	return "Hello, Gophers!"
}

func (uc StringUpperCaser) ToUpper() string {
	return strings.ToUpper(uc.Content.String())
}

func (m *MyInt) Double() {
	*m *= MyInt(2)
}
