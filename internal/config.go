package internal

import (
	"fmt"
	"os"
	"path"

	"gopkg.in/yaml.v3"
)

type ReposConfig struct {
	Wd string
}

func UpdateConfig(config ReposConfig) error {
	configPath, err := os.UserConfigDir()
	if err != nil {
		return err
	}

	b, err := yaml.Marshal(config)
	if err != nil {
		return err
	}

	return os.WriteFile(path.Join(configPath, "repos", "config.yaml"), b, 0644)
}

func ReadConfig() (ReposConfig, error) {
	configPath, err := os.UserConfigDir()
	if err != nil {
		return ReposConfig{}, err
	}

	f, err := os.ReadFile(path.Join(configPath, "repos", "config.yaml"))
	if err != nil {
		createFile(configPath)
		return ReposConfig{}, err
	}

	var c ReposConfig
	if err := yaml.Unmarshal(f, &c); err != nil {
		return ReposConfig{}, err
	}

	return c, nil
}

func createFile(configPath string) error {
	fmt.Println(path.Dir(configPath))
	err := os.MkdirAll(path.Dir(configPath), os.ModePerm)
	if err != nil {
		return err
	}
	f, err := os.Create(configPath)
	if err != nil {
		return err
	}

	c := ReposConfig{
		Wd: "",
	}

	b, err := yaml.Marshal(c)
	if err != nil {
		return err
	}

	_, err = f.Write(b)
	return err
}
