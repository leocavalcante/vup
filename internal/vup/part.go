package vup

type Inc interface {
	Inc(int) error
}

type Dec interface {
	Dec(int) error
}

type Clear interface {
	Clear()
}

type Part interface {
	Inc
	Dec
	Clear
	Value() int
	Set(int)
}
