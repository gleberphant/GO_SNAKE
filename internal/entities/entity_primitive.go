package entities

import (
	"go_snake/internal/tiles"
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type EntityPrimitive struct {
	PosX      int32
	PosY      int32
	Size      int32
	Direction float64
	Speed     float64
	IsAlive   bool
	Color     rl.Color
}

func (e *EntityPrimitive) Load() {

	e.Size = tiles.TSIZE / 2
	e.PosX = 0
	e.PosY = 0

	//	e.Speed = tiles.TSIZE / 4
	e.Direction = 45

	e.Color = rl.DarkBlue
	e.IsAlive = true
}

func (e *EntityPrimitive) Control(action ActionEnum) {

	// muda o angulo do movimento
	switch action {
	case ROTATE_LEFT:
		e.Direction += 45.0
	case ROTATE_RIGHT:
		e.Direction -= 45.0
	case MOVE:
		e.Speed = tiles.TSIZE / 4
	case STOP:
		e.Speed = 0
	default:
		e.Speed = 0
	}

}

func (e *EntityPrimitive) Update() {

}

func (e *EntityPrimitive) Move() {
	// Atualiza Position
	e.PosX += int32(e.Speed * math.Sin(e.Direction*rl.Deg2rad))
	e.PosY += int32(e.Speed * math.Cos(e.Direction*rl.Deg2rad))

	// Limpa Parametros das ações
	e.Speed = 0

}

func (e *EntityPrimitive) Collision() {}

func (e *EntityPrimitive) Draw() {
	rl.DrawCircleGradient(e.PosX, e.PosY, float32(e.Size), rl.Black, e.Color)
}

// Getters

func (e *EntityPrimitive) GetRect() rl.Rectangle {
	return rl.Rectangle{X: float32(e.PosY), Y: float32(e.PosY), Width: float32(e.PosX + e.Size), Height: float32(e.PosY + e.Size)}
}

func (e *EntityPrimitive) GetPositionVector() rl.Vector2 {
	return rl.Vector2{X: float32(e.PosX), Y: float32(e.PosY)}
}
