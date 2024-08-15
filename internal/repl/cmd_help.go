package repl

import (
	"errors"
	"fmt"
)

func execHelpCmd(args CmdArgs) error {
	cmds := GetCmds()

	fmt.Println("")
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")

	if len(cmds) == 0 {
		return errors.New(
			"could not display help message. No commands found",
		)
	}
	for _, v := range cmds {
		fmt.Printf("%s: %s\n", v.Name, v.Description)
	}

	return nil
}
