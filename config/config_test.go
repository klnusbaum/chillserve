package config

import (
	"testing"
	"os/user"

	"github.com/stretchr/testify/assert"
	"fmt"
	"errors"
)

func TestGetConfig(t *testing.T) {
	config, err := GetConfig(ParsedArgs{}, testGetUser, testHomeDirConfigFileReader)
	assert.NoError(t, err, "Reading config file should have been successful.")
	assert.Len(t, config.Phrases, 2, "Phrases did not have 2 phrases in it.")
	assert.Equal(t, config.Phrases[0], "chill", "First phrase was not \"chill\".")
	assert.Equal(t, config.Phrases[1], "most chill", "First phrase was not \"most chill\".")
	assert.Len(t, config.Replacements, 1, "Replacements had more than 1 replacment pair.")
	assert.Equal(t, config.Replacements["the"], "the chill", "\"the\" was not replaced properly.")
	assert.Equal(t, config.StateImagesLocation, "http://my.imagestore.com", "Incorrect image location.")

}

func TestGetConfigWithConfigArg(t *testing.T) {
	parsedArgs := ParsedArgs{
		ConfigFile: "./expected_name.yaml",
	}
	config, err := GetConfig(parsedArgs, testGetUser, testConfigFileReader)
	assert.NoError(t, err, "Reading config file should have been successful.")
	assert.Len(t, config.Phrases, 2, "Phrases did not have 2 phrases in it.")
	assert.Equal(t, config.Phrases[0], "chill", "First phrase was not \"chill\".")
	assert.Equal(t, config.Phrases[1], "most chill", "First phrase was not \"most chill\".")
	assert.Len(t, config.Replacements, 1, "Replacements had more than 1 replacment pair.")
	assert.Equal(t, config.Replacements["the"], "the chill", "\"the\" was not replaced properly.")
	assert.Equal(t, config.StateImagesLocation, "http://my.imagestore.com", "Incorrect image location.")
}

func TestGetConfigWithUserError(t *testing.T) {
	config, err := GetConfig(ParsedArgs{}, testGetUserError, testHomeDirConfigFileReader)
	assert.NotNil(t, err, "Error was nil.")
	assert.EqualError(t, err, "Error fetching user", "Incorrect error when fetching user.")
	assert.Nil(t, config, "Config was not nil.")
}

func TestGetConfigWithFileReadError(t *testing.T) {
	config, err := GetConfig(ParsedArgs{}, testGetUser, testConfigFileReaderError)
	assert.NotNil(t, err, "Error was nil.")
	assert.EqualError(t, err, "Couldn't read file \".chill/config.yaml\"", "Error was nil.")
	assert.Nil(t, config, "Config was not nil.")
}

func testGetUser() (*user.User, error) {
	return &user.User{
		HomeDir: ".",
	}, nil
}

func testGetUserError() (*user.User, error) {
	return nil, errors.New("Error fetching user")
}

func testHomeDirConfigFileReader(fileName string) ([]byte, error) {
	if fileName == ".chill/config.yaml" {
		return []byte(standardConfigContents), nil
	}

	return nil, fmt.Errorf("Couldn't read file %q", fileName)
}

func testConfigFileReader(fileName string) ([]byte, error) {
	if fileName == "./expected_name.yaml" {
		return []byte(standardConfigContents), nil
	}

	return nil, fmt.Errorf("Couldn't read file %q", fileName)
}

func testConfigFileReaderError(fileName string) ([]byte, error) {
	return nil, fmt.Errorf("Couldn't read file %q", fileName)
}

const standardConfigContents =
"phrases:\n" +
"  - chill\n" +
"  - most chill\n" +
"replacements:\n" +
"  the : the chill \n" +
"stateimageslocation: http://my.imagestore.com\n"
