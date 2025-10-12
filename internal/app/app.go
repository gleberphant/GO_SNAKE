package app

import (
	"fmt"
	"go_snake/internal/scenes"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	SCREENWIDTH  = 800
	SCREENHEIGHT = 600
)

// dados para configuração da aplicação
type App struct {
	Running      bool
	ScreenHeight int
	ScreenWidth  int

	currentScene     scenes.Scene
	currentSceneType scenes.SceneEnum
}

// métodos públicos
// EXECUTAR APLICAÇÃO
func (a *App) Run() error {

	// carregar aplicação
	if err := a.load(); err != nil {
		return fmt.Errorf("falha ao carregar a aplicação: %w", err)
	}

	// executar loop da aplicação
	if err := a.loop(); err != nil {
		return fmt.Errorf("erro durante a execução do loop: %w", err)
	}

	// encerra aplicação
	if err := a.unload(); err != nil {
		return fmt.Errorf("erro no encerramento/limpeza da aplicação: %w", err)
	}

	return nil
}

// métodos privados
// CONFIGURA A APLICAÇÃO
func (a *App) load() error {
	fmt.Println(" Configuração da aplicação")

	//configura parametros

	a.ScreenWidth = SCREENWIDTH
	a.ScreenHeight = SCREENHEIGHT
	a.Running = true

	//inicializa raylib
	rl.InitWindow(int32(a.ScreenWidth), int32(a.ScreenHeight), "GO SNAKE RAY LIB")
	rl.SetTraceLogLevel(rl.LogDebug)
	rl.SetTargetFPS(30)

	//configura cena inicial
	a.currentSceneType = scenes.PLAY
	a.currentScene = nil
	return nil

}

// SELECIONA A CENA
func (a *App) selectScene() scenes.Scene {

	fmt.Print("---- Cena ")
	switch a.currentSceneType {

	case scenes.INTRO:
		fmt.Println("... INTRO ")
		return new(scenes.Base)

	case scenes.MENU:
		fmt.Println("... MENU ")
		return new(scenes.Menu)

	case scenes.PLAY:
		fmt.Println("... PLAY GAME ")
		return new(scenes.Play)

	case scenes.OVER:
		fmt.Println("... OVER ")
		return new(scenes.Over)

	case scenes.QUIT:
		fmt.Println("... QUIT ")
		return nil
	}

	fmt.Println("... ERROR  ")
	return nil

}

// LOOP DA APLICAÇÃO
func (a *App) loop() error {

	// loop da aplicação
	fmt.Println(" Execução loop da aplicação")
	for a.Running {

		// seleciona cena
		fmt.Println("-- Selecionar Cena ")

		a.currentScene = nil

		if a.currentScene = a.selectScene(); a.currentScene == nil {
			a.Running = false
			break
		}

		// carrega cena
		fmt.Println("-- Carrega Cena ", a.currentSceneType)
		a.currentScene.Load()

		// executa cena
		fmt.Println("-- Executa Cena ", a.currentSceneType)
		a.currentScene.Loop()

		// informa proxima cena
		fmt.Println("-- Próxima Cena ", a.currentScene.NextSceneType())
		a.currentSceneType = a.currentScene.NextSceneType()

		// finaliza cena
		fmt.Println("-- Finaliza Cena  ")
		a.currentScene.Unload()
		a.currentScene = nil
	}

	return nil
}

func (a *App) unload() error {
	fmt.Println(" Encerramento/limpeza da aplicação")
	a.Running = false
	rl.CloseWindow()
	return nil
}
