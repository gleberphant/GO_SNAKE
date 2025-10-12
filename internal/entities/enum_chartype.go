package entities

type CharEnum int

func (c CharEnum) String() string {
	return [...]string{"player", "enemy"}[c]
}
