package mytypes

/*
	func (i int) Twice() int {
		return i * 2
	}

	because int is a non-local type; that is, it's not defined in our package
	we cannot define a method on it
*/

type MyInt int

func (i MyInt) Twice() MyInt {
	return i * 2
}

type MyString string

func (s MyString) Len() int {
	return len(s)
}
