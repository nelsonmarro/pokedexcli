package repl

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func StartRepl(args CmdArgs) {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		cmdName := words[0]
		if len(words) > 1 {
			args.Arguments = words[1:]
		}

		cmdAction, err := GetCmdAction(cmdName)
		if err != nil {
			fmt.Println(err)
			continue
		}

		err = cmdAction(args)
		if err != nil {
			fmt.Printf("error executing %s command!\n", cmdName)
			fmt.Println(err)
			continue
		}

		args.Arguments = []string{}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}
