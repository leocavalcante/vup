package vup

type Inc interface {
	Inc(int)
}

type Dec interface {
	Dec(int) error
}

type Part interface {
	Inc
	Dec
}
