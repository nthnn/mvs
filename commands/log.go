package commands

import (
	"fmt"
	"path/filepath"

	"github.com/nthnn/mvs/core"
	"github.com/nthnn/mvs/logger"
	"github.com/nthnn/mvs/utils"
	"github.com/vmihailenco/msgpack/v5"
)

func stylizeText(name string) string {
	startingColor := [3]int{0, 123, 255}
	endColor := [3]int{111, 66, 193}

	return logger.Colorize(
		name,
		startingColor,
		endColor,
		false, false,
	)
}

func LogCommand() {
	head := utils.ResolveHead()

	for head != "" {
		path := filepath.Join(core.CommitsPath, head)

		if err := utils.Verify(path); err != nil {
			logger.Error("Tamper detected at commit: %s", head)
			return
		}
		raw, _ := utils.ReadCompressed(path)

		var commit core.Commit
		msgpack.Unmarshal(raw, &commit)

		fmt.Printf(
			"\x1b[1mCommit\x1b[0m:\t%s\r\n"+
				"\x1b[1mAuthor\x1b[0m: %s\r\n"+
				"\x1b[1mDate\x1b[0m:\t%s"+
				"\r\n\r\n\t%s\r\n\r\n",
			stylizeText(commit.Hash),
			stylizeText(commit.Author+" <"+commit.Email+">"),
			stylizeText(commit.Timestamp),
			stylizeText(commit.Message),
		)
		head = commit.Parent
	}
}
