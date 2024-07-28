package internal

type CommandHandler func(args ...string) error

type Command struct {
	Name    string
	Desc    string
	Handler CommandHandler
}

func GetPokedexCommands(p *Pokedex) map[string]Command {
	return map[string]Command{
		"help": {Name: "help", Desc: "Prints the help message", Handler: func(args ...string) error {
			return p.help(args...)
		}},
		"exit": {Name: "help", Desc: "Exits Pokedex", Handler: func(args ...string) error {
			return p.exit(args...)
		}},
		"map": {Name: "map", Desc: "Next location areas", Handler: func(args ...string) error {
			return p.mapForward(args...)
		}},
		"mapb": {Name: "mapb", Desc: "Previous location areas", Handler: func(args ...string) error {
			return p.mapBack(args...)
		}},
		"explore": {Name: "explore", Desc: "Explore an area. Usage 'explore [area-name]'", Handler: func(args ...string) error {
			return p.explore(args...)
		}},
	}
}
