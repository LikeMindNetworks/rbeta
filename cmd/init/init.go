package init

import (
	"gopkg.in/alecthomas/kingpin.v2"
)

func Command(rbeta *kingpin.Application) *kingpin.CmdClause {
	initCmd := rbeta.Command("init", "init rbeta")

	initCmd.
		Flag(
			"namespace", "Name space for the rbeta deployment.",
		).Short('n').Required().String()

	return initCmd;
}
