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
	player   entities.EntityChar
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

	for i := 1; i < 4; i++ {
		item := entities.EntityItem{}
		item.Load()
		item.PosY = 3 * (tiles.TSIZE) * int32(i)
		item.PosX = 3 * (tiles.TSIZE)

		fmt.Println("adicionando item: pos %d %d", item.PosY, item.PosX)

		p.itemPool.PushFront(&item)
	}
	return nil
}

// gerenciar comandos
func (p *Play) eventHandler() error {
	p.ScenePrimitive.eventHandler()

	// leitura de controles do jogador

	if rl.IsKeyPressed(rl.KeyA) {
		fmt.Println("pressionado A")
		p.player.Control(entities.ROTATE_LEFT)
	}
	if rl.IsKeyPressed(rl.KeyD) {
		fmt.Println("pressionado D")
		p.player.Control(entities.ROTATE_RIGHT)
	}

	p.player.Control(entities.MOVE)

	return nil
}

// atualizar entidades
func (p *Play) updateEntities() error {

	p.ScenePrimitive.updateEntities()

	// atualizar personagem
	p.player.Update()

	// atualizar itens

	for node := p.itemPool.Front(); node != nil; node = node.Next() {

		item, ok := node.Value.(*entities.EntityItem)

		if !ok {
			continue
		}

		if item.IsAlive {

			// Checar Colisão
			if rl.CheckCollisionCircles(
				p.player.GetPositionVector(),
				float32(p.player.Size),
				item.GetPositionVector(),
				float32(item.Size)) {

				item.Collision(p.player)
			}
			item.Update()
		} else {
			currentNode := node.Next()
			p.itemPool.Remove(node)
			node = currentNode
		}

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

		item, ok := node.Value.(*entities.EntityItem)

		if !ok {
			fmt.Println("Erro do item morto")
			continue
		}

		item.Draw()

	}

	//desenhar personagens
	p.player.Draw()

	//desenhar hud
	rl.DrawText("CENAR PLAY:", 20, 100, 60, rl.White)
	rl.EndDrawing()

	return nil
}
