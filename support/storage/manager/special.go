package manager

import (
	"fmt"
	"support/logger"
	"support/storage/config"
	"support/storage/driver"
)

// Corrected to return *storage.StorageConfig instead of *storage.StorageHelper
type GetStorageConfig func() *config.StorageConfig

func initSpecialOss(GetConfig GetStorageConfig) (*driver.StorageHelper, error) {
	cfg := GetConfig()
	if cfg == nil {
		logger.Warn("InitSpecialOss storageConfig is nil")
		return nil, fmt.Errorf("InitSpecialOss storageConfig is nil")
	}
	return driver.New2(cfg)
}
