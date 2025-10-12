package entities

type ItemEnum int

func (i ItemEnum) String() string {
	return [...]string{"Food", "Trap"}[i]
}

const (
	FOOD ItemEnum = iota
	TRAP
)
