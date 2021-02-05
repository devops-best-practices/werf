package helm2

type ReleaseData struct{}

type MaintenanceHelperOptions struct {
	ReleaseStorageNamespace string
	ReleaseStorageType      string
}

func NewMaintenanceHelper(opts MaintenanceHelperOptions) *MaintenanceHelper {
	return &MaintenanceHelper{
		ReleaseStorageNamespace: opts.ReleaseStorageNamespace,
		ReleaseStorageType:      opts.ReleaseStorageType,
	}
}

type MaintenanceHelper struct {
	ReleaseStorageNamespace string
	ReleaseStorageType      string
}

func (helper *MaintenanceHelper) CheckStorageAvailable() (bool, error) {
	return true, nil
}

func (helper *MaintenanceHelper) GetReleasesList() ([]string, error) {
	return []string{"quickstart-application"}, nil
}

func (helper *MaintenanceHelper) GetReleaseData(releaseName string) (*ReleaseData, error) {
	return nil, nil
}

func (helper *MaintenanceHelper) ForgetReleaseStorageMetadata(releaseName string) error {
	return nil
}
