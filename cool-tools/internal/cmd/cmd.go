package cmd

import (
	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Main = gcmd.Command{
		Name:        "cool-tools",
		Usage:       "cool-tools [command] [args...]",
		Brief:       "cool-tools is a collection of tools for cool people.",
		Description: `cool-tools is a collection of tools for cool people.`,
	}
)
