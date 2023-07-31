package main

import (
	"bytes"
	"embed"
	"io/fs"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

//go:embed graphql
var schemas embed.FS

//go:embed model
var models embed.FS

type gqlgenConfig struct {
	Schema []string `yaml:"schema"`
	Model  *struct {
		Package  string `yaml:"package"`
		Filename string `yaml:"filename"`
	}
}

func main() {
	config := readConfigFile()
	generateSchema(config)
	generateModel(config)
}

func generateSchema(config gqlgenConfig) {
	dir, _ := schemas.ReadDir("graphql")
	schemaDir := getSchemaDir(config)

	for _, f := range dir {
		contents, _ := fs.ReadFile(schemas, "graphql/"+f.Name())
		os.WriteFile(schemaDir+"gorm-gqlgen-relay-"+f.Name(), contents, fs.ModePerm)
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

func generateModel(config gqlgenConfig) {
	dir, _ := models.ReadDir("model")
	modelDir := getModelDir(config)

	for _, f := range dir {
		contents, _ := fs.ReadFile(models, "model/"+f.Name())
		os.WriteFile(
			modelDir+"/gorm-gqlgen-relay-"+f.Name(),
			bytes.Replace(contents, []byte("package model"), []byte("package "+config.Model.Package), 1),
			fs.ModePerm,
		)
	}
}

func getModelDir(config gqlgenConfig) string {
	if config.Model == nil {
		return "./graph/model/"
	}

	split := strings.Split(config.Model.Filename, "/")
	split = split[:len(split)-1]
	return strings.Join(split, "/")
}
