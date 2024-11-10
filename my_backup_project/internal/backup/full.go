package backup

import (
	"fmt"
	"os"
	"path/filepath"
)

func CreateFullBackUp(workDir, backupDir, timeCreate string) error {
	backupPath := filepath.Join(backupDir, timeCreate)
	err := os.MkdirAll(backupPath, 0755)
	if err != nil {
		if os.IsPermission(err) {
			return fmt.Errorf("permission denied: cannot create backup directory %s", backupPath)
		}
		return fmt.Errorf("failed to create backup directory %s: %w", backupPath, err)
	}
	fmt.Println("Creating full backup...")
	if err := CopyDir(workDir, backupPath); err != nil {
		return fmt.Errorf("failed to complete full backup: %w", err)
	}
	return CopyDir(workDir, backupPath)
}
