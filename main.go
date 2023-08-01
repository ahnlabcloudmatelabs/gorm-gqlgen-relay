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

var graphqlExtension = ".graphql"

type gqlgenConfig struct {
	Schema []string `yaml:"schema"`
	Model  *struct {
		Package  string `yaml:"package"`
		Filename string `yaml:"filename"`
	}
	AutoBind []string `yaml:"autobind"`
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
		filename := strings.Replace(f.Name(), ".graphql", graphqlExtension, 1)
		contents, _ := fs.ReadFile(schemas, "graphql/"+filename)
		os.WriteFile(schemaDir+"gorm-gqlgen-relay-"+filename, contents, fs.ModePerm)
	}
}

func getSchemaDir(config gqlgenConfig) string {
	for _, path := range config.Schema {
		if strings.Contains(path, "*.graphql") {
			return strings.Replace(path, "*.graphql", "", 1)
		}

		if strings.Contains(path, "*.graphqls") {
			graphqlExtension = ".graphqls"
			return strings.Replace(path, "*.graphqls", "", 1)
		}
	}

	panic("No schema directory found (schema: *.graphql or *.graphqls)")
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

	if config.AutoBind == nil || len(config.AutoBind) == 0 {
		panic("No autobind files found")
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
