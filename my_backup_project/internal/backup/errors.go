package backup

import (
	"errors"
	"fmt"
	"os"
	"syscall"
)

func HandleFileError(path string, err error, action string) error {
	if errors.Is(err, syscall.ENOSPC) {
		return fmt.Errorf("disk full: cannot %s %s. Free up some disk space and try again", action, path)
	}
	if errors.Is(err, syscall.EACCES) || errors.Is(err, syscall.EPERM) || errors.Is(err, os.ErrPermission) {
		return fmt.Errorf("permission denied: cannot %s %s. Check file permissions", action, path)
	}
	if errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("%s does not exist. Please check the path and try again", path)
	}
	return fmt.Errorf("failed to %s %s: %w", action, path, err)
}
