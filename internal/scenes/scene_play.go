package scenes

import (
	"container/list"
	"fmt"
	"go_snake/internal/entities"
	"go_snake/internal/tiles"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Play struct {
	ScenePrimitive
	player   entities.EntitySnake
	itemPool *list.List
}

func (p *Play) Load() error {
	fmt.Println("Carregando Cena Play")

	//configurar variaveis
	p.ScenePrimitive.Load()
	p.sceneType = PLAY
	p.nextSceneType = OVER

	// configurar funções
	p.updateEntitiesFunc = p.updateEntities
	p.eventHandlerFunc = p.eventHandler
	p.drawSceneFunc = p.drawScene

	// configurar jogador
	p.player.Load()
	p.player.PosX = 2 * (tiles.TSIZE)
	p.player.PosY = 4 * (tiles.TSIZE)

	// configurar itens
	p.itemPool = list.New()

	for i := 1; i < 3; i++ {
		p.generateNewItem()
	}
	return nil
}

func (p *Play) generateNewItem() {
	newItem := entities.EntityFood{}
	newItem.Load()
	newItem.PosX = rl.GetRandomValue(50, 700)
	newItem.PosY = rl.GetRandomValue(50, 500)

	fmt.Printf("adicionando item: pos %d %d\n", newItem.PosY, newItem.PosX)

	p.itemPool.PushFront(&newItem)
}

// gerenciar comandos
func (p *Play) eventHandler() error {
	p.ScenePrimitive.eventHandler()

	// leitura de controles do jogador

	if rl.IsKeyPressed(rl.KeyA) {
		//fmt.Println("pressionado A")
		p.player.Control(entities.ROTATE_LEFT)
	}
	if rl.IsKeyPressed(rl.KeyD) {
		//fmt.Println("pressionado D")
		p.player.Control(entities.ROTATE_RIGHT)
	}

	p.player.Control(entities.MOVE)

	return nil
}

// atualizar entidades
func (p *Play) updateEntities() error {

	p.ScenePrimitive.updateEntities()

	// atualizar personagem
	if p.player.IsAlive {
		p.player.Update()
	} else {
		p.isRunning = false
		return nil
	}

	// atualizar itens

	for node := p.itemPool.Front(); node != nil; {

		// verifica error
		item, ok := node.Value.(*entities.EntityFood)

		if !ok {
			continue
		}

		item.Update()

		// Checar Colisão
		if rl.CheckCollisionCircles(
			p.player.GetPositionVector(),
			float32(p.player.Size),
			item.GetPositionVector(),
			float32(item.Size)) {

			item.Collision(&p.player)
			p.generateNewItem()
		}

		if !item.IsAlive {
			nextNode := node.Next()
			p.itemPool.Remove(node)
			node = nextNode
			continue
		}

		node = node.Next()

	}

	return nil
}

// desenhar cena
func (p *Play) drawScene() error {

	rl.BeginDrawing()

	//desenhar background
	rl.ClearBackground(rl.DarkGreen)

	//desenhar maptile

	//desenhar itens
	for node := p.itemPool.Front(); node != nil; node = node.Next() {

		item, ok := node.Value.(*entities.EntityFood)

		if !ok {
			fmt.Println("Erro do item morto")
			continue
		}

		item.Draw()

	}

	//desenhar personagens
	p.player.Draw()

	//desenhar hud
	//rl.DrawText("CENAR PLAY:", 20, 100, 60, rl.White)
	rl.EndDrawing()

	return nil
}
