package entities

import (
	"container/list"
	"fmt"
	"go_snake/internal/tiles"
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type EntityChar struct {
	EntityPrimitive
	CharType CharEnum
	life     int
	bodyPos  *list.List
}

func (e *EntityChar) Load() {
	e.EntityPrimitive.Load()
	e.Color = rl.Green
	e.life = 1
	e.bodyPos = list.New()

	body := rl.Vector2{X: float32(e.PosX), Y: float32(e.PosY)}
	fmt.Printf("Adicionado cabeça X: %d Y %d\n", e.PosX, e.PosY)

	e.bodyPos.PushBack(&body)

	e.Speed = tiles.TSIZE / 2

}

func (e *EntityChar) Update() {

	// se a vida for maior que quantidade de corpos, adiciona corpo
	if e.bodyPos.Len() < e.life {
		body := rl.Vector2{X: float32(e.PosX), Y: float32(e.PosY)}
		fmt.Printf("corpo cresceu  X: %d Y %d\n", e.PosX, e.PosY)
		e.bodyPos.PushBack(&body)

	}

	// movimenta personagem
	oldPosX := e.PosX
	oldPosY := e.PosY

	e.PosX += int32(e.Speed * math.Sin(e.Direction*rl.Deg2rad))
	e.PosY += int32(e.Speed * math.Cos(e.Direction*rl.Deg2rad))

	// atualiza a cabeça com posicao antiga
	nodeLast := e.bodyPos.Back()
	nodePrev := nodeLast.Prev()

	// registra posição anterior
	for nodePrev != nil {
		// Create a new Vector2 with the values from nodePrev
		newVec := rl.Vector2{X: nodePrev.Value.(*rl.Vector2).X, Y: nodePrev.Value.(*rl.Vector2).Y}
		nodeLast.Value = &newVec

		nodeLast = nodePrev
		nodePrev = nodePrev.Prev()
	}

	head, ok := nodeLast.Value.(*rl.Vector2)

	// atualiza cabeça
	if ok {
		head.X = float32(oldPosX)
		head.Y = float32(oldPosY)
	}
	// Limpa Parametros das ações
	e.Speed = 0

	// e.Move()
}

func (e *EntityChar) Draw() {

	for node := e.bodyPos.Front(); node != nil; node = node.Next() {

		body, ok := node.Value.(*rl.Vector2)
		if !ok {
			continue
		}

		rl.DrawCircleLines(int32(body.X), int32(body.Y), float32(e.Size), rl.Blue)

	}

	rl.DrawCircle(e.PosX, e.PosY, float32(e.Size), rl.Red)

}

func (e *EntityChar) Collision(target EntityChar) {}
