package healthcheck

import (
	"fmt"
	"path/filepath"

	"github.com/NII-DG/gogs/internal/conf"
	"github.com/NII-DG/gogs/internal/utils"
)

/*
 Check connecting File System
*/
func CheckFileSystem() error {
	filepath := filepath.Join(conf.DG.HealthFilePath, conf.DG.HealthFileName)
	if utils.ExistData(filepath) {
		return nil
	} else {
		return fmt.Errorf("Not exits %s for health check", filepath)
	}
}
