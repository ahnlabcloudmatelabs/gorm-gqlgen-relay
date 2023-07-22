package main

import (
	"bytes"
	"embed"
	"io/fs"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

//go:embed schema
var schemas embed.FS

//go:embed resolvers
var resolvers embed.FS

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
	generateResolvers(config)
}

func generateSchema(config gqlgenConfig) {
	dir, _ := schemas.ReadDir("schema")
	schemaDir := getSchemaDir(config)

	for _, path := range config.Schema {
		strings.Contains(path, "*.graphql")
	}

	for _, f := range dir {
		contents, _ := fs.ReadFile(schemas, "schema/"+f.Name())
		os.WriteFile(schemaDir+f.Name(), contents, fs.ModePerm)
	}
}

func getSchemaDir(config gqlgenConfig) string {
	for _, path := range config.Schema {
		strings.Contains(path, "*.graphql")
		return strings.Replace(path, "*.graphql", "", 1)
	}

	panic("No schema directory found")
}

func generateResolvers(config gqlgenConfig) {
	dir, _ := resolvers.ReadDir("resolvers")
	resolverDir := getResolverDir(config)
	resolverPackage := config.Exec.Package

	for _, f := range dir {
		contents, _ := fs.ReadFile(schemas, "resolvers/"+f.Name())

		os.WriteFile(
			resolverDir+strings.Replace(f.Name(), "txt", "go", 1),
			bytes.Replace(contents, []byte("package resolver"), []byte("package "+resolverPackage), 1),
			fs.ModePerm,
		)
	}
}

func getResolverDir(config gqlgenConfig) string {
	if config.Exec == nil {
		return ""
	}

	paths := strings.Split(config.Exec.Filename, "/")
	paths = paths[:len(paths)-1]
	return strings.Join(paths, "/")
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
