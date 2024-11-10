package backup

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"time"
)

func CopyFile(sourcePath, targetPath string) error {
	srcFile, err := os.Open(sourcePath)
	if err != nil {
		return HandleFileError(sourcePath, err, "open source file")
	}
	defer srcFile.Close()
	dstFile, err := os.Create(targetPath)
	if err != nil {
		return HandleFileError(targetPath, err, "create destination file")
	}
	defer dstFile.Close()
	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		return HandleFileError(targetPath, err, "copy data")
	}
	srcInfo, err := os.Stat(sourcePath)
	if err != nil {
		return HandleFileError(sourcePath, err, "get file info")
	}
	if err := os.Chmod(targetPath, srcInfo.Mode()); err != nil {
		return HandleFileError(targetPath, err, "set permissions")
	}
	return nil
}

func CopyDir(srcDir, dstDir string) error {
	return filepath.WalkDir(srcDir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return HandleFileError(path, err, "access")
		}
		realPath, err := filepath.Rel(srcDir, path)
		if err != nil {
			return fmt.Errorf("failed to calculate relative path for %s: %w", path, err)
		}
		targetPath := filepath.Join(dstDir, realPath)
		if d.IsDir() {
			if mkErr := os.MkdirAll(targetPath, d.Type().Perm()); mkErr != nil {
				return HandleFileError(targetPath, mkErr, "create directory")
			}
			return nil
		}
		if copyErr := CopyFile(path, targetPath); copyErr != nil {
			return HandleFileError(path, copyErr, "copy file")
		}
		return nil
	})
}

func IsFileModified(srcFile, dstFile string) (bool, error) {
	srcInfo, err := os.Stat(srcFile)
	if err != nil {
		if os.IsNotExist(err) {
			return true, fmt.Errorf("source file %s does not exist", srcFile)
		}
		if os.IsPermission(err) {
			return false, fmt.Errorf("permission denied: cannot access source file %s", srcFile)
		}
		return false, fmt.Errorf("error accessing source file %s: %w", srcFile, err)
	}

	dstInfo, err := os.Stat(dstFile)
	if err != nil {
		if os.IsNotExist(err) {
			return true, nil
		}
		if os.IsPermission(err) {
			return false, fmt.Errorf("permission denied: cannot access destination file %s", dstFile)
		}
		return false, fmt.Errorf("error accessing destination file %s: %w", dstFile, err)
	}

	isModified := srcInfo.ModTime() != dstInfo.ModTime() || srcInfo.Size() != dstInfo.Size()
	return isModified, nil
}

func GetLastFullBackUp(backupDir string) (string, error) {
	var backupPaths []string
	entries, err := os.ReadDir(backupDir)
	if err != nil {
		return "", fmt.Errorf("failed to read backup directory %s: %w", backupDir, err)
	}
	for _, entry := range entries {
		if entry.IsDir() {
			_, err := time.Parse("2006-01-02_15-04-05", entry.Name())
			if err == nil {
				backupPaths = append(backupPaths, filepath.Join(backupDir, entry.Name()))
			}
		}
	}

	if len(backupPaths) == 0 {
		return "", fmt.Errorf("no valid full backups found in %s", backupDir)
	}
	sort.Slice(backupPaths, func(i, j int) bool {
		return backupPaths[i] > backupPaths[j]
	})
	return backupPaths[0], nil
}
