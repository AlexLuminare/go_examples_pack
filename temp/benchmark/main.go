package main

type str struct {
	a string
}

var (
	a1 = str{a: "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
	a2 = str{a: "bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb"}
)

func swapValuesP(first, second *str) (*str, *str) {
	temp := *first
	*first = *second
	*second = temp
	return first, second
}

func swapValuesV(first, second str) (str, str) {
	temp := first
	first = second
	second = temp
	return first, second
}
