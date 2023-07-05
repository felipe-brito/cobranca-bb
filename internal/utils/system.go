package utils

import (
	"os/exec"
	"strings"
)

func GoVersion() string {
	cmd, err := exec.Command("go", "version").Output()
	if err != nil {
		return ""
	}
	return strings.Split(string(cmd), " ")[2]
}
