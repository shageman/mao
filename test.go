package mao

type Test struct {
	Title string
	Fn func(Expect)
}

func (test *Test) Run () {
	test.Fn(func (val interface{}) *Expected {
		return &Expected{ test, val }
	})
}
