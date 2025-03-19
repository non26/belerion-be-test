package model

type Factor struct {
	Number string
	Power  string
}

func NewFactor(number string, power string) Factor {
	return Factor{Number: number, Power: power}
}
