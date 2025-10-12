package entities

import rl "github.com/gen2brain/raylib-go/raylib"

type Entity interface {
	Load()
	GetPositionVector() rl.Vector2
	GetRect() rl.Rectangle
	Control(action ActionEnum)
	Collision()
	Update()
	Draw()
}
