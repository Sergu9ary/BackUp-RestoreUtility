package backup

import (
	"fmt"
	"os"
	"path/filepath"
)

func CreateIncrementalBackUp(workDir, backupDir, timeCreate string) error {
	lastFullBackup, err := GetLastFullBackUp(backupDir)
	if err != nil {
		return fmt.Errorf("could not find last full backup: %w", err)
	}
	incrementalBackupPath := filepath.Join(backupDir, timeCreate)
	if err := os.MkdirAll(incrementalBackupPath, 0755); err != nil {
		return HandleFileError(incrementalBackupPath, err, "create incremental backup directory")
	}
	fmt.Println("Creating incremental backup from", workDir, "to", incrementalBackupPath)
	return filepath.WalkDir(workDir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return HandleFileError(path, err, "access")
		}
		relPath, err := filepath.Rel(workDir, path)
		if err != nil {
			return fmt.Errorf("failed to calculate relative path for %s: %w", path, err)
		}
		backupFilePath := filepath.Join(lastFullBackup, relPath)
		newBackupFilePath := filepath.Join(incrementalBackupPath, relPath)
		if d.IsDir() {
			if mkErr := os.MkdirAll(newBackupFilePath, 0755); mkErr != nil {
				return HandleFileError(newBackupFilePath, mkErr, "create directory")
			}
			return nil
		}
		modified, modErr := IsFileModified(path, backupFilePath)
		if modErr != nil {
			return HandleFileError(path, modErr, "check file modification")
		}
		if modified {
			if copyErr := CopyFile(path, newBackupFilePath); copyErr != nil {
				return HandleFileError(path, copyErr, "copy file")
			}
		}

		return nil
	})
}
