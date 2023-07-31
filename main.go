package main

import (
	"embed"
	"io/fs"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

//go:embed graphql
var schemas embed.FS

type gqlgenConfig struct {
	Exec *struct {
		Package  string `yaml:"package"`
		Filename string `yaml:"filename"`
	} `yaml:"exec,omitempty"`
	Schema []string `yaml:"schema"`
}

func main() {
	config := readConfigFile()
	generateSchema(config)
}

func generateSchema(config gqlgenConfig) {
	dir, _ := schemas.ReadDir("graphql")
	schemaDir := getSchemaDir(config)

	for _, path := range config.Schema {
		if strings.Contains(path, "*.graphql") {
			for _, f := range dir {
				contents, _ := fs.ReadFile(schemas, "graphql/"+f.Name())
				os.WriteFile(schemaDir+"gorm-gqlgen-relay-"+f.Name(), contents, fs.ModePerm)
			}

			return
		}
	}
}

func getSchemaDir(config gqlgenConfig) string {
	for _, path := range config.Schema {
		strings.Contains(path, "*.graphql")
		return strings.Replace(path, "*.graphql", "", 1)
	}

	panic("No schema directory found")
}

func readConfigFile() gqlgenConfig {
	contents, err := os.ReadFile("./gqlgen.yml")
	if err == nil {
		return parseConfigFile(contents)
	}

	contents, err = os.ReadFile("./gqlgen.yaml")
	if err != nil {
		panic(err)
	}

	return parseConfigFile(contents)
}

func parseConfigFile(contents []byte) (config gqlgenConfig) {
	if err := yaml.Unmarshal(contents, &config); err != nil {
		panic(err)
	}

	return
}
