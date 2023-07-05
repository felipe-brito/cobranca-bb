package config

import (
	"context"
	resource "github.com/felipe-brito/cobranca-bb"
	"github.com/felipe-brito/cobranca-bb/internal/utils"
	"html/template"
	"os"
	"runtime"
)

type Banner struct {
	Black   string
	Blue    string
	Yellow  string
	Red     string
	Green   string
	White   string
	Default string
	Info    RuntimeInfo
}

type RuntimeInfo struct {
	GoVersion string
	GoOS      string
	GoARCH    string
	NumCPU    int
}

func (b *Banner) Init(_ context.Context) error {

	b.setDefaultBanner()

	content, err := resource.Resource.ReadFile(utils.BannerFileResourcePath)
	if err != nil {
		return err
	}

	tmp, err := template.New(utils.BannerTemplateName).Parse(string(content))
	if err != nil {
		return err
	}

	err = tmp.Execute(os.Stderr, b)
	if err != nil {
		return err
	}

	return nil
}

func (b *Banner) Stop(_ context.Context) error {
	return nil
}

func (b *Banner) setDefaultBanner() {
	b.Black = utils.Black
	b.Blue = utils.Blue
	b.Yellow = utils.Yellow
	b.Red = utils.Red
	b.Green = utils.Green
	b.White = utils.White
	b.Default = utils.Reset
	b.Info = RuntimeInfo{
		GoVersion: utils.GoVersion(),
		GoOS:      runtime.GOOS,
		GoARCH:    runtime.GOARCH,
		NumCPU:    runtime.NumCPU(),
	}
}
