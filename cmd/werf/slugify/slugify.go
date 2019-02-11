package slugify

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/flant/werf/pkg/slug"
)

var CmdData struct {
	Format string
}

func NewCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "slugify STRING",
		DisableFlagsInUseLine: true,
		Short: "Print slugged string by specified format",
		Example: `  $ werf slugify -f kubernetes-namespace feature-fix-2
  feature-fix-2

  $ werf slugify -f kubernetes-namespace 'branch/one/!@#4.4-3'
  branch-one-4-4-3-4fe08955

  $ werf slugify -f kubernetes-namespace My_branch
  my-branch-8ebf2d1d

  $ werf slugify -f helm-release my_release-NAME
  my_release-NAME

  # The result has been trimmed to fit maximum bytes limit:
  $ werf slugify -f helm-release looooooooooooooooooooooooooooooooooooooooooong_string
  looooooooooooooooooooooooooooooooooooooooooong-stri-b150a895

  $ werf slugify -f docker-tag helo/ehlo
  helo-ehlo-b6f6ab1f

  $ werf slugify -f docker-tag 16.04
  16.04`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := runSlugify(args); err != nil {
				cmd.Help()
				fmt.Println()

				return err
			}

			return nil
		},
	}

	cmd.Flags().StringVarP(&CmdData.Format, "format", "f", "", `  r|helm-release:         suitable for Helm Release
 ns|kubernetes-namespace: suitable for Kubernetes Namespace
tag|docker-tag:           suitable for Docker Tag`)

	return cmd
}

func runSlugify(args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("accepts 1 position argument, received %d", len(args))
	}

	data := args[0]

	switch CmdData.Format {
	case "helm-release", "r":
		fmt.Println(slug.HelmRelease(data))
	case "kubernetes-namespace", "ns":
		fmt.Println(slug.KubernetesNamespace(data))
	case "docker-tag", "tag":
		fmt.Println(slug.DockerTag(data))
	case "":
		return fmt.Errorf("--format FORMAT argument required")
	default:
		return fmt.Errorf("unknown format '%s'", CmdData.Format)
	}

	return nil
}
