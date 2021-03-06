package commands

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/tamada/flaver"
	"github.com/tamada/flaver/errors"
)

func runsOnInteractiveMode(command *cobra.Command, ec *errors.Center) error {
	return ec
}

func performEach(command *cobra.Command, arg string, ec *errors.Center) ([]flaver.Release, error) {
	repo, err := flaver.BuildRepository(arg)
	if err != nil {
		ec.Push(err)
		return nil, ec
	}
	flaverObject, err := flaver.NewFlaver(repo)
	allFlag, err := command.Flags().GetBool("all")
	if allFlag && err == nil {
		return flaverObject.FindAll(repo)
	}
	release, err := flaverObject.Find(repo)
	if release == nil {
		return []flaver.Release{}, err
	}
	return []flaver.Release{release}, err
}

func printReleases(cmd *cobra.Command, name string, releases []flaver.Release) {
	cmd.Printf(`{"product":"%s","versions":[`, name)
	for index, r := range releases {
		if index != 0 {
			cmd.Print(",")
		}
		cmd.Printf(`{"version":"%s","release_date": "%s","url":"%s"}`, r.Version(), r.Date(), r.Url())
	}
	cmd.Println("]}")
}

func perform(command *cobra.Command, args []string) error {
	ec := errors.New()
	if len(args) == 0 {
		runsOnInteractiveMode(command, ec)
	}
	for _, arg := range args {
		releases, err := performEach(command, arg, ec)
		ec.Push(err)
		if err == nil {
			printReleases(command, arg, releases)
		}
	}
	if ec.IsEmpty() {
		return nil
	}
	return ec
}

func Execute() {
	cmd := NewRootCommand()
	cmd.SetOut(os.Stdout)
	if err := cmd.Execute(); err != nil {
		cmd.SetOut(os.Stderr)
		cmd.Println(err)
		os.Exit(1)
	}
}

func init() {
}

// NewRootCommand creates the root command with cobra.
func NewRootCommand() *cobra.Command {
	var root = &cobra.Command{
		Use:  "flaver [PRODUCTs...]",
		Args: cobra.ArbitraryArgs,
		RunE: perform,
		Example: `  flaver tamada/flaver
  flaver github.com/tamada/flaver
  flaver gitlab.com/gitlab-org/gitlab-foss`,
		Version: flaver.Version,
	}
	root.Flags().BoolP("all", "a", false, "finds all versions of the products")
	return root
}
