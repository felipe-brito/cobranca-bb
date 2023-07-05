package utils

// Status
const (
	UP   = "UP"
	DOWN = "DOWN"
)

// Color
const (
	Black  = "\x1b[0m\x1b[30m"
	Blue   = "\x1b[0m\x1b[34m"
	Yellow = "\x1b[0m\x1b[33m"
	Red    = "\x1b[0m\x1b[31m"
	Green  = "\x1b[0m\x1b[32m"
	White  = "\x1b[0m\x1b[37m"
	Reset  = "\033[0m"
)

// Banner
const (
	BannerFileResourcePath = "resource/banner.txt"
	BannerTemplateName     = "banner"
)

// Viper
const (
	ApplicationFileResourcePath = "resource/application.yaml"
	ApplicationFileName         = "application"
	Extension                   = "yaml"
)

// KB Unit
const (
	KB = uint64(1024)
)
