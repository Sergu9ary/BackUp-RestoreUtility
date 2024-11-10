package restore

import (
	"fmt"
	"my_backup_project/internal/backup"
	"os"
	"path/filepath"
)

func RestoreBackUp(backupDir, restoreDir string) error {
	fmt.Println("Restoring backup from", backupDir, "to", restoreDir)
	return filepath.WalkDir(backupDir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return backup.HandleFileError(path, err, "access")
		}
		relPath, err := filepath.Rel(backupDir, path)
		if err != nil {
			return fmt.Errorf("failed to calculate relative path for %s: %w", path, err)
		}
		targetPath := filepath.Join(restoreDir, relPath)
		if d.IsDir() {
			if mkErr := os.MkdirAll(targetPath, d.Type().Perm()); mkErr != nil {
				return backup.HandleFileError(targetPath, mkErr, "create directory")
			}
			return nil
		}
		if copyErr := backup.CopyFile(path, targetPath); copyErr != nil {
			return backup.HandleFileError(path, copyErr, "copy file")
		}
		return nil
	})
}
