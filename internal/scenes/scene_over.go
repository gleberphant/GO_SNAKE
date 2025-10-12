package scenes

import rl "github.com/gen2brain/raylib-go/raylib"

type Over struct {
	ScenePrimitive
}

func (o *Over) Load() error {
	//configurar variaveis
	o.ScenePrimitive.Load()
	o.sceneType = OVER
	o.nextSceneType = MENU

	// configurar funções
	o.drawSceneFunc = o.drawScene
	o.updateEntitiesFunc = o.updateEntities
	o.eventHandlerFunc = o.eventHandler

	return nil
}

func (o *Over) drawScene() error {
	rl.BeginDrawing()

	rl.ClearBackground(rl.Brown)
	rl.DrawText("CENAR OVER:", 20, 100, 60, rl.White)
	rl.EndDrawing()
	return nil
}
