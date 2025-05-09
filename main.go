package main

import (
	"os"
	"strings"

	"github.com/nthnn/mvs/commands"
	"github.com/nthnn/mvs/core"
	"github.com/nthnn/mvs/logger"
)

var ActualVersion string

func usage() {
	startingColor := [3]int{111, 66, 193}
	endColor := [3]int{0, 123, 255}

	verLength := len(ActualVersion) - 1
	if verLength < 30 {
		ActualVersion = ActualVersion +
			strings.Repeat(" ", 9)
	}

	var (
		title = logger.Colorize(
			"Minimal Versioning System     ",
			startingColor, endColor,
			true, false,
		)

		buildVersion = logger.Colorize(
			ActualVersion,
			startingColor, endColor,
			false, false,
		)

		url = logger.Colorize(
			"(https://github.com/nthnn/mvs)",
			startingColor, endColor,
			false, true,
		)
	)

	core.PrintBanner(
		title, buildVersion,
		url,
	)
}

func main() {
	if len(os.Args) < 2 {
		usage()
		return
	}

	core.LoadGlobalConfig()
	switch os.Args[1] {
	case "init":
		commands.InitializeCommand()

	case "add":
		commands.AddCommand(os.Args[2:])

	case "remove":
		commands.RemoveCommand(os.Args[2:])

	case "commit":
		commands.CommitCommand()

	case "log":
		commands.LogCommand()

	case "branch":
		commands.BranchCommand(os.Args[2:])

	case "checkout":
		commands.CheckoutCommand(os.Args[2:])

	case "status":
		commands.StatusCommand()

	case "tree":
		commands.TreeCommand()

	case "amend":
		commands.AmendCommand()

	default:
		usage()
	}
}
