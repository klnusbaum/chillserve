package main

import (
	"github.com/klnusbaum/chillserve/handlers"
	"net/http"

	"github.com/go-yaml/yaml"
	"io/ioutil"
	"os/user"
	"path/filepath"
)

const DefaultConfigFile = ".chill/config.yaml"

type Config struct {
	Phrases []string
	Replacements map[string]string
}

func main() {
	config, err := getConfig()
	if err != nil {
		panic("Could not read config file")
	}

	ch := handlers.NewChillifierHandler(config.Replacements)
	rch := handlers.NewRandomChillHandler(config.Phrases...)

	http.Handle("/chillify", ch)
	http.Handle("/chill", rch)
	http.ListenAndServe(":8080", nil)
}

func getConfig() (Config, error) {
	usr, err := user.Current()
	if err != nil {
		return Config{}, err
	}

	configFile := filepath.Join(usr.HomeDir, DefaultConfigFile)
	fileContents, err := ioutil.ReadFile(configFile)

	if err != nil {
		return Config{}, err
	}

	config := Config{}
	err = yaml.Unmarshal(fileContents, &config)
	return config, err
}
