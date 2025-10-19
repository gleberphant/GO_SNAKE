package entities

import (
	"container/list"
	"fmt"
	"go_snake/internal/tiles"
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type SnakeBody struct {
	X         int32
	Y         int32
	Direction float64
	Speed     float64
}

type EntitySnake struct {
	EntityPrimitive
	CharType         CharEnum
	life             int
	bodyList         *list.List
	headNode         *list.Element
	oldPosX, oldPosY int32
}

func (e *EntitySnake) Load() {
	e.EntityPrimitive.Load()
	e.Color = rl.Green
	e.life = 1
	e.bodyList = list.New()
	e.Speed = tiles.TSIZE

	fmt.Printf("Adicionado cabeça X: %d Y %d\n", e.PosX, e.PosY)

	e.bodyList.PushFront(&SnakeBody{X: e.PosX, Y: e.PosY, Speed: e.Speed, Direction: e.Direction})
	e.headNode = e.bodyList.Front()

}

func (e *EntitySnake) Update() {

	// atualiza cabeça
	e.oldPosX = e.PosX
	e.oldPosY = e.PosY

	e.PosX += int32(e.Speed * math.Sin(e.Direction))
	e.PosY += int32(e.Speed * math.Cos(e.Direction))

	nodeHead := e.bodyList.Front()

	head, ok := nodeHead.Value.(*SnakeBody)

	if ok {
		head.X = e.PosX
		head.Y = e.PosY
		head.Direction = e.Direction
		head.Speed = e.Speed
	}

	// atualiza o corpo
	nodeBody := nodeHead.Next()

	for nodeBody != nil {

		current, ok := nodeBody.Value.(*SnakeBody)
		if !ok {
			break
		}

		// checka colisão com cabeça
		if rl.CheckCollisionCircles(
			rl.Vector2{X: float32(head.X), Y: float32(head.Y)}, float32(e.Size-2),
			rl.Vector2{X: float32(current.X), Y: float32(current.Y)}, float32(e.Size-2),
		) {
			fmt.Println("MORREU")
			e.IsAlive = false
		}

		target, ok := nodeHead.Value.(*SnakeBody)
		if !ok {
			fmt.Print("ERRO")
			break
		}

		current.Direction = (math.Atan2(float64(target.X-current.X), float64(target.Y-current.Y)))

		distance := math.Sqrt(math.Pow(float64(target.X-current.X), 2) + math.Pow(float64(target.Y-current.Y), 2))
		current.Speed = e.Speed

		if distance > (tiles.TSIZE + 10) {

			current.X += int32(current.Speed * math.Sin(current.Direction))
			current.Y += int32(current.Speed * math.Cos(current.Direction))

		}
		nodeHead = nodeBody
		nodeBody = nodeBody.Next()
	}

	// Limpa Parametros das ações
	e.Speed = 0

}

func (e *EntitySnake) Draw() {

	for node := e.bodyList.Front(); node != nil; node = node.Next() {

		body, ok := node.Value.(*SnakeBody)
		if !ok {
			continue
		}

		rl.DrawLine(body.X, body.Y, body.X+int32(32*math.Sin(body.Direction)), body.Y+int32(32*math.Cos(body.Direction)), rl.Red)
		rl.DrawCircleLines(body.X, body.Y, float32(e.Size), rl.Blue)

		rl.DrawText(fmt.Sprintf("Angulo %f", e.Direction*rl.Rad2deg), 20, 20, 14, rl.White)
	}
}

func (e *EntitySnake) EatFood() {

	tail := e.bodyList.Back().Value.(*SnakeBody)

	e.bodyList.PushBack(&SnakeBody{
		X:         tail.X - int32((float64(tiles.TSIZE+10))*math.Sin(tail.Direction)),
		Y:         tail.Y - int32((float64(tiles.TSIZE+10))*math.Cos(tail.Direction)),
		Direction: e.Direction,
		Speed:     e.Speed,
	})
	fmt.Printf("corpo cresceu  X: %d Y %d\n", e.PosX, e.PosY)
	e.life = e.bodyList.Len()

}

func (e *EntitySnake) Collision(target EntityPrimitive) {}
