package urlshortner

import (
	"errors"
	"flag"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

var paths map[string]string

//ResolvePath resolves resirect path from memory based on given path
func ResolvePath(path string) (string, error) {
	if resolvedPath, ok := paths[path]; ok {
		return resolvedPath, nil
	}
	return "", errors.New("Path doesnt exists")
}

func init() {
	paths = resolvePaths()
}

func resolvePaths() map[string]string {
	filepath := flag.String("pathFile", "paths.yml", "The yml file to be used as source")
	flag.Parse()
	ymlData, err := ioutil.ReadFile(*filepath)
	if err != nil {
		log.Fatalf("File Opening/Reading error: %v \n", err)
	}
	paths = parseYAML(ymlData)
	return paths
}

func parseYAML(ymlData []byte) map[string]string {
	var paths struct {
		Paths map[string]string `yaml:"paths"`
	}
	err := yaml.Unmarshal(ymlData, &paths)
	if err != nil {
		log.Fatalf("Unable to convert data to yml \n")
	}
	log.Printf("paths data %v \n", paths)
	return paths.Paths
}
