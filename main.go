package main

import (
	"net/http"

	"github.com/klnusbaum/chillserve/config"
	"github.com/klnusbaum/chillserve/handlers"
	"io/ioutil"
	"os/user"
	"flag"
)

func main() {
	parsedArgs := parseArgs()
	config, err := config.GetConfig(parsedArgs, user.Current, ioutil.ReadFile)
	if err != nil {
		panic("Could not read config file")
	}
	setupHandlers(config)
	http.ListenAndServe(":8080", nil)
}

func parseArgs() config.ParsedArgs {
	parsedArgs := config.ParsedArgs{}
	flag.StringVar(&parsedArgs.ConfigFile, "config", "", "File containing configuration for chillserve")
	flag.Parse()
	return parsedArgs
}

func setupHandlers(config *config.Config) {
	ch := handlers.NewChillifierHandler(config.Replacements)
	rch := handlers.NewRandomChillHandler(config.Phrases...)
	sc := handlers.NewStateChillHandler(config.StateImagesLocation)

	http.Handle("/chillify", ch)
	http.Handle("/chill", rch)
	http.Handle("/states_chill", sc)
}
