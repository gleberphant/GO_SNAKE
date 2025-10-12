package entities

import (
	"go_snake/internal/tiles"
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Entity struct {
	TileX     int32
	TileY     int32
	Direction float64
	Speed     float64

	Color rl.Color
}

func (e *Entity) Load() {

	e.Speed = tiles.TSIZE / 4
	e.Direction = 45
	e.TileX = 16
	e.TileY = 16
	e.Color = rl.DarkBlue

}

func (e *Entity) ChangeDirection(newDirection float64) {
	e.Direction += newDirection

}

func (e *Entity) UpdatePosition() {

	e.TileX += int32(e.Speed * math.Sin(e.Direction*rl.Deg2rad))
	e.TileY += int32(e.Speed * math.Cos(e.Direction*rl.Deg2rad))

}

func (e *Entity) Draw() {
	rl.DrawCircle((e.TileX), (e.TileY), tiles.TSIZE/2, e.Color)
}
