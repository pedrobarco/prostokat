package configs

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path"
)

const (
	dirMode  = fs.FileMode(0775)
	fileMode = fs.FileMode(0600)
)

func pathExists(path string) bool {
	_, err := os.Stat(path)
	if os.IsExist(err) {
		return false
	}
	if err != nil {
		log.Fatalf("could not determine if path exists: %s \n", err)
	}
	return true
}

type ConfigFile struct {
	Name string
	Type string
	Path string
}

func (cf *ConfigFile) getConfigFileByName(name string) string {
	return path.Join(cf.Path, name+"."+cf.Type)
}

func (cf *ConfigFile) getConfigFile() string {
	return cf.getConfigFileByName(cf.Name)
}

func (cf *ConfigFile) hasConfigFileByName(name string) bool {
	file := cf.getConfigFileByName(name)
	return pathExists(file)
}

func (cf *ConfigFile) hasConfigFile() bool {
	return cf.hasConfigFileByName(cf.Name)
}

func (cf *ConfigFile) createConfigFolder() {
	err := os.Mkdir(cf.Path, dirMode)
	if err != nil {
		log.Fatalf("could not create config folder: %s \n", err)
	}
	fmt.Printf("+ %s \n", cf.Path)
}

func (cf *ConfigFile) deleteConfigFolder() {
	err := os.RemoveAll(cf.Path)
	if err != nil {
		log.Fatalf("could not delete config folder: %s \n", err)
	}
	fmt.Printf("- %s \n", cf.Path)
}

func (cf *ConfigFile) createConfigFile(name string, content []byte) {
	file := cf.getConfigFileByName(name)
	err := os.WriteFile(file, content, fileMode)
	if err != nil {
		log.Fatalf("could not create config file: %s \n", err)
	}
	fmt.Printf("+ %s \n", file)
}

func (cf *ConfigFile) createDefaultConfigFile(content []byte) {
	cf.createConfigFile(cf.Name, content)
}
