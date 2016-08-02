package config

import (
	"os/user"
	"path/filepath"
	"github.com/go-yaml/yaml"
)


// A CurrentUser is a function that can return the user running the current program.
type CurrentUser func () (*user.User, error)

// A ReadFile is a funcion that, given a fileName, will read the entire file into a byte array.
type ReadFile func(fileName string) ([]byte, error)

// A Config contains the various configuration settings available for the chill server.
type Config struct {
	Phrases             []string // A list of really chill phrases
	Replacements        map[string]string // A mapping of non-chill strings to chill strings
	StateImagesLocation string // The url location of chill state images
}

// A ParsedArgs represents arguments given to the program at run time which can be used in the calculation of
// a Config struct.
type ParsedArgs struct {
	ConfigFile string
}

const defaultConfigFile = ".chill/config.yaml"

// GetConfig calculates the Config from a given config file. The config file can be specified
// using the command line argument "-config" or can simply exist in the users home directory under
// the name ".chill/config.yaml". Config files should be in the yaml file format.
func GetConfig(parsedArgs ParsedArgs, currentUser CurrentUser, readFile ReadFile) (*Config, error) {
	configFile, err := getConfigFile(parsedArgs, currentUser)
	if err != nil {
		return nil, err
	}

	return readConfigFile(configFile, readFile)
}

func getConfigFile(parsedArgs ParsedArgs, currentUser CurrentUser) (string, error) {
	if parsedArgs.ConfigFile != "" {
		return parsedArgs.ConfigFile, nil
	}

	return getConfigFileFromHomeDir(currentUser)
}

func getConfigFileFromHomeDir(currentUser CurrentUser)  (string, error) {
	usr, err := currentUser()
	if err != nil {
		return "", err
	}

	return filepath.Join(usr.HomeDir, defaultConfigFile), nil
}

func readConfigFile(configFile string, readFile ReadFile) (*Config, error) {
	fileContents, err := readFile(configFile)
	if err != nil {
		return nil, err
	}

	config := Config{}
	err = yaml.Unmarshal(fileContents, &config)
	return &config, err
}
