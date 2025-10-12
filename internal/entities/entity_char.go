package entities

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type EntityChar struct {
	EntityPrimitive
	CharType CharEnum
	life     int
}

func (e *EntityChar) Load() {
	e.EntityPrimitive.Load()
	e.Color = rl.Green
}

func (e *EntityChar) Update() {
	e.Move()
}

func (e *EntityChar) Collision(target EntityChar) {}
