package helm

import (
	"fmt"
	"os"

	"github.com/werf/werf/pkg/deploy/helm2"

	"github.com/spf13/cobra"
	"github.com/werf/werf/cmd/werf/common"
	"github.com/werf/werf/pkg/werf"
)

var migrate2To3CommonCmdData common.CmdData

var migrate2ToCmdData struct {
	ReleaseStorageNamespace string
	ReleaseStorageType      string

	Release   string
	Namespace string
}

func NewMigrate2To3Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                   "migrate2to3",
		DisableFlagsInUseLine: true,
		Short:                 "Print Helm Release name that will be used in current configuration with specified params",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := common.ProcessLogOptions(&migrate2To3CommonCmdData); err != nil {
				common.PrintHelp(cmd)
				return err
			}

			return runMigrate2To3()
		},
	}

	common.SetupTmpDir(&migrate2To3CommonCmdData, cmd)
	common.SetupHomeDir(&migrate2To3CommonCmdData, cmd)

	common.SetupKubeConfig(&migrate2To3CommonCmdData, cmd)
	common.SetupKubeConfigBase64(&migrate2To3CommonCmdData, cmd)
	common.SetupKubeContext(&migrate2To3CommonCmdData, cmd)

	common.SetupLogOptions(&migrate2To3CommonCmdData, cmd)

	cmd.Flags().StringVarP(&migrate2ToCmdData.Release, "release", "", os.Getenv("WERF_RELEASE"), "Target release name which should be migrated from helm 2 to helm 3 (default $WERF_RELEASE)")
	cmd.Flags().StringVarP(&migrate2ToCmdData.Namespace, "namespace", "", os.Getenv("WERF_NAMESPACE"), "Target kubernetes namespace for helm 3 release (default $WERF_NAMESPACE)")

	return cmd
}

func runMigrate2To3() error {
	if err := werf.Init(*migrate2To3CommonCmdData.TmpDir, *migrate2To3CommonCmdData.HomeDir); err != nil {
		return fmt.Errorf("initialization error: %s", err)
	}

	releaseName := migrate2ToCmdData.Namespace
	if releaseName == "" {
		return fmt.Errorf("--release (or WERF_RELEASE env var) required")
	}

	helm2MaintenanceHelper := helm2.NewMaintenanceHelper(helm2.MaintenanceHelperOptions{})
	if available, err := helm2MaintenanceHelper.CheckStorageAvailable(); err != nil {
		return err
	} else if !available {
		return fmt.Errorf("helm 2 release storage is not available")
	}

	existingReleases, err := helm2MaintenanceHelper.GetReleasesList()
	if err != nil {
		return fmt.Errorf("error getting existing helm 2 releases to perform check: %s", err)
	}

	foundHelm2Release := false
	for _, existingReleaseName := range existingReleases {
		if releaseName == existingReleaseName {
			foundHelm2Release = true
			break
		}
	}

	if !foundHelm2Release {
		return fmt.Errorf("not found helm 2 release %q", releaseName)
	}

	releaseData, err := helm2MaintenanceHelper.GetReleaseData(releaseName)
	if err != nil {
		return fmt.Errorf("unable to get helm 2 release %q info: %s", releaseName, err)
	}

	_ = releaseData
	// TODO: set resource annotations and labels
	// TODO: set migration in progress kind of "lock", release it when migration is done

	if err := helm2MaintenanceHelper.ForgetReleaseStorageMetadata(releaseName); err != nil {
		return fmt.Errorf("unable to forget helm 2 release storage metadata for the release %q: %s", releaseName, err)
	}

	return nil
}
