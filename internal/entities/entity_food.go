package entities

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type EntityFood struct {
	EntityPrimitive
	ItemType ItemEnum
}

func (e *EntityFood) Load() {
	e.EntityPrimitive.Load()
	e.Color = rl.Gray
}

func (e *EntityFood) Update() {

}

func (e *EntityFood) Collision(target *EntitySnake) {
	if e.IsAlive && target.IsAlive {
		target.EatFood()
	}
	
	e.IsAlive = false

}
