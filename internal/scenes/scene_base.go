package scenes

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Base struct {
	sceneType     SceneEnum
	nextSceneType SceneEnum

	eventHandlerFunc   func() error
	updateEntitiesFunc func() error
	drawSceneFunc      func() error

	isRunning bool
}

func (b *Base) Load() error {
	//configurar variaveis
	b.sceneType = INTRO
	b.nextSceneType = MENU
	b.isRunning = true

	// configurar funções
	b.drawSceneFunc = b.drawScene
	b.updateEntitiesFunc = b.updateEntities
	b.eventHandlerFunc = b.eventHandler

	return nil
}

// LOOP DA CENA
func (b *Base) Loop() error {

	fmt.Println(".. Iniciando loop da cena")

	// LOOP DA CENA
	for b.isRunning {

		// event handler
		b.eventHandlerFunc()

		// update Entities
		b.updateEntitiesFunc()

		//b.drawScene()
		b.drawSceneFunc()

	}

	return nil

}

func (b *Base) eventHandler() error {
	// handle events
	if rl.WindowShouldClose() {
		b.nextSceneType = QUIT
		b.isRunning = false
		return nil

	}

	if rl.IsKeyPressed(rl.KeyEnter) {
		b.isRunning = false
		return nil
	}
	return nil
}

func (b *Base) updateEntities() error {
	return nil
}

func (b *Base) drawScene() error {
	return nil
}

func (b *Base) Unload() error {
	return nil
}

func (b *Base) NextSceneType() SceneEnum {
	return b.nextSceneType
}

func (b *Base) GetSceneType() SceneEnum {
	return b.sceneType
}
