package entities

import (
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type EntityItem struct {
	EntityPrimitive
	ItemType ItemEnum
}

func (e *EntityItem) Load() {
	e.EntityPrimitive.Load()
	e.Color = rl.Gray
}

func (e *EntityItem) Update() {

}

func (e *EntityItem) Collision(target EntityChar) {
	target.life++
	e.IsAlive = false
	fmt.Println("Colidiu %d", target.life)

}
