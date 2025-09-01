package main

import(
	"fmt"
	"strings"
	"bufio"
	"os"
	"net/http"
	"io"
	"encoding/json"
)

type cliCommand struct {
    name        string
    description string
    callback    func(config *Config) error
}

type LocResponse struct {
	Count 		int 	`json:"count"`
	Next		string	`json:"next"`
	Previous 	string	`json:"previous"`
	Results 	[]Location	`json:"results"`
}
type Location struct {
	Name	string	`json:"name"`
	URL 	string 	`json:"url"`
}
type Config struct {
	Next		string	`json:"next"`
	Previous 	string	`json:"previous"`
}
var registry map[string]cliCommand


func init() {
    registry = map[string]cliCommand{
        "exit": {name: "exit", description: "Exit the Pokedex", callback: commandExit},
        "help": {name: "help", description: "Displays a help message", callback: commandHelp},
        "map": {name: "map", description: "Show next 20 locations", callback:commandMap},
        "mapb": {name: "mapb", description: "Show previous 20 locations", callback:commandMapb},
    }
}

func commandMap(config *Config) error {
	if config.Next == "" {
		config.Next = "https://pokeapi.co/api/v2/location-area" 
	}
	url := config.Next
	res, err := http.Get(url)
	if err != nil {
		return err
	}  
	defer res.Body.Close()
	var resources LocResponse
	data, _ := io.ReadAll(res.Body)
	err = json.Unmarshal(data, &resources)
	if err != nil {
		return err
	}

	for _, result := range resources.Results{
		fmt.Println(result.Name)
	}
	config.Next = resources.Next
	config.Previous = resources.Previous
	return nil
}
func commandMapb(config *Config) error {
	
	if config.Previous == "" {
		fmt.Println("you're on the first page")
		return nil
	}
	url := config.Previous
	res, err := http.Get(url)
	if err != nil {
		return err
	}  
	defer res.Body.Close()
	var resources LocResponse
	data, _ := io.ReadAll(res.Body)
	err = json.Unmarshal(data, &resources)
	if err != nil {
		return err
	}

	for _, result := range resources.Results{
		fmt.Println(result.Name)
	}
	config.Previous = resources.Previous
	config.Next = resources.Next
	return nil
}
func commandExit(config *Config) error {
    fmt.Println("Closing the Pokedex... Goodbye!")
    os.Exit(0)
    return nil
}

func commandHelp(config *Config) error {
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
	locURLs := Config{}
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
		err := command.callback(&locURLs)
		if err != nil {
			fmt.Println(err)
			continue
		}
	} 
}