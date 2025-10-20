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
	CharType  CharEnum
	life      int
	bodyList  *list.List
	SnakeHead *SnakeBody
}

func (e *EntitySnake) Load() {
	e.EntityPrimitive.Load()
	e.Color = rl.Green
	e.life = 1
	e.bodyList = list.New()
	e.Speed = tiles.TSIZE

	fmt.Printf("Adicionando cabeça X: %d Y %d\n", e.PosX, e.PosY)

	e.bodyList.PushFront(&SnakeBody{X: e.PosX, Y: e.PosY, Speed: e.Speed, Direction: e.Direction})

	e.SnakeHead, _ = e.bodyList.Front().Value.(*SnakeBody)

}

func (e *EntitySnake) Update() {
	e.Move()

}

func (e *EntitySnake) Move() {

	// movimento personagem
	e.EntityPrimitive.Move()

	// movimenta a cabeça
	nodeHead := e.bodyList.Front()

	head, ok := nodeHead.Value.(*SnakeBody)

	if !ok {
		return
	}

	head.X = e.PosX
	head.Y = e.PosY
	head.Direction = e.Direction
	head.Speed = e.Speed

	// movimenta o corpo
	nodeBody := nodeHead.Next()

	for nodeBody != nil {

		// define o corpo atual
		current, ok := nodeBody.Value.(*SnakeBody)
		if !ok {
			fmt.Printf("Error: %v\n", ok)
			break
		}
		// define alvo - corpo a frente
		target, ok := nodeHead.Value.(*SnakeBody)
		if !ok {
			fmt.Printf("Error: %v\n", ok)
			break
		}
		// Define direção e velocidade do corpo
		current.Direction = (math.Atan2(float64(target.X-current.X), float64(target.Y-current.Y)))
		current.Speed = e.Speed

		//calcula distancia do corpo da frente
		distance := math.Sqrt(math.Pow(float64(target.X-current.X), 2) + math.Pow(float64(target.Y-current.Y), 2))

		//movimento apenas se não estiver muito proximo
		if distance > (tiles.TSIZE + 10) {

			current.X += int32(current.Speed * math.Sin(current.Direction))
			current.Y += int32(current.Speed * math.Cos(current.Direction))

		}

		// checka colisão da cabeça com  o corpo
		if rl.CheckCollisionCircles(
			rl.Vector2{X: float32(head.X), Y: float32(head.Y)}, float32(e.Size-2),
			rl.Vector2{X: float32(current.X), Y: float32(current.Y)}, float32(e.Size-2),
		) {
			fmt.Println("Cabeça mordeu o corpo")
			e.IsAlive = false
		}

		// proximo corpo
		nodeHead = nodeBody
		nodeBody = nodeBody.Next()
	}

}

func (e *EntitySnake) Collision(target EntityPrimitive) {}

func (e *EntitySnake) EatFood() {

	tail := e.bodyList.Back().Value.(*SnakeBody)

	e.bodyList.PushBack(&SnakeBody{
		X:         tail.X - int32((float64(tiles.TSIZE+10))*math.Sin(tail.Direction)),
		Y:         tail.Y - int32((float64(tiles.TSIZE+10))*math.Cos(tail.Direction)),
		Direction: e.Direction,
		Speed:     e.Speed,
	})

	e.life = e.bodyList.Len()

	fmt.Printf("Eat food  X: %d Y %d\n", e.PosX, e.PosY)

}

func (e *EntitySnake) Draw() {

	prev := &SnakeBody{
		X: e.PosX + int32(32*math.Sin(e.Direction)),
		Y: e.PosY + int32(32*math.Cos(e.Direction)),
	}

	for node := e.bodyList.Front(); node != nil; node = node.Next() {

		body, ok := node.Value.(*SnakeBody)

		if !ok {
			continue
		}

		rl.DrawCircleLines(body.X, body.Y, float32(e.Size), rl.Blue)

		rl.DrawText(fmt.Sprintf("Angulo %f", e.Direction*rl.Rad2deg), 20, 20, 14, rl.White)

		rl.DrawLine(body.X, body.Y, prev.X, prev.Y, rl.Red)

		prev = body

	}
}
