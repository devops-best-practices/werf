package main

import (
	"errors"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"

	"github.com/flant/werf/cmd/werf/build"
	"github.com/flant/werf/cmd/werf/build_and_publish"
	"github.com/flant/werf/cmd/werf/cleanup"
	"github.com/flant/werf/cmd/werf/deploy"
	"github.com/flant/werf/cmd/werf/dismiss"
	"github.com/flant/werf/cmd/werf/publish"
	"github.com/flant/werf/cmd/werf/purge"

	helm_secret_edit "github.com/flant/werf/cmd/werf/helm/secret/edit"
	helm_secret_extract "github.com/flant/werf/cmd/werf/helm/secret/extract"
	helm_secret_generate "github.com/flant/werf/cmd/werf/helm/secret/generate"
	helm_secret_key_generate "github.com/flant/werf/cmd/werf/helm/secret/key_generate"
	helm_secret_regenerate "github.com/flant/werf/cmd/werf/helm/secret/regenerate"

	"github.com/flant/werf/cmd/werf/tools/slugify"

	images_cleanup "github.com/flant/werf/cmd/werf/images/cleanup"
	images_publish "github.com/flant/werf/cmd/werf/images/publish"
	images_purge "github.com/flant/werf/cmd/werf/images/purge"

	stages_build "github.com/flant/werf/cmd/werf/stages/build"
	stages_cleanup "github.com/flant/werf/cmd/werf/stages/cleanup"
	stages_purge "github.com/flant/werf/cmd/werf/stages/purge"

	host_cleanup "github.com/flant/werf/cmd/werf/host/cleanup"
	host_purge "github.com/flant/werf/cmd/werf/host/purge"

	helm_deploy_chart "github.com/flant/werf/cmd/werf/helm/deploy_chart"
	helm_generate_chart "github.com/flant/werf/cmd/werf/helm/generate_chart"
	helm_get_service_values "github.com/flant/werf/cmd/werf/helm/get_service_values"

	"github.com/flant/werf/cmd/werf/completion"
	"github.com/flant/werf/cmd/werf/docs"
	"github.com/flant/werf/cmd/werf/version"

	"github.com/flant/werf/cmd/werf/common"
	"github.com/flant/werf/cmd/werf/common/templates"
	"github.com/flant/werf/pkg/logger"
	"github.com/flant/werf/pkg/process_exterminator"
)

func main() {
	trapTerminationSignals()

	logger.Init(logger.Options{})

	if err := process_exterminator.Init(); err != nil {
		logger.LogError(fmt.Errorf("process exterminator initialization error: %s", err))
		os.Exit(1)
	}

	rootCmd := &cobra.Command{
		Use:   "werf",
		Short: "Werf helps to implement and support Continuous Integration and Continuous Delivery",
		Long: common.GetLongCommandDescription(`Werf helps to implement and support Continuous Integration and Continuous Delivery.

Find more information at https://flant.github.io/werf`),
		SilenceUsage:  true,
		SilenceErrors: true,
	}

	groups := templates.CommandGroups{
		{
			Message: "Main Commands:",
			Commands: []*cobra.Command{
				build.NewCmd(),
				publish.NewCmd(),
				build_and_publish.NewCmd(),
				deploy.NewCmd(),
				dismiss.NewCmd(),
				cleanup.NewCmd(),
				purge.NewCmd(),
			},
		},
		{
			Message: "Tools Commands:",
			Commands: []*cobra.Command{
				slugify.NewCmd(),
			},
		},
		{
			Message: "Lowlevel Management Commands:",
			Commands: []*cobra.Command{
				stagesCmd(),
				imagesCmd(),
				helmCmd(),
				hostCmd(),
			},
		},
	}
	groups.Add(rootCmd)

	templates.ActsAsRootCommand(rootCmd, groups...)

	rootCmd.AddCommand(
		completion.NewCmd(rootCmd),
		version.NewCmd(),
		docs.NewCmd(),
	)

	if err := rootCmd.Execute(); err != nil {
		logger.LogError(err)
		os.Exit(1)
	}
}

func imagesCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "images",
		Short: "Commands to work with images",
	}
	cmd.AddCommand(
		images_publish.NewCmd(),
		images_cleanup.NewCmd(),
		images_purge.NewCmd(),
	)

	return cmd
}

func stagesCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "stages",
		Short: "Commands to work with stages, which are cache for images",
	}
	cmd.AddCommand(
		stages_build.NewCmd(),
		stages_cleanup.NewCmd(),
		stages_purge.NewCmd(),
	)

	return cmd
}

func helmCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "helm",
		Short: "Commands to manage deployment with helm",
	}
	cmd.AddCommand(
		helm_get_service_values.NewCmd(),
		helm_generate_chart.NewCmd(),
		helm_deploy_chart.NewCmd(),
		secretCmd(),
	)

	return cmd
}

func hostCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "host",
	}
	cmd.AddCommand(
		host_cleanup.NewCmd(),
		host_purge.NewCmd(),
	)

	return cmd
}

func secretCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "secret",
		Short: "Commands to work with secrets",
	}
	cmd.AddCommand(
		helm_secret_key_generate.NewCmd(),
		helm_secret_generate.NewCmd(),
		helm_secret_extract.NewCmd(),
		helm_secret_edit.NewCmd(),
		helm_secret_regenerate.NewCmd(),
	)

	return cmd
}

func trapTerminationSignals() {
	c := make(chan os.Signal, 1)
	signals := []os.Signal{os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT}
	signal.Notify(c, signals...)
	go func() {
		<-c

		logger.LogError(errors.New("interrupted"))
		os.Exit(17)
	}()
}