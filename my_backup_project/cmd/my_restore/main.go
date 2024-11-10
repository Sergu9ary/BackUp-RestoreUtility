package main

import (
	"fmt"
	"my_backup_project/internal/restore"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("My_restore input: my_restore <backupDir> <restoreDir>")
		return
	}
	backupDir := os.Args[1]
	workDir := os.Args[2]

	err := restore.RestoreBackUp(backupDir, workDir)
	if err != nil {
		fmt.Println("Error restoring backup:", err)
		return
	}

}
