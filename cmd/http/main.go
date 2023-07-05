package main

import (
	"context"
	"fmt"
	"github.com/felipe-brito/cobranca-bb/internal/config"
	"os"
	"os/signal"
	"syscall"
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8081
// @BasePath  /api/v1

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	ctx := context.Background()

	sign := make(chan os.Signal, 1)
	signal.Notify(sign, syscall.SIGINT, syscall.SIGTERM)

	for _, c := range config.GetConfigs() {
		err := c.Init(ctx)
		if err != nil {
			fmt.Println("Initialization error")
		}
	}

	sig := <-sign
	fmt.Println()
	fmt.Println(fmt.Sprintf("Sinal %s: terminated...", sig))

	for _, c := range config.GetConfigs() {
		err := c.Stop(ctx)
		if err != nil {
			fmt.Println(err)
		}
	}
}
