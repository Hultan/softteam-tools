package resources

import (
	"os"
	"path"
	"path/filepath"
)

// Resources : Handles SoftTeam resources
type Resources struct {
}

// GetExecutablePath : Returns the path of the executable
func (r *Resources) GetExecutablePath() string {
	ex, err := os.Executable()
	if err != nil {
		return ""
	}
	return filepath.Dir(ex)
}

// GetResourcesPath : Returns the resources path
func (r *Resources) GetResourcesPath() string {
	executablePath:=r.GetExecutablePath()
	return path.Join(executablePath, "resources")
}

// GetResourcePath : Gets the path for a single resource file
func (r *Resources) GetResourcePath(fileName string) string {
	resourcesPath:=r.GetResourcesPath()
	resourcePath:=path.Join(resourcesPath, fileName)
	return resourcePath
}
