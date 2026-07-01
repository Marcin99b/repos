package cmd

import (
	"fmt"
	"os"
	"path"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

var defdirCmd = &cobra.Command{
	Use:   "defdir",
	Short: "Get or set default repos directory",
	Long: `defdir without parameter will return default directory, 
	but defdir with argument will update default directory`,
	Run: run,
}

func init() {
	rootCmd.AddCommand(defdirCmd)
}

func run(cmd *cobra.Command, args []string) {
	configPath, err := os.UserConfigDir()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	reposConfigPath := path.Join(configPath, "repos", "config.yaml")

	if len(args) == 0 {
		value, err := get(reposConfigPath)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println(value)
	} else if len(args) == 1 {
		_, err := save(reposConfigPath, args[0])
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	} else {
		fmt.Println("Too many arguments...")
	}
}

func get(configPath string) (string, error) {
	c, err := readConfig(configPath)
	if err != nil {
		return "", err
	}
	return c.Defdir, nil
}

func save(configPath string, value string) (string, error) {
	c, err := readConfig(configPath)
	if err != nil {
		return "", err
	}

	c.Defdir = value
	b, err := yaml.Marshal(c)
	if err != nil {
		return "", err
	}

	return "", os.WriteFile(configPath, b, 0644)
}

func readConfig(configPath string) (ReposConfig, error) {
	f, err := os.ReadFile(configPath)
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
		fmt.Println(err.Error())
		return err
	}
	f, err := os.Create(configPath)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	c := ReposConfig{
		Defdir: "",
	}

	b, err := yaml.Marshal(c)
	if err != nil {
		return err
	}

	_, err = f.Write(b)
	return err
}
