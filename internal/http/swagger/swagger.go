package swagger

import (
	"fmt"
	"github.com/felipe-brito/cobranca-bb/internal/global"
	"github.com/felipe-brito/cobranca-bb/internal/http/swagger/docs"
	httpSwagger "github.com/swaggo/http-swagger"
	"github.com/swaggo/swag"
	"net/http"
)

type Swagger struct {
}

func (s Swagger) Config() http.HandlerFunc {

	docs.SwaggerInfo.Title = global.AppConfig.App.Title
	docs.SwaggerInfo.Version = global.AppConfig.App.Version
	docs.SwaggerInfo.Description = global.AppConfig.App.Description
	docs.SwaggerInfo.BasePath = global.AppConfig.Swagger.Base
	docs.SwaggerInfo.Host = fmt.Sprintf("%s:%d", global.AppConfig.Swagger.Host, global.AppConfig.Swagger.Port)
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	return httpSwagger.Handler(httpSwagger.URL("/swagger.json"))
}

func (s Swagger) Read() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusOK)
		s, _ := swag.ReadDoc()
		_, _ = writer.Write([]byte(s))
	}
}
