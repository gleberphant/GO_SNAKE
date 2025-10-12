package scenes

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Menu struct {
	Base
}

func (m *Menu) Load() error {
	//configurar variaveis
	m.Base.Load()
	m.sceneType = MENU
	m.nextSceneType = PLAY

	// configurar funções
	m.drawSceneFunc = m.drawScene
	m.updateEntitiesFunc = m.updateEntities
	m.eventHandlerFunc = m.eventHandler

	return nil
}

func (m *Menu) drawScene() error {

	rl.BeginDrawing()

	rl.ClearBackground(rl.DarkGray)
	rl.DrawText("CENA MENU:", 20, 100, 60, rl.White)
	rl.EndDrawing()
	return nil
}
