package config

import "context"

type Configuration interface {
	Init(ctx context.Context) error
	Stop(ctx context.Context) error
}

var data []Configuration

func init() {
	data = make([]Configuration, 0)
	data = append(data, &Banner{})
	data = append(data, &Viper{})
	data = append(data, &Server{})
}

func GetConfigs() []Configuration {
	return data
}
