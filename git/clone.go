package git

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
)

var CloneCmd = &cobra.Command{
	Use:   "clone",
	Short: "clone a GitHub repository",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			log.Fatalln("repository must be specified")
			return
		}

		if err := CloneRepository(args[0], ref, create); err != nil {
			log.Fatalln("error while cloning repository: ", err)
		}
	},
}

var ref string
var create bool

func CloneRepository(repository, ref string, shouldCreate bool) error {
	repo, err := NewGHRepo(repository)
	if err != nil {
		return err
	}

	if err := repo.Clone(viper.GetString("location")); err != nil {
		return err
	}

	if err := repo.Checkout(ref, shouldCreate); err != nil {
		return err
	}

	fmt.Printf("Cloned repository to: %s\n", repo.RepoDir)
	return nil
}

func init() {
	CloneCmd.PersistentFlags().StringVar(&ref, "ref", "master",
		"specific reference to check out")

	CloneCmd.PersistentFlags().BoolVar(&create, "create", false,
		"create the reference if it does not exist")
}
