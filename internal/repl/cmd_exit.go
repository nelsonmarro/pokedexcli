package repl

import (
	"os"
)

func execExitCmd(args CmdArgs) error {
	os.Exit(0)
	return nil
}
