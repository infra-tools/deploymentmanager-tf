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

type fileLoader func(string, string)

func Setup(f fileLoader) *cobra.Command {
	c := &cobra.Command {
		Use: "deploymentmanager-tf",
		Short: "A YAML-to-Terraform converter",
		Long: "A YAML-to-Terraform converter created to facilitate the usage of Google Cloud Platform",
		Run: func(c *cobra.Command, args []string) {
			file, err := c.Flags().GetString("file")
			if err != nil {
				failAndExit(err)
			}

			if file == "" {
				failAndExit(errors.New("No file specified"))
			}

			output, err := c.Flags().GetString("output")
			fmt.Println(output)
			if err != nil {
				failAndExit(err)
			}

			if output == "" {
				failAndExit(errors.New("No output location specified"))
			}

			f(file, output)
		},
	}
	c.Flags().String("file", "", "The file to be loaded")
	c.Flags().String("output", "", "The write location of the new file")
	return c
}