package config

import (
	"bytes"
	"context"
	resource "github.com/felipe-brito/cobranca-bb"
	"github.com/felipe-brito/cobranca-bb/internal/global"
	"github.com/felipe-brito/cobranca-bb/internal/utils"
	"github.com/spf13/viper"
	"strings"
)

type Viper struct {
}

func (v Viper) Init(_ context.Context) error {

	app := global.Application{}

	content, err := resource.Resource.ReadFile(utils.ApplicationFileResourcePath)
	if err != nil {
		return err
	}

	setup := viper.GetViper()

	setup.SetConfigType(utils.Extension)
	setup.SetConfigName(utils.ApplicationFileName)
	setup.AllowEmptyEnv(true)
	setup.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	setup.AutomaticEnv()

	err = viper.ReadConfig(bytes.NewBuffer(content))
	if err != nil {
		return err
	}

	err = setup.Unmarshal(&app)
	if err != nil {
		return err
	}

	global.AppConfig = app

	return nil
}

func (v Viper) Stop(_ context.Context) error {
	return nil
}
