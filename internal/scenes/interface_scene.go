package scenes

type Scene interface {
	//public methods
	Load() error
	Loop() error
	Unload() error
	NextSceneType() SceneEnum
	GetSceneType() SceneEnum

	//private methods

	eventHandler() error
	updateEntities() error
	drawScene() error
}
