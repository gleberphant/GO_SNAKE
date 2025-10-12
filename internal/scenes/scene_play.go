package scenes

import (
	"fmt"
	"go_snake/internal/entities"
	"go_snake/internal/tiles"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Play struct {
	Base
	player entities.Entity
}

func (p *Play) Load() error {
	//configurar variaveis
	p.Base.Load()
	p.sceneType = PLAY
	p.nextSceneType = OVER

	// configurar funções
	p.updateEntitiesFunc = p.updateEntities
	p.eventHandlerFunc = p.eventHandler
	p.drawSceneFunc = p.drawScene

	// configurar entidades
	p.player.Load()
	p.player.TileX = 400 / (tiles.TSIZE)
	p.player.TileY = 300 / (tiles.TSIZE)

	return nil
}

// gerenciar comandos
func (p *Play) eventHandler() error {
	p.Base.eventHandler()

	// leitura de controles do jogador
	p.player.ChangeDirection(0)

	if rl.IsKeyPressed(rl.KeyA) {
		fmt.Println("pressionado A")
		p.player.ChangeDirection(45)
	}
	if rl.IsKeyPressed(rl.KeyD) {
		fmt.Println("pressionado D")
		p.player.ChangeDirection(-45)
	}

	return nil
}

// atualizar entidades
func (p *Play) updateEntities() error {

	p.Base.updateEntities()

	// atualizar personagem
	p.player.UpdatePosition()

	// atualizar inimigos

	return nil
}

// desenhar cena
func (p *Play) drawScene() error {

	rl.BeginDrawing()

	//desenhar background
	rl.ClearBackground(rl.DarkGreen)

	//desenhar maptile

	//desenhar entidades

	//desenhar personagens
	p.player.Draw()

	//desenhar hud
	rl.DrawText("CENAR PLAY:", 20, 100, 60, rl.White)
	rl.EndDrawing()

	return nil
}
