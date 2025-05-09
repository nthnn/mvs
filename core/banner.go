/*
 * Copyright 2025 Nathanne Isip
 * This file is part of Minimal Versioning System (https://github.com/nthnn/mvs)
 * This code is licensed under MIT license (see LICENSE for details)
 */

package core

import (
	"fmt"

	"github.com/nthnn/mvs/logger"
)

const asciiArt = "\r\n" +
	"\x1b[1;36m███\x1b[1;30m╗\x1b[1;36m   \x1b[1;36m███\x1b[1;30m╗\x1b[1;36m██\x1b[1;30m╗   \x1b[1;36m██\x1b[1;30m╗\x1b[1;36m███████\x1b[1;30m╗\x1b[0m\r\n" +
	"\x1b[1;36m████\x1b[1;30m╗ \x1b[1;36m████\x1b[1;30m║\x1b[1;36m██\x1b[1;30m║   \x1b[1;36m██\x1b[1;30m║\x1b[1;36m██\x1b[1;30m╔════╝\x1b[0m    %s\r\n" +
	"\x1b[1;36m██\x1b[1;30m╔\x1b[1;36m████\x1b[1;30m╔\x1b[1;36m██\x1b[1;30m║\x1b[1;36m██\x1b[1;30m║   \x1b[1;36m██\x1b[1;30m║\x1b[1;36m███████\x1b[1;30m╗\x1b[0m    %s\r\n" +
	"\x1b[1;36m██\x1b[1;30m║╚\x1b[1;36m██\x1b[1;30m╔╝\x1b[1;36m██\x1b[1;30m║╚\x1b[1;36m██\x1b[1;30m╗ \x1b[1;36m██\x1b[1;30m╔╝╚════\x1b[1;36m██\x1b[1;30m║\x1b[0m    %s\r\n" +
	"\x1b[1;36m██\x1b[1;30m║ ╚═╝ \x1b[1;36m██\x1b[1;30m║ ╚\x1b[1;36m████\x1b[1;30m╔╝ \x1b[1;36m███████\x1b[1;30m║\x1b[0m"

func stylizeCommandName(name string) string {
	startingColor := [3]int{111, 66, 193}
	endColor := [3]int{0, 123, 255}

	return logger.Colorize(
		name,
		startingColor,
		endColor,
		true, false,
	)
}

func PrintBanner(title, buildVersion, url string) {
	banner := fmt.Sprintf(asciiArt, title, buildVersion, url) +
		"\r\n\r\n" +
		"\x1b[1mUsage\x1b[0m:\r\n" +
		"  " + stylizeCommandName("mvs      ") + "<command>  [options]\r\n\r\n" +
		"\x1b[1mCommands\x1b[0m:\r\n" +
		"  " + stylizeCommandName("init     ") + "           Initialize a new repository\r\n" +
		"  " + stylizeCommandName("branch   ") + "[name]     List or create branch\r\n" +
		"  " + stylizeCommandName("checkout ") + "<name>     Switch branch or commit\r\n" +
		"  " + stylizeCommandName("add      ") + "<paths>    Stage file changes\r\n" +
		"  " + stylizeCommandName("remove   ") + "<paths>    Unstage file changes\r\n" +
		"  " + stylizeCommandName("commit   ") + "<message>  Commit staged changes\r\n" +
		"  " + stylizeCommandName("amend    ") + "<message>  Amend the message of previous commit\r\n" +
		"  " + stylizeCommandName("log      ") + "           Show commit history\r\n" +
		"  " + stylizeCommandName("status   ") + "           Show staged/modified/untracked files\r\n" +
		"  " + stylizeCommandName("tree     ") + "           Render an ASCII branch tree\r\n\r\n"

	fmt.Print(banner)
}
