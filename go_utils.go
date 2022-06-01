package go_utils

import (
	"os"
	"os/exec"

	"github.com/BurntSushi/toml"
)

type BuildpackToml struct {
	API string `toml:"api"`

	Buildpack struct {
		ID string `toml:"id"`
		VERSION string `toml:"version"`
	} `toml:"buildpack"`

	Metadata struct {
		TEAMNAME string `toml:"team-name"`
		CONTACT string `toml:"contact"`

		Dependencies []struct {
			URI string `toml:"uri"`
			VERSION string `toml:"version"`
		} `toml:"dependencies"`
	} `toml:"metadata"`

	Stacks []struct {
		ID string `toml:"id"`
	} `toml:"stacks"`
}

// Execute a command and print the stdout and stderr streams
func ExecuteCmd(cmd exec.Cmd) {

}

// Get full buildpack metadata
func GetBuildpackMetadata(tomlPath string) (BuildpackToml, error) {
	var mData BuildpackToml

	file, err := os.Open(tomlPath)
	if err != nil { return mData, err }

	_, err = toml.NewDecoder(file).Decode(&mData)
	if err != nil {
		return mData, err
	}

	return mData, nil
}