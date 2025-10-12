package scenes

// typedef identificado da cena atual
type SceneEnum int

func (s SceneEnum) String() string {
	return [...]string{"Intro", "Menu", "Play", "Over", "Quit"}[s]
}

const (
	INTRO SceneEnum = iota
	MENU
	PLAY
	OVER
	QUIT
)
