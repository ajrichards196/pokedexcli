package main

import(
	"fmt"
	"strings"
	"bufio"
	"os"
)

type cliCommand struct {
    name        string
    description string
    callback    func() error
}

var registry map[string]cliCommand

func init() {
    registry = map[string]cliCommand{
        "exit": {name: "exit", description: "Exit the Pokedex", callback: commandExit},
        "help": {name: "help", description: "Displays a help message", callback: commandHelp},
    }
}

func commandExit() error {
    fmt.Println("Closing the Pokedex... Goodbye!")
    os.Exit(0)
    return nil
}

func commandHelp() error {
    fmt.Println("Welcome to the Pokedex!")
    fmt.Println("Usage:")
    fmt.Println()
    for _, cmd := range registry {
        fmt.Printf("%s: %s\n", cmd.name, cmd.description)
    }
    return nil
}


func cleanInput(text string) []string {
	s := strings.TrimSpace(text)
	s = strings.ToLower(s)
	return strings.Fields(s)
}


func main() {

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scan := scanner.Scan() 
		if !scan {
			fmt.Print(scanner.Err())
			break
		}	
		parts := cleanInput(scanner.Text())
		if len(parts) == 0 {
			continue
		}

		command, commandExists := registry[parts[0]]
		if !commandExists {
			fmt.Print("Unknown command\n")
		}
		command.callback()
		} 
}