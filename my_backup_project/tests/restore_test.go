package tests

import (
	"my_backup_project/internal/backup"
	"my_backup_project/internal/restore"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestRestoreBackup(t *testing.T) {
	workDir := createTestWorkDir(t)
	defer os.RemoveAll(workDir)

	backupDir := createTestBackupDir(t)
	defer os.RemoveAll(backupDir)

	restoreDir, err := os.MkdirTemp("", "restoreDir")
	if err != nil {
		t.Fatalf("failed to create restore directory: %v", err)
	}
	defer os.RemoveAll(restoreDir)
	timeCreate := time.Now().Format("2006-01-02_15-04-05")
	err = backup.CreateFullBackUp(workDir, backupDir, timeCreate)
	if err != nil {
		t.Fatalf("CreateFullBackup failed: %v", err)
	}
	fullBackupPath := filepath.Join(backupDir, timeCreate)
	err = restore.RestoreBackUp(fullBackupPath, restoreDir)
	if err != nil {
		t.Fatalf("RestoreBackup from full backup failed: %v", err)
	}
	checkFileContent(t, filepath.Join(restoreDir, "file1.txt"), "original content")
	checkFileContent(t, filepath.Join(restoreDir, "folder1/file2.txt"), "original content")
}
