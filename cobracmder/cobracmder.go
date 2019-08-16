package cobracmder

import (
	"os"
	"fmt"
	"errors"
	"github.com/spf13/cobra"
)

func failAndExit(e error) {
	fmt.Println(e)
	os.Exit(1)
}

type fileLoader func(string)

func Setup(f fileLoader) *cobra.Command {
	c := &cobra.Command {
		Use: "deploymentmanager-tf",
		Short: "A YAML-to-Terraform converter",
		Long: "A YAML-to-Terraform converter created to facilitate the usage of Google Cloud Platform",
		Run: func(c *cobra.Command, args []string) {
			path, err := c.Flags().GetString("path")
			if err != nil {
				failAndExit(err)
			}

			if path == "" {
				failAndExit(errors.New("No path provided"))
			}

			f(path)
		},
	}
	c.Flags().String("path", "", "The file to be loaded")
	return c
}