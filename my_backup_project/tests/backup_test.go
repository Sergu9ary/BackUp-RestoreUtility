package tests

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	"my_backup_project/internal/backup"
)

func checkFileContent(t *testing.T, path, expectedContent string) {
	content, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("failed to read file %s: %v", path, err)
	}
	if string(content) != expectedContent {
		t.Errorf("expected content %q, got %q", expectedContent, content)
	}
}

func findLastBackUp(t *testing.T, backupDir string) string {
	var lastBackupPath string
	entries, err := os.ReadDir(backupDir)
	if err != nil {
		t.Fatalf("failed to read backup directory: %v", err)
	}

	for _, entry := range entries {
		if entry.IsDir() && (lastBackupPath == "" || entry.Name() > filepath.Base(lastBackupPath)) {
			lastBackupPath = filepath.Join(backupDir, entry.Name())
		}
	}

	if lastBackupPath == "" {
		t.Fatal("no backups found")
	}
	return lastBackupPath
}

func createTestWorkDir(t *testing.T) string {
	t.Helper()
	workDir, err := os.MkdirTemp("", "workDir")
	if err != nil {
		t.Fatalf("failed to create temp work directory: %v", err)
	}
	err = os.Mkdir(filepath.Join(workDir, "folder1"), 0755)
	if err != nil {
		t.Fatalf("failed to create folder in workDir: %v", err)
	}

	err = os.WriteFile(filepath.Join(workDir, "file1.txt"), []byte("original content"), 0644)
	if err != nil {
		t.Fatalf("failed to create file1 in workDir: %v", err)
	}

	err = os.Chmod(filepath.Join(workDir, "file1.txt"), 0644)
	if err != nil {
		t.Fatalf("failed to set permissions for file1 in workDir: %v", err)
	}

	err = os.WriteFile(filepath.Join(workDir, "folder1/file2.txt"), []byte("original content"), 0644)
	if err != nil {
		t.Fatalf("failed to create file2 in workDir: %v", err)
	}

	err = os.Chmod(filepath.Join(workDir, "folder1/file2.txt"), 0644)
	if err != nil {
		t.Fatalf("failed to set permissions for file2 in workDir: %v", err)
	}

	return workDir
}

func createTestBackupDir(t *testing.T) string {
	t.Helper()
	backupDir, err := os.MkdirTemp("", "backupDir")
	if err != nil {
		t.Fatalf("failed to create temp backup directory: %v", err)
	}
	if err := os.Chmod(backupDir, 0755); err != nil {
		t.Fatalf("failed to set permissions for backupDir: %v", err)
	}
	return backupDir
}

func TestCreateIncrementalBackup(t *testing.T) {
	workDir := createTestWorkDir(t)
	defer os.RemoveAll(workDir)

	backupDir := createTestBackupDir(t)
	defer os.RemoveAll(backupDir)
	timestamp := time.Now().Format("2006-01-02_15-04-05")
	err := backup.CreateFullBackUp(workDir, backupDir, timestamp)
	if err != nil {
		t.Fatalf("CreateFullBackup failed: %v", err)
	}
	lastFullBackup, err := backup.GetLastFullBackUp(backupDir)
	if err != nil {
		t.Fatalf("failed to find last full backup%v: %v", lastFullBackup, err)
	}
	err = os.WriteFile(filepath.Join(workDir, "file1.txt"), []byte("modified content"), 0644)
	if err != nil {
		t.Fatalf("failed to modify file1 in workDir: %v", err)
	}
	err = backup.CreateIncrementalBackUp(workDir, backupDir, timestamp)
	if err != nil {
		t.Fatalf("CreateIncrementalBackup failed: %v", err)
	}
	lastIncrementalBackup := findLastBackUp(t, backupDir)
	checkFileContent(t, filepath.Join(lastIncrementalBackup, "file1.txt"), "modified content")
}
