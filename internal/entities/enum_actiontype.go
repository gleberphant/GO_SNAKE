package entities

type ActionEnum int

func (a ActionEnum) String() string {
	return [...]string{"RotateLeft", "RotateRight", "Move"}[a]
}

const (
	ROTATE_LEFT ActionEnum = iota
	ROTATE_RIGHT
	MOVE
	STOP
	MOVE_FOWARD
	MOVE_BACK
)
