// main APP

package main

import (
	"fmt"
	"go_snake/internal/app"
	"log"
)

func main() {

	var appInstance app.App
	var err error

	fmt.Println(" ---- INICIALIZANDO APLICAÇÃO  ")

	if err = appInstance.Run(); err != nil {
		log.Fatal("Erro na execução da aplicação: ", err)
	}

	fmt.Println("---- APLICAÇÃO FINALIZADA")

}
